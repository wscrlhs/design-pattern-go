package composite

import "fmt"

// File 组件接口
type File struct {
	Name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.Name)
}

func (f *File) GetName() string {
	return f.Name
}

// Folder 组合
type Folder struct {
	components []Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, composite := range f.components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}

// Component 叶子
type Component interface {
	Search(string)
}
