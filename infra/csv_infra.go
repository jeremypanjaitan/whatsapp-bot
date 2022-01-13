package infra

import (
	"encoding/csv"
	"os"
)

type CsvInfraEntity interface {
	ReadCsvFile(fileName string) ([][]string, error)
}

type CsvInfra struct{}

func NewCsvInfra() CsvInfraEntity {
	return &CsvInfra{}
}

func (c *CsvInfra) ReadCsvFile(fileName string) ([][]string, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return csv.NewReader(csvFile).ReadAll()
}
