package infrastructure

import (
	"back-transportes-fisi/cmd/terminal/application"
	"back-transportes-fisi/cmd/terminal/infrastructure"
	"database/sql"
)

type TerminalService struct {
	Create  *application.TerminalCreate
	GetAll  *application.TerminalGetAll
	GetById *application.TerminalGetById
	Update  *application.TerminalUpdate
	Delete  *application.TerminalDelete
}

func NewTerminalService(db *sql.DB) TerminalService {
	terminalContianer := infrastructure.NewSQLiteTerminalRepository(db)

	return TerminalService{
		Create:  application.NewTerminalCreate(terminalContianer),
		GetAll:  application.NewTerminalGetAll(terminalContianer),
		GetById: application.NewTerminalGetById(terminalContianer),
		Update:  application.NewTerminalUpdate(terminalContianer),
		Delete:  application.NewTerminalDelete(terminalContianer),
	}
}
