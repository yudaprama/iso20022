package model

// Details of the non extension request.
type UndertakingNonExtensionRequest1 struct {

	// Details related to the requesting party.
	RequestingParty *PartyIdentification43 `xml:"RqstngPty"`

	// Details related to the identification of the undertaking.
	UndertakingIdentification *Undertaking9 `xml:"UdrtkgId"`
}

func (u *UndertakingNonExtensionRequest1) AddRequestingParty() *PartyIdentification43 {
	u.RequestingParty = new(PartyIdentification43)
	return u.RequestingParty
}

func (u *UndertakingNonExtensionRequest1) AddUndertakingIdentification() *Undertaking9 {
	u.UndertakingIdentification = new(Undertaking9)
	return u.UndertakingIdentification
}
