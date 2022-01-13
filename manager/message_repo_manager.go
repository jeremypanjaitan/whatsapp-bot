package manager

import (
	"whatsapp-bot/config"
	"whatsapp-bot/infra"
	"whatsapp-bot/repo"
)

type MessageRepoManagerEntity interface {
	GetMessageRepo() repo.MessageRepoEntity
}

type MessageRepoManager struct {
	messageRepo repo.MessageRepoEntity
}

func NewMessageRepoManager(
	appConfig config.AppConfigEntity,
	textInfra infra.TextInfraEntity,
) MessageRepoManagerEntity {
	messageRepo := repo.NewMessageRepo(appConfig, textInfra)
	return &MessageRepoManager{messageRepo: messageRepo}
}

func (m *MessageRepoManager) GetMessageRepo() repo.MessageRepoEntity {
	return m.messageRepo
}
