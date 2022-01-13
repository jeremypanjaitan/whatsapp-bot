package manager

import (
	"whatsapp-bot/config"
	"whatsapp-bot/infra"
	"whatsapp-bot/repo"
)

type PhoneNumRepoManagerEntity interface {
	GetPhoneNumRepo() repo.PhoneNumRepoEntity
}

type PhoneNumRepoManager struct {
	csvRepo repo.PhoneNumRepoEntity
}

func NewPhoneNumRepoManager(
	appConfig config.AppConfigEntity,
	csvInfra infra.CsvInfraEntity,
) PhoneNumRepoManagerEntity {
	csvRepo := repo.NewPhoneNumRepo(appConfig, csvInfra)
	return &PhoneNumRepoManager{csvRepo: csvRepo}
}

func (c *PhoneNumRepoManager) GetPhoneNumRepo() repo.PhoneNumRepoEntity {
	return c.csvRepo
}
