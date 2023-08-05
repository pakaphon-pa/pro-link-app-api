package service

import (
	"pro-link-api/internal/config"
	"pro-link-api/internal/storage"
)

type (
	Service struct {
		Config                *config.Config
		Storage               *storage.Storage
		AccountStorage        storage.IAccountStorage
		EducationStorage      storage.IEducationStorage
		ExperienceStorage     storage.IExperienceStorage
		LanguageStorage       storage.ILanguageStorage
		ProfileStorage        storage.IProfileStorage
		SkillStorage          storage.ISkillStorage
		WebsiteProfileStorage storage.IWebsiteProfileStorage
	}

	AuthService struct {
		*Service
	}
)

func New(connection *storage.Storage, config *config.Config) *Service {
	accountStorage := storage.NewAccountStorage(connection)
	educationStorage := storage.NewEducationStorage(connection)
	experienceStorage := storage.NewExperienceStorage(connection)
	languageStorage := storage.NewLanguageStorage(connection)
	profileStorage := storage.NewProfileStorage(connection)
	skillStorage := storage.NewSkillStorage(connection)
	websiteProfileStorage := storage.NewWebsiteProfileStorage(connection)
	return &Service{
		Config:                config,
		Storage:               connection,
		AccountStorage:        accountStorage,
		EducationStorage:      educationStorage,
		ExperienceStorage:     experienceStorage,
		LanguageStorage:       languageStorage,
		ProfileStorage:        profileStorage,
		SkillStorage:          skillStorage,
		WebsiteProfileStorage: websiteProfileStorage,
	}
}

func NewAuthService(service *Service) *AuthService {
	return &AuthService{
		Service: service,
	}
}
