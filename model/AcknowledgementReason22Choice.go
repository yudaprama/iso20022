package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason22Choice struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason3Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AcknowledgementReason22Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason3Code)(&value)
}

func (a *AcknowledgementReason22Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
