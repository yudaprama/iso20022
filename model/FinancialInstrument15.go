package model

// Identifies the financial instrument.
type FinancialInstrument15 struct {

	// Identifies the financial instrument using a choice of either ISIN,  local code, or a description of the instrument. ISIN is the preferred format.
	Identification *SecurityIdentification6Choice `xml:"Id"`

	// Provides the ability to describe the instrument through a description and main characteristics.
	InstrumentDescription *SecurityInstrumentDescription2 `xml:"InstrmDesc,omitempty"`

	// Provides details of the underlying financial instrument for which the transaction report is being sent. If there is more than one underlying financial instrument then it is the dominant/ultimate instrument that should be identified here.
	UnderlyingInstrumentIdentification *SecurityIdentification6Choice `xml:"UndrlygInstrmId,omitempty"`
}

func (f *FinancialInstrument15) AddIdentification() *SecurityIdentification6Choice {
	f.Identification = new(SecurityIdentification6Choice)
	return f.Identification
}

func (f *FinancialInstrument15) AddInstrumentDescription() *SecurityInstrumentDescription2 {
	f.InstrumentDescription = new(SecurityInstrumentDescription2)
	return f.InstrumentDescription
}

func (f *FinancialInstrument15) AddUnderlyingInstrumentIdentification() *SecurityIdentification6Choice {
	f.UnderlyingInstrumentIdentification = new(SecurityIdentification6Choice)
	return f.UnderlyingInstrumentIdentification
}
