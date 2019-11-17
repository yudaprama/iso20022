package model

// Information specific to an amendment.
type AmendInformation2 struct {

	// Identifies the MeetingResultDissemination essage to be amended.
	PreviousReference *MessageIdentification `xml:"PrvsRef"`
}

func (a *AmendInformation2) AddPreviousReference() *MessageIdentification {
	a.PreviousReference = new(MessageIdentification)
	return a.PreviousReference
}
