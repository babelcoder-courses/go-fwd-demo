package engine

type Engine interface {
	Dump() error
	Filter(string) ([]Resource, error)
}
