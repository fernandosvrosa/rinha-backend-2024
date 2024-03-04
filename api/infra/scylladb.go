package infra

import (
	"github.com/gocql/gocql"
	"sync"
)

type ConnectionManager struct {
	Cluster *gocql.ClusterConfig
	Session *gocql.Session
	Lock    sync.Mutex
}

func (cm *ConnectionManager) Connect() error {
	cm.Lock.Lock()
	defer cm.Lock.Unlock()

	session, err := cm.Cluster.CreateSession()
	if err != nil {
		return err
	}

	cm.Session = session
	return nil
}

func (cm *ConnectionManager) Close() {
	cm.Lock.Lock()
	defer cm.Lock.Unlock()

	if cm.Session != nil {
		cm.Session.Close()
	}
}
