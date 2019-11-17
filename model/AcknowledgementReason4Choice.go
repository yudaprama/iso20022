package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason4Choice struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason5Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AcknowledgementReason4Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason5Code)(&value)
}

func (a *AcknowledgementReason4Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
