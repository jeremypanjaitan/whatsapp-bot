package repo

import (
	"whatsapp-bot/config"
	"whatsapp-bot/infra"
)

type MessageRepoEntity interface {
	ReadMessage() (*string, error)
}

type MessageRepo struct {
	textInfra infra.TextInfraEntity
	appConfig config.AppConfigEntity
}

func NewMessageRepo(
	appConfig config.AppConfigEntity,
	textInfra infra.TextInfraEntity,
) MessageRepoEntity {
	return &MessageRepo{
		textInfra: textInfra,
		appConfig: appConfig,
	}
}

func (m *MessageRepo) ReadMessage() (*string, error) {
	message, err := m.textInfra.ReadFile(m.appConfig.GetMessageFileName())
	if err != nil {
		return nil, err
	}
	return &message, nil
}
