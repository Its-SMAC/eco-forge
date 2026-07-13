package registry

type Registry interface {
	Search(name string)
	Install(name string)
}
