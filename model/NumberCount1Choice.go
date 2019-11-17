package model

// Choice of number count type.
type NumberCount1Choice struct {

	// Sequential number of the instruction in a range of linked settlement instructions.
	CurrentInstructionNumber *Exact3NumericText `xml:"CurInstrNb"`

	// Total numbers of settlement transactions, receipts and deliveries, and the concerned settlement transaction number.
	TotalNumber *TotalNumber1 `xml:"TtlNb"`
}

func (n *NumberCount1Choice) SetCurrentInstructionNumber(value string) {
	n.CurrentInstructionNumber = (*Exact3NumericText)(&value)
}

func (n *NumberCount1Choice) AddTotalNumber() *TotalNumber1 {
	n.TotalNumber = new(TotalNumber1)
	return n.TotalNumber
}
