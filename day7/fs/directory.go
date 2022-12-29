package fs

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Directory struct {
	name    string
	parent  *Directory
	Content []Entry
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{name, parent, []Entry{}}
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int {
	size := 0
	for _, entry := range d.Content {
		size += entry.Size()
	}
	return size
}

func (d *Directory) Parent() *Directory {
	return d.parent
}

func (d *Directory) SetParent(p *Directory) {
	d.parent = p
}

func (d *Directory) SubDir(name string) *Directory {
	idx := slices.IndexFunc(d.Content, func(e Entry) bool {
		return e.Name() == name
	})
	if idx != -1 {
		subDir, ok := (d.Content[idx]).(*Directory)
		if ok {
			return subDir
		}
	}
	panic(fmt.Sprintf("SubDir: subdirectory with %s not found", name))
}

func (d *Directory) SubDirs() []*Directory {
	dirs := []*Directory{}

	var rec func(d *Directory)
	rec = func(d *Directory) {
		subDirs := d.subDirs()
		dirs = append(dirs, subDirs...)
		for _, subDir := range subDirs {
			rec(subDir)
		}
	}
	rec(d)

	return dirs
}

func (d *Directory) subDirs() []*Directory {
	dirs := []*Directory{}
	for _, e := range d.Content {
		dir, ok := e.(*Directory)
		if ok {
			dirs = append(dirs, dir)
		}
	}
	return dirs
}
