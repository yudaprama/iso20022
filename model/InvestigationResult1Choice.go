package model

// Provides the investigation results.
type InvestigationResult1Choice struct {

	// Provides the result.
	Result *SupplementaryDataEnvelope1 `xml:"Rslt"`

	// Provides the investigation status.
	InvestigationStatus *InvestigationStatus1Code `xml:"InvstgtnSts"`
}

func (i *InvestigationResult1Choice) AddResult() *SupplementaryDataEnvelope1 {
	i.Result = new(SupplementaryDataEnvelope1)
	return i.Result
}

func (i *InvestigationResult1Choice) SetInvestigationStatus(value string) {
	i.InvestigationStatus = (*InvestigationStatus1Code)(&value)
}
