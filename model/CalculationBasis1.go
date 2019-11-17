package model

// Specification of the charge basis.
type CalculationBasis1 struct {

	// Structured format.
	Structured *CalculationBasis1Code `xml:"Strd"`

	// Additional information about the calculation basis.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CalculationBasis1) SetStructured(value string) {
	c.Structured = (*CalculationBasis1Code)(&value)
}

func (c *CalculationBasis1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
