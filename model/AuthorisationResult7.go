package model

// Outcome of the authorisation.
type AuthorisationResult7 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification75 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	TransactionResponse *ResponseType2 `xml:"TxRspn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Additional information relevant for the destination.
	AdditionalInformation []*ActionMessage3 `xml:"AddtlInf,omitempty"`
}

func (a *AuthorisationResult7) AddAuthorisationEntity() *GenericIdentification75 {
	a.AuthorisationEntity = new(GenericIdentification75)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult7) AddTransactionResponse() *ResponseType2 {
	a.TransactionResponse = new(ResponseType2)
	return a.TransactionResponse
}

func (a *AuthorisationResult7) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult7) AddAdditionalInformation() *ActionMessage3 {
	newValue := new(ActionMessage3)
	a.AdditionalInformation = append(a.AdditionalInformation, newValue)
	return newValue
}
