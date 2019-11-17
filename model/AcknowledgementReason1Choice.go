package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason1Choice struct {

	// Specifies additional information about the processed instruction.
	Code *AcknowledgementReason3Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AcknowledgementReason1Choice) SetCode(value string) {
	a.Code = (*AcknowledgementReason3Code)(&value)
}

func (a *AcknowledgementReason1Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
