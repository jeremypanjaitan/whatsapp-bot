package infra

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"time"
	"whatsapp-bot/config"

	"github.com/Rhymen/go-whatsapp"
	"github.com/skip2/go-qrcode"
)

type WhatsappInfraEntity interface {
	Login()
	GenerateQrCode()
	SendMessage(phoneNumber string, message string) string
}

type WhatsappInfra struct {
	wac             *whatsapp.Conn
	qrChan          chan string
	qrFileName      string
	remoteJidSuffix string
}

func NewWhatsappInfra(appConfig config.AppConfigEntity) WhatsappInfraEntity {
	wac, err := whatsapp.NewConn(time.Duration(appConfig.GetWhatsAppConnTimeOut()) * time.Second)
	wac.SetClientVersion(3, 2123, 7)
	if err != nil {
		log.Panicln(err)
	}
	qrChan := make(chan string)
	return &WhatsappInfra{
		wac:             wac,
		qrChan:          qrChan,
		qrFileName:      appConfig.GetQrFileName(),
		remoteJidSuffix: appConfig.GetRemoteJidSuffix(),
	}
}

func (w *WhatsappInfra) Login() {
	_, err := w.wac.Login(w.qrChan)
	if err != nil {
		log.Panicln(err)
	}
}

func (w *WhatsappInfra) GenerateQrCode() {
	go func() {
		err := qrcode.WriteFile(<-w.qrChan, qrcode.Medium, 256, w.qrFileName+".png")

		if err != nil {
			log.Panicln(err)
		}
	}()
}

func (w *WhatsappInfra) SendMessage(phoneNumber string, message string) string {
	b := make([]byte, 10)
	rand.Read(b)
	msgID := strings.ToUpper(hex.EncodeToString(b))
	textMessage := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: fmt.Sprintf("%s@%s", phoneNumber, w.remoteJidSuffix),
			Id:        msgID,
			FromMe:    true,
			Status:    whatsapp.DeliveryAck,
		},
		Text: message,
	}
	res, err := w.wac.Send(textMessage)
	if err != nil {
		log.Panicln(err)
	}
	return res
}
