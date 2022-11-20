package bridge

import "fmt"

type HpPrinter struct{}

func (h *HpPrinter) PrintFile() {
	fmt.Println("Printing file using HP printer")
}
