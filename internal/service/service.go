package service

import (
	"pro-link-api/internal/client"
	"pro-link-api/internal/config"
	"pro-link-api/internal/storage"
)

type (
	Service struct {
		Config                *config.Config
		NotificationClient    *client.NotificationClient
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

	UserService struct {
		*Service
	}
)

func New(connection *storage.Storage, clientConn *client.Client, config *config.Config) *Service {
	notificationClient := client.NewNotificationClient(config, clientConn.SMTP)
	accountStorage := storage.NewAccountStorage(connection)
	educationStorage := storage.NewEducationStorage(connection)
	experienceStorage := storage.NewExperienceStorage(connection)
	languageStorage := storage.NewLanguageStorage(connection)
	profileStorage := storage.NewProfileStorage(connection)
	skillStorage := storage.NewSkillStorage(connection)
	websiteProfileStorage := storage.NewWebsiteProfileStorage(connection)
	return &Service{
		Config:                config,
		NotificationClient:    notificationClient,
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

func NewUserService(service *Service) *UserService {
	return &UserService{
		Service: service,
	}
}
