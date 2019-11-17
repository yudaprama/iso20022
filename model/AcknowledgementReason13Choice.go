package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason13Choice struct {

	// Specifies additional information about the processed instruction.
	Code *RepoCallAcknowledgementReason2Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AcknowledgementReason13Choice) SetCode(value string) {
	a.Code = (*RepoCallAcknowledgementReason2Code)(&value)
}

func (a *AcknowledgementReason13Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}
