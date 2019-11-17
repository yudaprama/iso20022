package model

// Outcome of the authorisation.
type AuthorisationResult12 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification90 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType5 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Acquirer has requested a contact to the maintenance host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AuthorisationResult12) AddAuthorisationEntity() *GenericIdentification90 {
	a.AuthorisationEntity = new(GenericIdentification90)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult12) AddResponseToAuthorisation() *ResponseType5 {
	a.ResponseToAuthorisation = new(ResponseType5)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult12) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult12) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
