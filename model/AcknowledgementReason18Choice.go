package model

// Choice of format for the acknowledgement reason.
type AcknowledgementReason18Choice struct {

	// Specifies additional information about the processed instruction.
	Code *RepoCallAcknowledgementReason2Code `xml:"Cd"`

	// Specifies additional information about the processed instruction.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AcknowledgementReason18Choice) SetCode(value string) {
	a.Code = (*RepoCallAcknowledgementReason2Code)(&value)
}

func (a *AcknowledgementReason18Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
