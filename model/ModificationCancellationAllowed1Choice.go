package model

// Choice of format for the ModificationCancellation information.
type ModificationCancellationAllowed1Choice struct {

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Modification, cancellation allowed information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (m *ModificationCancellationAllowed1Choice) SetIndicator(value string) {
	m.Indicator = (*YesNoIndicator)(&value)
}

func (m *ModificationCancellationAllowed1Choice) AddProprietary() *GenericIdentification20 {
	m.Proprietary = new(GenericIdentification20)
	return m.Proprietary
}
