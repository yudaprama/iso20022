package model

// Specification of the EU capital gain.
type EUCapitalGain1 struct {

	// Structured format.
	Structured *EUCapitalGain1Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (e *EUCapitalGain1) SetStructured(value string) {
	e.Structured = (*EUCapitalGain1Code)(&value)
}

func (e *EUCapitalGain1) SetAdditionalInformation(value string) {
	e.AdditionalInformation = (*Max350Text)(&value)
}
