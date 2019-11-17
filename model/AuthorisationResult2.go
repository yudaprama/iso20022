package model

// Outcome of the authorisation.
type AuthorisationResult2 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification33 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`
}

func (a *AuthorisationResult2) AddAuthorisationEntity() *GenericIdentification33 {
	a.AuthorisationEntity = new(GenericIdentification33)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult2) AddResponseToAuthorisation() *ResponseType1 {
	a.ResponseToAuthorisation = new(ResponseType1)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult2) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}
