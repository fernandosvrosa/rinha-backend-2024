package port

type ContaLockPort interface {
	Execute(clientID int) (bool, error)
}
