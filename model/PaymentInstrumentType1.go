package model

// Identifies a payment instrument type as the search criteria for the financial institution to do the investigation.
type PaymentInstrumentType1 struct {

	// Provides the card number.
	CardNumber *Min8Max28NumericText `xml:"CardNb"`

	// Identifies the authority request type as a code.
	AuthorityRequestType []*AuthorityRequestType1 `xml:"AuthrtyReqTp"`

	// Additional information, in free text form, to complement the requested information.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`
}

func (p *PaymentInstrumentType1) SetCardNumber(value string) {
	p.CardNumber = (*Min8Max28NumericText)(&value)
}

func (p *PaymentInstrumentType1) AddAuthorityRequestType() *AuthorityRequestType1 {
	newValue := new(AuthorityRequestType1)
	p.AuthorityRequestType = append(p.AuthorityRequestType, newValue)
	return newValue
}

func (p *PaymentInstrumentType1) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max500Text)(&value)
}
