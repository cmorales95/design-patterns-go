package solid

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print(document Document) {}

func (m *MultiFunctionPrinter) Fax(document Document) {}

func (m *MultiFunctionPrinter) Scan(document Document) {}

// OldFashionPrinter does not make sense implement all the methods of the interface
type OldFashionPrinter struct{}

func (o OldFashionPrinter) Print(d Document) {
}

func (o OldFashionPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated: ...
func (o OldFashionPrinter) Scan(d Document) {
	panic("operation not supported")
}

// ISP

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// MyPrinter chooses the interface that it needs to work, no a big one with everything
type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {}

type Photocopier struct{}

func (p Photocopier) Scan(d Document) {}

func (p Photocopier) Print(d Document) {}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

// decorator design pattern
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func ExecuteISP() {

}