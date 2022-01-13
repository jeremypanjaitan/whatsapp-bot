package infra

import "os"

type TextInfraEntity interface {
	ReadFile(fileName string) (string, error)
}

type TextInfra struct {
}

func NewTextInfra() TextInfraEntity {
	return &TextInfra{}
}

func (t *TextInfra) ReadFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	return string(content), err
}
