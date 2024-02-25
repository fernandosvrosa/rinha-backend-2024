package adapter

import (
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	"github.com/gocql/gocql"
)

type UpdateAmountClientAdapter struct {
	session *gocql.Session
}

func NewUpdateAmountClientAdapter(session *gocql.Session) *UpdateAmountClientAdapter {
	return &UpdateAmountClientAdapter{session: session}
}

func (u UpdateAmountClientAdapter) Execute(client entity.Client) (entity.Client, error) {
	newVersion := client.Version + 1
	if err := u.session.Query("UPDATE conta SET saldo_inicial = ?, version = ? WHERE id = ?;", client.Amount, newVersion, client.ID).Exec(); err != nil {
		return entity.Client{}, err
	}
	return client, nil
}
