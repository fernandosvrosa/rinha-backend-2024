package client

import (
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	appError "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/error"
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/port"
)

type FindAccountByIdUsecase struct {
	findAccountPort port.FindAccountPort
}

func NewFindAccountByIdUsecase(findAccountPort port.FindAccountPort) *FindAccountByIdUsecase {
	return &FindAccountByIdUsecase{
		findAccountPort: findAccountPort,
	}
}

func (f *FindAccountByIdUsecase) Execute(id int) (entity.Account, error) {

	account, err := f.findAccountPort.Execute(id)

	if err != nil {
		if err.Error() == "not found" {
			return entity.Account{}, appError.NotFound{Message: "Account not found."}

		}
		return entity.Account{}, err
	}

	return account, nil
}
