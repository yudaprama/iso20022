package model

// Indicates that the claim for non receipt is effectively a missing cover.
type MissingCover struct {

	// Indicates whether or not the claim is related to a missing cover.
	MissingCoverIndication *YesNoIndicator `xml:"MssngCoverIndctn"`
}

func (m *MissingCover) SetMissingCoverIndication(value string) {
	m.MissingCoverIndication = (*YesNoIndicator)(&value)
}
