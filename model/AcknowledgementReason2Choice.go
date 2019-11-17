package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason2Choice struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason5Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AcknowledgementReason2Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason5Code)(&value)
}

func (a *AcknowledgementReason2Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
