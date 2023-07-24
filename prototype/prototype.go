package prototype

import "fmt"

// Inode 原型接口
type Inode interface {
	Print(string)
	Clone() Inode
}

// File
type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() Inode {
	return &File{Name: f.Name + "_Clone"}
}

// Folder 具体原型
type Folder struct {
	Children []Inode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Children {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	CloneFolder := &Folder{Name: f.Name + "_Clone"}
	var tempChildren []Inode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	CloneFolder.Children = tempChildren
	return CloneFolder
}
