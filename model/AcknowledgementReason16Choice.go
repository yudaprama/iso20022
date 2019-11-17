package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason16Choice struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason5Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AcknowledgementReason16Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason5Code)(&value)
}

func (a *AcknowledgementReason16Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
