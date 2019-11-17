package model

// Outcome of the authorisation.
type AuthorisationResult5 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification70 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`
}

func (a *AuthorisationResult5) AddAuthorisationEntity() *GenericIdentification70 {
	a.AuthorisationEntity = new(GenericIdentification70)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult5) AddResponseToAuthorisation() *ResponseType1 {
	a.ResponseToAuthorisation = new(ResponseType1)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult5) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}
