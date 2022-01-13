package config

import (
	"os"
	"strconv"
)

type AppConfigEntity interface {
	GetWhatsAppConnTimeOut() int
	GetQrFileName() string
	GetRemoteJidSuffix() string
}

type AppConfig struct {
	whatsAppConnTimeOut int
	qrFileName          string
	remoteJidSuffix     string
}

func NewAppConfig() AppConfigEntity {
	whatsAppConnTimeOut, _ := strconv.Atoi(os.Getenv("WHATSAPP_CONN_TIMEOUT"))
	qrFileName := os.Getenv("QR_FILE_NAME")
	remoteJidSuffix := os.Getenv("REMOTE_JID_SUFFIX")

	return &AppConfig{
		whatsAppConnTimeOut: whatsAppConnTimeOut,
		qrFileName:          qrFileName,
		remoteJidSuffix:     remoteJidSuffix,
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
