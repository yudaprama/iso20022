package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason21Choice struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason6Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AcknowledgementReason21Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason6Code)(&value)
}

func (a *AcknowledgementReason21Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
