package docker

type TagStore interface {
	Name() string
	Init(opts ...Option) (err error)
	TagInset()
	TagUpdate(tag string)
	TagFind() (tag string)
}
