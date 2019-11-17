package model

// Information specific to an amendment or cancellation.
type AmendInformation1 struct {

	// Identifies the linked message which was previously sent.
	PreviousReference *MessageIdentification `xml:"PrvsRef"`

	// Indicates whether instructions must be resent (in case of modification of the parameters of a meeting for which instructions have already been sent).
	ReconfirmInstructions *YesNoIndicator `xml:"RcnfrmInstrs"`
}

func (a *AmendInformation1) AddPreviousReference() *MessageIdentification {
	a.PreviousReference = new(MessageIdentification)
	return a.PreviousReference
}

func (a *AmendInformation1) SetReconfirmInstructions(value string) {
	a.ReconfirmInstructions = (*YesNoIndicator)(&value)
}
