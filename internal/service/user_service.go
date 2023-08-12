package service

import (
	"context"
	"pro-link-api/api"
	"pro-link-api/internal/model"
	utils "pro-link-api/internal/pkg/utils"
	"time"

	"github.com/samber/lo"
	"gorm.io/gorm"
)

type IUserService interface {
	SaveUserInfo(c context.Context, data *api.ProfileRequest) (*api.SaveResponse, error)
}

func (s *UserService) SaveUserInfo(c context.Context, data *api.ProfileRequest) (*api.SaveResponse, error) {

	tx, id, _, err := utils.GetUserIdAndTrx(c)
	if err != nil {
		return nil, err
	}

	currentDate := time.Now()

	err = s.saveProfile(c, tx, data, id, currentDate)
	if err != nil {
		return nil, err
	}

	return &api.SaveResponse{
		Message: "success",
		Code:    "200",
	}, nil
}

func (s *UserService) saveProfile(c context.Context, tx *gorm.DB, data *api.ProfileRequest, accId int, now time.Time) error {
	profileReq := data.Data
	profile, err := s.ProfileStorage.FindByAccId(c, data.Data.AccId)
	if err != nil {
		return err
	}

	prfModel := &model.Profile{}

	if profile.PrfID == 0 {
		prfModel = profile
		prfModel.PrfCreatedBy = accId
	} else {
		prfModel.PrfUpdatedBy = accId
		prfModel.PrfUpdatedDate = &now
	}

	prfModel.PrfFirstName = profileReq.FirstName
	prfModel.PrfLastName = profileReq.LastName
	prfModel.PrfAbout = profileReq.About
	prfModel.PrfAddress = profileReq.Address
	prfModel.PrfPhoneNumber = profileReq.PhoneNumber
	prfModel.PrfPhoneType = profileReq.PhoneType
	prfModel.PrfBirthDate = profileReq.BirthDay
	prfModel.PrfBirthMonth = profileReq.BirthMonth

	prfRes, err := s.ProfileStorage.Save(tx, c, prfModel)
	if err != nil {
		return err
	}

	err = s.saveWebProfile(c, tx, prfRes, profileReq.Website, now)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) saveWebProfile(c context.Context, tx *gorm.DB, prfRes *model.Profile, webList []*api.Website, now time.Time) error {

	webProList, err := s.WebsiteProfileStorage.FindByPrfId(c, prfRes.PrfID)

	if err != nil {
		return err
	}

	toMap := lo.SliceToMap(webProList, func(data *model.WebsiteProfile) (int, *model.WebsiteProfile) {
		return data.PrfID, data
	})

	toModelList := make([]*model.WebsiteProfile, 0)
	for _, req := range webList {
		var toModel *model.WebsiteProfile
		if data, found := toMap[req.Id]; found {
			toModel = data
			toModel.WebUpdatedDate = &now
			toModel.WebUpdatedBy = prfRes.AccID
			delete(toMap, req.Id)
		} else {
			toModel = &model.WebsiteProfile{}
			toModel.WebCreatedBy = prfRes.AccID
		}

		toModel.WebName = req.Website
		toModel.WebType = req.WebsiteType
		toModelList = append(toModelList, toModel)
	}

	err = s.WebsiteProfileStorage.BulkSave(tx, c, toModelList)
	if err != nil {
		return err
	}

	deleteList := make([]int, 0)
	for _, v := range toMap {
		deleteList = append(deleteList, v.PrfID)
	}

	err = s.WebsiteProfileStorage.BulkDelete(tx, c, deleteList)
	if err != nil {
		return err
	}

	return nil
}
