package bridge_test

import (
	"design-pattern/bridge"
	"fmt"
)

func Example() {
	hpPrinter := &bridge.Hp{}
	epsonPrinter := &bridge.Epson{}

	macComputer := &bridge.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &bridge.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()

	// Output:
	// Print request for mac
	// Printing by a HP Printer
	//
	// Print request for mac
	// Printing by a EPSON Printer
	//
	// Print request for windows
	// Printing by a HP Printer
	//
	// Print request for windows
	// Printing by a EPSON Printer
}
