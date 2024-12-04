package closer

type Closer interface {
	AddCloser(closer func())
}

type closer struct {
	closers []func()
}

func New() *closer {
	return &closer{}
}

func (c *closer) AddCloser(closer func()) {
	c.closers = append(c.closers, closer)
}

func (c *closer) Close() {
	for _, closer := range c.closers {
		if err := closer(); err != nil {
			return
		}
	}

	return
}