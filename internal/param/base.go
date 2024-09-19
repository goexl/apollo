package param

type Base struct {
	Cluster    string
	Namespaces []string
}

func newBase() *Base {
	return &Base{
		Cluster:    "default",
		Namespaces: make([]string, 0, 2),
	}
}
