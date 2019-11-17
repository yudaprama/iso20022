package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason12Choice struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason5Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AcknowledgementReason12Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason5Code)(&value)
}

func (a *AcknowledgementReason12Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}
