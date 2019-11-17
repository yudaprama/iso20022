package model

// Additional information about a request.
type AdditionalInformation6 struct {

	// Specifies the type of additional information.
	InformationType *ExternalInformationType1Code `xml:"InfTp"`

	// Contents of the additional information.
	InformationValue *Max350Text `xml:"InfVal"`
}

func (a *AdditionalInformation6) SetInformationType(value string) {
	a.InformationType = (*ExternalInformationType1Code)(&value)
}

func (a *AdditionalInformation6) SetInformationValue(value string) {
	a.InformationValue = (*Max350Text)(&value)
}
