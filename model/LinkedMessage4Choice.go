package model

// Reference to a message.
type LinkedMessage4Choice struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference6 `xml:"PrvsRef"`

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference *AdditionalReference6 `xml:"OthrRef"`
}

func (l *LinkedMessage4Choice) AddPreviousReference() *AdditionalReference6 {
	l.PreviousReference = new(AdditionalReference6)
	return l.PreviousReference
}

func (l *LinkedMessage4Choice) AddOtherReference() *AdditionalReference6 {
	l.OtherReference = new(AdditionalReference6)
	return l.OtherReference
}
