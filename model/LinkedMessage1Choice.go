package model

// Reference to a message.
type LinkedMessage1Choice struct {

	// Linked previous reference.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef"`

	// Linked other reference.
	OtherReference *AdditionalReference3 `xml:"OthrRef"`

	// Linked related reference.
	RelatedReference *AdditionalReference3 `xml:"RltdRef"`
}

func (l *LinkedMessage1Choice) AddPreviousReference() *AdditionalReference3 {
	l.PreviousReference = new(AdditionalReference3)
	return l.PreviousReference
}

func (l *LinkedMessage1Choice) AddOtherReference() *AdditionalReference3 {
	l.OtherReference = new(AdditionalReference3)
	return l.OtherReference
}

func (l *LinkedMessage1Choice) AddRelatedReference() *AdditionalReference3 {
	l.RelatedReference = new(AdditionalReference3)
	return l.RelatedReference
}
