package model

// Expiry and extension information.
type ExpiryDetails1 struct {

	// Terms defining when the undertaking will cease to be available.
	ExpiryTerms *ExpiryTerms1 `xml:"XpryTerms,omitempty"`

	// Additional information related to the expiry and expiry extension.
	AdditionalExpiryInformation []*Max2000Text `xml:"AddtlXpryInf,omitempty"`
}

func (e *ExpiryDetails1) AddExpiryTerms() *ExpiryTerms1 {
	e.ExpiryTerms = new(ExpiryTerms1)
	return e.ExpiryTerms
}

func (e *ExpiryDetails1) AddAdditionalExpiryInformation(value string) {
	e.AdditionalExpiryInformation = append(e.AdditionalExpiryInformation, (*Max2000Text)(&value))
}
