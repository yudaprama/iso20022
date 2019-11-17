package model

// Choice of format for the modification cancellation information.
type ModificationCancellationAllowed4Choice struct {

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Modification, cancellation allowed information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (m *ModificationCancellationAllowed4Choice) SetIndicator(value string) {
	m.Indicator = (*YesNoIndicator)(&value)
}

func (m *ModificationCancellationAllowed4Choice) AddProprietary() *GenericIdentification30 {
	m.Proprietary = new(GenericIdentification30)
	return m.Proprietary
}
