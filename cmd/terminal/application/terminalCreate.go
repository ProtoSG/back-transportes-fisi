package application

import (
	"back-transportes-fisi/cmd/terminal/domain"
	"back-transportes-fisi/cmd/terminal/domain/repository"
)

type TerminalCreate struct {
	repository repository.TerminalRepository
}

func NewTerminalCreate(repository repository.TerminalRepository) *TerminalCreate {
	return &TerminalCreate{repository}
}

func (this *TerminalCreate) Execute(idTerminal int, nombre string, ubicacion string) error {
	terminal := domain.NewTerminal(idTerminal, nombre, ubicacion)
	return this.repository.Create(terminal)
}
