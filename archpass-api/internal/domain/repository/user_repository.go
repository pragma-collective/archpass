package repository

import (
	"github.com/pragma-collective/archpass/ent"
	"github.com/pragma-collective/archpass/internal/domain/dto"
)

type IUserRepository interface {
	Create(input dto.CreateUserInput) (ent.User, error)
	List() (ent.Users, error)
	FindByWalletAddress(walletAddress string) (ent.User, error)
}
