package chapter3

// Printer is interface
type Printer interface {
	open()
	print()
	close()
}

// AbstractDisplay is struct
type AbstractDisplay struct {
	printer Printer
}

// Display for printout result
func (a *AbstractDisplay) Display() {
	a.printer.open()
	for i := 0; i < 5; i++ {
		a.printer.print()
	}
	a.printer.close()
}
