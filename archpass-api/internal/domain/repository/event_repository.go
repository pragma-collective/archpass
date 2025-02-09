package repository

import (
	"github.com/pragma-collective/archpass/ent"
	"github.com/pragma-collective/archpass/internal/domain/dto"
)

type IEventRepository interface {
	Create(input dto.CreateEventInput, userId int) (ent.Event, error)
	ListByOrganizerId(limit int, offset int, userId int) (ent.Events, error)
	GetBySlugAndOrganizerId(slug string, userId int) (ent.Event, error)
	GetBySlug(eventSlug string) (ent.Event, error)
	GetByEventHash(hash string) (ent.Event, error)
	GetById(id int) (ent.Event, error)
}
