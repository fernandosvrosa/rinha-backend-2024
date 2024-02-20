package port

import "github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"

type UpdateAmountClientPort interface {
	Execute(entity entity.Client) (entity.Client, error)
}
