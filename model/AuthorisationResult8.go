package model

// Outcome of the authorisation.
type AuthorisationResult8 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification75 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	TransactionResponse *ResponseType2 `xml:"TxRspn"`

	// Set of actions to be performed by the card acceptor.
	Action []*Action4 `xml:"Actn,omitempty"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Additional information relevant for the destination.
	AdditionalInformation []*ActionMessage3 `xml:"AddtlInf,omitempty"`
}

func (a *AuthorisationResult8) AddAuthorisationEntity() *GenericIdentification75 {
	a.AuthorisationEntity = new(GenericIdentification75)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult8) AddTransactionResponse() *ResponseType2 {
	a.TransactionResponse = new(ResponseType2)
	return a.TransactionResponse
}

func (a *AuthorisationResult8) AddAction() *Action4 {
	newValue := new(Action4)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *AuthorisationResult8) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult8) AddAdditionalInformation() *ActionMessage3 {
	newValue := new(ActionMessage3)
	a.AdditionalInformation = append(a.AdditionalInformation, newValue)
	return newValue
}
