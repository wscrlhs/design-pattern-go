package composite_test

import (
	"design-pattern/composite"
	"fmt"
)

func Example() {
	file1 := &composite.File{Name: "File1"}
	file2 := &composite.File{Name: "File2"}
	file3 := &composite.File{Name: "File3"}

	name1 := file1.GetName()
	if name1 != "File1" {
		fmt.Println("file1.GetName is wrong")
	}

	folder1 := &composite.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &composite.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")

	// Output:
	// Serching recursively for keyword rose in folder Folder2
	// Searching for keyword rose in file File2
	// Searching for keyword rose in file File3
	// Serching recursively for keyword rose in folder Folder1
	// Searching for keyword rose in file File1

}
