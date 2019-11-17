package model

// Set of elements used to provide additional cover details for the claim non receipt.
type MissingCover2 struct {

	// Indicates whether or not the claim is related to a missing cover.
	MissingCoverIndicator *YesNoIndicator `xml:"MssngCoverInd"`

	// Set of elements provided to update incorrect settlement information for the cover related to the received payment instruction.
	CoverCorrection *SettlementInformation15 `xml:"CoverCrrctn,omitempty"`
}

func (m *MissingCover2) SetMissingCoverIndicator(value string) {
	m.MissingCoverIndicator = (*YesNoIndicator)(&value)
}

func (m *MissingCover2) AddCoverCorrection() *SettlementInformation15 {
	m.CoverCorrection = new(SettlementInformation15)
	return m.CoverCorrection
}
