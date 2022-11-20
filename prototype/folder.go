package prototype

import "fmt"

type Folder struct {
	Childrens []Inode
	Name      string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Childrens {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{
		Name: f.Name,
	}
	var tempChildrens []Inode
	for _, i := range f.Childrens {
		c := i.Clone()
		tempChildrens = append(tempChildrens, c)
	}
	cloneFolder.Childrens = tempChildrens
	return cloneFolder
}
