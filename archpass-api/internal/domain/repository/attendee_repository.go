package repository

import (
	"github.com/pragma-collective/archpass/internal/domain/dto"
)

type IAttendeeRepository interface {
	Create(input dto.CreateAttendeeInput) (int, error)
}
