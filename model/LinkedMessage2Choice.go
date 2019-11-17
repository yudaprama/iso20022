package model

// Reference to a message.
type LinkedMessage2Choice struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef"`

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef"`
}

func (l *LinkedMessage2Choice) AddPreviousReference() *AdditionalReference3 {
	l.PreviousReference = new(AdditionalReference3)
	return l.PreviousReference
}

func (l *LinkedMessage2Choice) AddOtherReference() *AdditionalReference3 {
	l.OtherReference = new(AdditionalReference3)
	return l.OtherReference
}
