package adapter

import (
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	"github.com/gocql/gocql"
)

type FindClientAdapter struct {
	session *gocql.Session
}

func NewFindClientAdapter(session *gocql.Session) *FindClientAdapter {
	return &FindClientAdapter{session: session}
}

func (f FindClientAdapter) Execute(clientID int) (entity.Client, error) {
	client := entity.Client{}

	if err := f.session.Query("SELECT id, limite, saldo_inicial, version FROM conta WHERE id = ?", clientID).Scan(&client.ID, &client.Limit, &client.Amount, &client.Version); err != nil {
		return entity.Client{}, err
	}

	return client, nil
}
