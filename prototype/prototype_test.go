package prototype_test

import (
	"design-pattern/prototype"
	"fmt"
)

func Example() {
	file1 := &prototype.File{Name: "File1"}
	file2 := &prototype.File{Name: "File2"}
	file3 := &prototype.File{Name: "File3"}

	folder1 := &prototype.Folder{
		Children: []prototype.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &prototype.Folder{
		Children: []prototype.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")

	// Output:
	// Printing hierarchy for Folder2
	//   Folder2
	//     Folder1
	//         File1
	//     File2
	//     File3
	//
	// Printing hierarchy for clone Folder
	//   Folder2_Clone
	//     Folder1_Clone
	//         File1_Clone
	//     File2_Clone
	//     File3_Clone
}
