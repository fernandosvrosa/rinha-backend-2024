package port

type ContaUnLockPort interface {
	Execute(clientID int) (bool, error)
}
