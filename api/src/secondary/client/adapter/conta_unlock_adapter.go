package adapter

import "github.com/gocql/gocql"

type ContaUnLockAdapter struct {
	session *gocql.Session
}

func NewContaUnLockAdapter(session *gocql.Session) *ContaUnLockAdapter {
	return &ContaUnLockAdapter{session: session}
}

func (c ContaUnLockAdapter) Execute(clientID int) (bool, error) {
	query := "DELETE FROM conta_lock WHERE id = ?"

	if err := c.session.Query(query, clientID).Exec(); err != nil {
		return false, err
	}

	return true, nil
}
