package repo

import (
	"whatsapp-bot/config"
	"whatsapp-bot/infra"
	"whatsapp-bot/model"
)

type PhoneNumRepoEntity interface {
	ReadCsvFile() (*[]model.PhoneNumber, error)
}

type PhoneNumRepo struct {
	appConfig config.AppConfigEntity
	csvInfra  infra.CsvInfraEntity
}

func NewPhoneNumRepo(
	appConfig config.AppConfigEntity,
	csvInfra infra.CsvInfraEntity,
) PhoneNumRepoEntity {
	return &PhoneNumRepo{appConfig: appConfig, csvInfra: csvInfra}
}

func (c *PhoneNumRepo) ReadCsvFile() (*[]model.PhoneNumber, error) {
	csvLines, err := c.csvInfra.ReadCsvFile(c.appConfig.GetPhoneNumFileName())
	if err != nil {
		return nil, err
	}

	var phoneNumbers []model.PhoneNumber
	for _, line := range csvLines {
		phoneNumber := model.PhoneNumber{
			Number: line[0],
		}
		phoneNumbers = append(phoneNumbers, phoneNumber)
	}

	return &phoneNumbers, nil
}
