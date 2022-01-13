package app

import (
	"log"
	"whatsapp-bot/config"
	"whatsapp-bot/infra"
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
	csvRepoManager       manager.PhoneNumRepoManagerEntity
	appConfig            config.AppConfigEntity
	messageRepoManager   manager.MessageRepoManagerEntity
}

func NewApp() AppEntity {
	appConfig := config.NewAppConfig()
	csvInfra := infra.NewCsvInfra()
	textInfra := infra.NewTextInfra()
	whatsappInfraManager := manager.NewWhatsappManager(appConfig)
	messageRepoManager := manager.NewMessageRepoManager(appConfig, textInfra)
	csvRepoManager := manager.NewPhoneNumRepoManager(appConfig, csvInfra)
	return &App{
		appConfig:            appConfig,
		whatsappInfraManager: whatsappInfraManager,
		csvRepoManager:       csvRepoManager,
		messageRepoManager:   messageRepoManager,
	}

}

func (a *App) Run() {
	whatsappInfra := a.whatsappInfraManager.GetWhatsappInfra()
	whatsappInfra.GenerateQrCode()
	whatsappInfra.Login()
	phoneNumbers, err := a.csvRepoManager.GetPhoneNumRepo().ReadCsvFile()
	if err != nil {
		log.Panicln(err)
	}

	messsage, err := a.messageRepoManager.GetMessageRepo().ReadMessage()
	if err != nil {
		log.Panicln(err)
	}
	for _, phoneNumber := range *phoneNumbers {
		whatsappInfra.SendMessage(phoneNumber.Number, *messsage)
		log.Println(phoneNumber.Number)
	}
}
