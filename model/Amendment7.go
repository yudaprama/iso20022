package model

// Details of the amendment.
type Amendment7 struct {

	// Identification of the proposed amendment.
	AmendmentIdentification *Amendment8 `xml:"AmdmntId"`

	// Proposed undertaking amendment status.
	AmendmentStatus *UndertakingStatus2Code `xml:"AmdmntSts"`
}

func (a *Amendment7) AddAmendmentIdentification() *Amendment8 {
	a.AmendmentIdentification = new(Amendment8)
	return a.AmendmentIdentification
}

func (a *Amendment7) SetAmendmentStatus(value string) {
	a.AmendmentStatus = (*UndertakingStatus2Code)(&value)
}
