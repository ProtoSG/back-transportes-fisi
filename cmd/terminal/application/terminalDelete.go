package application

import (
	"back-transportes-fisi/cmd/shared/domain"
	"back-transportes-fisi/cmd/terminal/domain/repository"
)

type TerminalDelete struct {
	repository repository.TerminalRepository
}

func NewTerminalDelete(repository repository.TerminalRepository) *TerminalDelete {
	return &TerminalDelete{repository}
}

func (this *TerminalDelete) Execute(idTerminal int) error {
	if terminal, _ := this.repository.GetById(idTerminal); terminal == nil {
		return domain.NewEntityNotFound(idTerminal, "Terminal")
	}

	return this.repository.Delete(idTerminal)
}
