package model

// Provides additional cover details for the claim non receipt.
type MissingCover3 struct {

	// Indicates whether or not the claim is related to a missing cover.
	MissingCoverIndicator *YesNoIndicator `xml:"MssngCoverInd"`

	// Set of elements provided to update incorrect settlement information for the cover related to the received payment instruction.
	CoverCorrection *SettlementInstruction3 `xml:"CoverCrrctn,omitempty"`
}

func (m *MissingCover3) SetMissingCoverIndicator(value string) {
	m.MissingCoverIndicator = (*YesNoIndicator)(&value)
}

func (m *MissingCover3) AddCoverCorrection() *SettlementInstruction3 {
	m.CoverCorrection = new(SettlementInstruction3)
	return m.CoverCorrection
}
