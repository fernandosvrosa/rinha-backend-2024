package adapter

import (
	"github.com/gocql/gocql"
	"time"
)

type ContaLockAdapter struct {
	session *gocql.Session
}

func NewContaLockAdapter(session *gocql.Session) *ContaLockAdapter {
	return &ContaLockAdapter{session: session}
}

func (c ContaLockAdapter) Execute(clientID int) (bool, error) {
	var timestamp time.Time
	var applied bool
	query := "INSERT INTO conta_lock (id, time) VALUES (?, toTimestamp(now())) IF NOT EXISTS"

	applied, err := c.session.Query(query, clientID).ScanCAS(&clientID, &timestamp)

	if err != nil {
		return false, err

	}
	return applied, nil
}
