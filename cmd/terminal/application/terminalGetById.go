package application

import (
	domainShared "back-transportes-fisi/cmd/shared/domain"
	"back-transportes-fisi/cmd/terminal/domain"
	"back-transportes-fisi/cmd/terminal/domain/repository"
)

type TerminalGetById struct {
	repository repository.TerminalRepository
}

func NewTerminalGetById(repository repository.TerminalRepository) *TerminalGetById {
	return &TerminalGetById{repository}
}

func (this *TerminalGetById) Execute(idTerminal int) (*domain.Terminal, error) {
	terminal, _ := this.repository.GetById(idTerminal)
	if terminal == nil {
		return nil, domainShared.NewEntityNotFound(idTerminal, "Terminal")
	}

	return terminal, nil
}
