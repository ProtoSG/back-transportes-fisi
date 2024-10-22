package application

import (
	"back-transportes-fisi/cmd/terminal/domain"
	"back-transportes-fisi/cmd/terminal/domain/repository"
)

type TerminalGetAll struct {
	repository repository.TerminalRepository
}

func NewTerminalGetAll(repository repository.TerminalRepository) *TerminalGetAll {
	return &TerminalGetAll{repository}
}

func (this *TerminalGetAll) Execute() ([]*domain.Terminal, error) {
	return this.repository.GetAll()
}
