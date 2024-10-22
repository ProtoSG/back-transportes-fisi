package application

import (
	domainShared "back-transportes-fisi/cmd/shared/domain"
	"back-transportes-fisi/cmd/terminal/domain"
	"back-transportes-fisi/cmd/terminal/domain/repository"
)

type TerminalUpdate struct {
	repository repository.TerminalRepository
}

func NewTerminalUpdate(repository repository.TerminalRepository) *TerminalUpdate {
	return &TerminalUpdate{repository}
}

func (this *TerminalUpdate) Execute(idTerminal int, nombre, ubicacion string) error {
	if terminal, _ := this.repository.GetById(idTerminal); terminal == nil {
		return domainShared.NewEntityNotFound(idTerminal, "Terminal")
	}

	terminal := domain.NewTerminal(idTerminal, nombre, ubicacion)
	return this.repository.Update(terminal)
}
