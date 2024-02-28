package port

import "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"

type FindAccountPort interface {
	Execute(clientID int) (entity.Account, error)
}
