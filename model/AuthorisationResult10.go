package model

// Outcome of the authorisation, and actions to perform.
type AuthorisationResult10 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification90 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType5 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Indicates whether the acquirer requires a further exchange completion after the completion of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Instructs the point of interaction (POI) how to contact the host to initiate the maintenance of the terminal.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AuthorisationResult10) AddAuthorisationEntity() *GenericIdentification90 {
	a.AuthorisationEntity = new(GenericIdentification90)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult10) AddResponseToAuthorisation() *ResponseType5 {
	a.ResponseToAuthorisation = new(ResponseType5)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult10) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult10) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *AuthorisationResult10) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}
