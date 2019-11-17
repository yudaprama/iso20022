package model

// Outcome of the withdrawal authorisation.
type AuthorisationResult13 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *PartyType16Code `xml:"AuthstnNtty,omitempty"`

	// Result of the authorisation.
	AuthorisationResponse *ResponseType7 `xml:"AuthstnRspn"`

	// Trace of response by the entities in the path from the issuer to the ATM.
	ResponseTrace []*ResponseType8 `xml:"RspnTrac,omitempty"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action7 `xml:"Actn,omitempty"`
}

func (a *AuthorisationResult13) SetAuthorisationEntity(value string) {
	a.AuthorisationEntity = (*PartyType16Code)(&value)
}

func (a *AuthorisationResult13) AddAuthorisationResponse() *ResponseType7 {
	a.AuthorisationResponse = new(ResponseType7)
	return a.AuthorisationResponse
}

func (a *AuthorisationResult13) AddResponseTrace() *ResponseType8 {
	newValue := new(ResponseType8)
	a.ResponseTrace = append(a.ResponseTrace, newValue)
	return newValue
}

func (a *AuthorisationResult13) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult13) AddAction() *Action7 {
	newValue := new(Action7)
	a.Action = append(a.Action, newValue)
	return newValue
}
