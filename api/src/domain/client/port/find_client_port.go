package port

import "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"

type FindClientPort interface {
	Execute(clientID int) (entity.Client, error)
}
