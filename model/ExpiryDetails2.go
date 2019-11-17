package model

// Expiry and extension information.
type ExpiryDetails2 struct {

	// Terms defining when the undertaking will cease to be available.
	ExpiryTerms *ExpiryTerms2 `xml:"XpryTerms,omitempty"`

	// Additional information related to the expiry and expiry extension.
	AdditionalExpiryInformation []*Max2000Text `xml:"AddtlXpryInf,omitempty"`
}

func (e *ExpiryDetails2) AddExpiryTerms() *ExpiryTerms2 {
	e.ExpiryTerms = new(ExpiryTerms2)
	return e.ExpiryTerms
}

func (e *ExpiryDetails2) AddAdditionalExpiryInformation(value string) {
	e.AdditionalExpiryInformation = append(e.AdditionalExpiryInformation, (*Max2000Text)(&value))
}
