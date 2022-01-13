package app

import (
	"log"
	"whatsapp-bot/config"
	"whatsapp-bot/manager"

	"github.com/joho/godotenv"
)

type AppEntity interface {
	Run()
}

func init() {
	log.SetFlags(log.Lshortfile)
	godotenv.Load()
}

type App struct {
	whatsappInfraManager manager.WhatsappManagerEntity
	appConfig            config.AppConfigEntity
}

func NewApp() AppEntity {
	appConfig := config.NewAppConfig()
	whatsappInfraManager := manager.NewWhatsappManager(appConfig)
	return &App{
		appConfig:            appConfig,
		whatsappInfraManager: whatsappInfraManager}
}

func (a *App) Run() {
	whatsappInfra := a.whatsappInfraManager.GetWhatsappInfra()
	whatsappInfra.GenerateQrCode()
	whatsappInfra.Login()
	whatsappInfra.SendMessage("6287883139903", "Hello Jeremy Panjaitan")
}
