package model

// Additional information about a request (e.g. financing request).
type AdditionalInformation1 struct {

	// Specifies the type of additional information.
	InformationType *InformationType1Choice `xml:"InfTp"`

	// Contents of the additional information.
	InformationValue *Max350Text `xml:"InfVal"`
}

func (a *AdditionalInformation1) AddInformationType() *InformationType1Choice {
	a.InformationType = new(InformationType1Choice)
	return a.InformationType
}

func (a *AdditionalInformation1) SetInformationValue(value string) {
	a.InformationValue = (*Max350Text)(&value)
}
