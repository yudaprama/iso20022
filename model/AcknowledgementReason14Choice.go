package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason14Choice struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason6Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AcknowledgementReason14Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason6Code)(&value)
}

func (a *AcknowledgementReason14Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}
