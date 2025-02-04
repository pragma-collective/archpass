package repository

import (
	"github.com/pragma-collective/archpass/internal/domain/dto"
)

type ITransactionRepository interface {
	Create(input dto.CreateTransactionInput) (int, error)
}
