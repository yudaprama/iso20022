package model

// Outcome of the authorisation.
type AuthorisationResult11 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification90 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType5 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`
}

func (a *AuthorisationResult11) AddAuthorisationEntity() *GenericIdentification90 {
	a.AuthorisationEntity = new(GenericIdentification90)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult11) AddResponseToAuthorisation() *ResponseType5 {
	a.ResponseToAuthorisation = new(ResponseType5)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult11) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}
