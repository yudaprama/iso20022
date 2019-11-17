package model

// Outcome of the withdrawal authorisation.
type AuthorisationResult9 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *PartyType13Code `xml:"AuthstnNtty,omitempty"`

	// Result of the authorisation.
	AuthorisationResponse *ResponseType3 `xml:"AuthstnRspn"`

	// Trace of response by the entities in the path from the issuer to the ATM.
	ResponseTrace []*ResponseType4 `xml:"RspnTrac,omitempty"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action5 `xml:"Actn,omitempty"`
}

func (a *AuthorisationResult9) SetAuthorisationEntity(value string) {
	a.AuthorisationEntity = (*PartyType13Code)(&value)
}

func (a *AuthorisationResult9) AddAuthorisationResponse() *ResponseType3 {
	a.AuthorisationResponse = new(ResponseType3)
	return a.AuthorisationResponse
}

func (a *AuthorisationResult9) AddResponseTrace() *ResponseType4 {
	newValue := new(ResponseType4)
	a.ResponseTrace = append(a.ResponseTrace, newValue)
	return newValue
}

func (a *AuthorisationResult9) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult9) AddAction() *Action5 {
	newValue := new(Action5)
	a.Action = append(a.Action, newValue)
	return newValue
}
