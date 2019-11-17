package model

// Settlement transaction numbering information
type TotalNumber1 struct {

	// Sequential number of the instruction in a range of linked settlement instructions.
	CurrentInstructionNumber *Exact3NumericText `xml:"CurInstrNb"`

	// Total number of settlement instructions that are linked together.
	TotalOfLinkedInstructions *Exact3NumericText `xml:"TtlOfLkdInstrs"`
}

func (t *TotalNumber1) SetCurrentInstructionNumber(value string) {
	t.CurrentInstructionNumber = (*Exact3NumericText)(&value)
}

func (t *TotalNumber1) SetTotalOfLinkedInstructions(value string) {
	t.TotalOfLinkedInstructions = (*Exact3NumericText)(&value)
}
