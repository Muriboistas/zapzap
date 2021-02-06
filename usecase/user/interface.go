package user

import "github.com/muriboistas/zapzap/entity"

// Repository interface
type Repository interface {
	Reader
	Writer
}

// Reader interface
type Reader interface {
	Get(urn string) (*entity.User, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.User) (entity.ID, error)
	Update(e *entity.User) error
	Delete(id entity.ID) error
}

// UseCase interface
type UseCase interface {
	GetUser(urn string) (*entity.User, error)
	CreateUser(name, urn, language string) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(urn string) error
}
