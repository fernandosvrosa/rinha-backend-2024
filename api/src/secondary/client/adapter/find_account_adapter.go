package adapter

import (
	"github.com/fernandosvrosa/rinha-backend/api/src/domain/client/entity"
	"github.com/gocql/gocql"
)

type FindAccountAdapter struct {
	session *gocql.Session
}

func NewFindAccountAdapter(session *gocql.Session) *FindAccountAdapter {
	return &FindAccountAdapter{session: session}
}

func (f FindAccountAdapter) Execute(clientID int) (entity.Account, error) {
	client := entity.Account{}

	if err := f.session.Query("SELECT id, limite, saldo_inicial, version FROM conta WHERE id = ?", clientID).Scan(&client.ID, &client.Limit, &client.Amount, &client.Version); err != nil {
		return entity.Account{}, err
	}

	return client, nil
}
