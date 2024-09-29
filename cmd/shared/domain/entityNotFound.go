package domain

import "fmt"

type EntityNotFound struct {
	Id   int
	Name string
}

func NewEntityNotFound(id int, name string) EntityNotFound {
	return EntityNotFound{
		Id:   id,
		Name: name,
	}
}

func (this EntityNotFound) Error() string {
	return fmt.Sprintf("%s con Id %d no encontrada", this.Name, this.Id)
}
