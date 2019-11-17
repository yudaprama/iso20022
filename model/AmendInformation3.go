package model

// Information specific to an amendment.
type AmendInformation3 struct {

	// Identifies the MeetingResultDissemination message to be amended.
	PreviousReference *MessageIdentification `xml:"PrvsRef"`
}

func (a *AmendInformation3) AddPreviousReference() *MessageIdentification {
	a.PreviousReference = new(MessageIdentification)
	return a.PreviousReference
}
