package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason9Choice struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason6Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AcknowledgementReason9Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason6Code)(&value)
}

func (a *AcknowledgementReason9Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
