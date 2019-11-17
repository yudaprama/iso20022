package model

// Choice of format for the ModificationCancellation information.
type ModificationCancellationAllowed3Choice struct {

	// Modification, cancellation allowed information expressed as a indicator.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Modification, cancellation allowed information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (m *ModificationCancellationAllowed3Choice) SetIndicator(value string) {
	m.Indicator = (*YesNoIndicator)(&value)
}

func (m *ModificationCancellationAllowed3Choice) AddProprietary() *GenericIdentification38 {
	m.Proprietary = new(GenericIdentification38)
	return m.Proprietary
}
