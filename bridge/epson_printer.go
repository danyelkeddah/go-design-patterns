package bridge

import "fmt"

type EpsonPrinter struct{}

func (e *EpsonPrinter) PrintFile() {
	fmt.Println("Printing file using EPSON printer")
}
