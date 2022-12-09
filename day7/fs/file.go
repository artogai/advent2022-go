package fs

type file struct {
	name   string
	size   int
	parent *Directory
}

func NewFile(name string, size int) *file {
	return &file{name, size, nil}
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Size() int {
	return f.size
}

func (f *file) SetParent(p *Directory) {
	f.parent = p
}
