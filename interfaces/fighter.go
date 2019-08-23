package interfaces

type Fighter interface {
	Owner() Entity
	SetOwner(Entity)
}
