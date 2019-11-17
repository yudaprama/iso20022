package model

// Choice of format for the modification cancellation information.
type ModificationCancellationAllowed5Choice struct {

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Modification, cancellation allowed information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (m *ModificationCancellationAllowed5Choice) SetIndicator(value string) {
	m.Indicator = (*YesNoIndicator)(&value)
}

func (m *ModificationCancellationAllowed5Choice) AddProprietary() *GenericIdentification47 {
	m.Proprietary = new(GenericIdentification47)
	return m.Proprietary
}
