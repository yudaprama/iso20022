package model

// Specifies the authority request type.
type AuthorityRequestType1 struct {

	// Specifies the requested message name identifier.
	MessageNameIdentification *Max35Text `xml:"MsgNmId"`

	// Specifies the message name.
	MessageName *Max140Text `xml:"MsgNm,omitempty"`
}

func (a *AuthorityRequestType1) SetMessageNameIdentification(value string) {
	a.MessageNameIdentification = (*Max35Text)(&value)
}

func (a *AuthorityRequestType1) SetMessageName(value string) {
	a.MessageName = (*Max140Text)(&value)
}
