package slinterfaces

// main interface for the Simple Output
type ISimpleOutput interface {
	GetSessionID() string

	Open() (ISimpleChannel, error)
	Close() error

	GetDetails() string
}
