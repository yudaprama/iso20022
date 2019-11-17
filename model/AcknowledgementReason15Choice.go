package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason15Choice struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason3Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AcknowledgementReason15Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason3Code)(&value)
}

func (a *AcknowledgementReason15Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}
