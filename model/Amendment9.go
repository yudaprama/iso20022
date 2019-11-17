package model

// Details of the amendment.
type Amendment9 struct {

	// Contents of the related UndertakingAmendmentResponse message.
	UndertakingAmendmentResponseMessage *UndertakingAmendmentResponseMessage1 `xml:"UdrtkgAmdmntRspnMsg"`
}

func (a *Amendment9) AddUndertakingAmendmentResponseMessage() *UndertakingAmendmentResponseMessage1 {
	a.UndertakingAmendmentResponseMessage = new(UndertakingAmendmentResponseMessage1)
	return a.UndertakingAmendmentResponseMessage
}
