package manager

import (
	"whatsapp-bot/config"
	"whatsapp-bot/infra"
)

type WhatsappManagerEntity interface {
	GetWhatsappInfra() infra.WhatsappInfraEntity
}

type WhatsappManager struct {
	whatsAppInfra infra.WhatsappInfraEntity
}

func NewWhatsappManager(appConfig config.AppConfigEntity) WhatsappManagerEntity {
	whatsAppInfra := infra.NewWhatsappInfra(appConfig)
	return &WhatsappManager{whatsAppInfra: whatsAppInfra}
}

func (w *WhatsappManager) GetWhatsappInfra() infra.WhatsappInfraEntity {
	return w.whatsAppInfra
}
