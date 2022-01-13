package config

import (
	"os"
	"strconv"
)

type AppConfigEntity interface {
	GetWhatsAppConnTimeOut() int
	GetQrFileName() string
	GetRemoteJidSuffix() string
	GetPhoneNumFileName() string
	GetMessageFileName() string
}

type AppConfig struct {
	whatsAppConnTimeOut int
	qrFileName          string
	remoteJidSuffix     string
	phoneNumFileName    string
	messageFileName     string
}

func NewAppConfig() AppConfigEntity {
	whatsAppConnTimeOut, _ := strconv.Atoi(os.Getenv("WHATSAPP_CONN_TIMEOUT"))
	qrFileName := os.Getenv("QR_FILE_NAME")
	remoteJidSuffix := os.Getenv("REMOTE_JID_SUFFIX")
	phoneNumFileName := os.Getenv("PHONE_NUMBER_FILE_NAME")
	messageFileName := os.Getenv("MESSAGE_FILE_NAME")

	return &AppConfig{
		whatsAppConnTimeOut: whatsAppConnTimeOut,
		qrFileName:          qrFileName,
		remoteJidSuffix:     remoteJidSuffix,
		phoneNumFileName:    phoneNumFileName,
		messageFileName:     messageFileName,
	}
}

func (a *AppConfig) GetWhatsAppConnTimeOut() int {
	return a.whatsAppConnTimeOut
}

func (a *AppConfig) GetQrFileName() string {
	return a.qrFileName
}

func (a *AppConfig) GetRemoteJidSuffix() string {
	return a.remoteJidSuffix
}

func (a *AppConfig) GetPhoneNumFileName() string {
	return a.phoneNumFileName
}

func (a *AppConfig) GetMessageFileName() string {
	return a.messageFileName
}
