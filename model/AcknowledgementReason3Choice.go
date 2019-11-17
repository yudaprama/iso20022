package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason3Choice struct {

	// Specifies additional information about the processed instruction.
	Code *RepoCallAcknowledgementReason2Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AcknowledgementReason3Choice) SetCode(value string) {
	a.Code = (*RepoCallAcknowledgementReason2Code)(&value)
}

func (a *AcknowledgementReason3Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}
