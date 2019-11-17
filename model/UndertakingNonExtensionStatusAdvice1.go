package model

// Details of a non-extension advice.
type UndertakingNonExtensionStatusAdvice1 struct {

	// Details related to the notifying party.
	NotifyingParty *PartyIdentification43 `xml:"NtifngPty"`

	// Details related to the identification of the undertaking.
	UndertakingIdentification *Undertaking7 `xml:"UdrtkgId"`
}

func (u *UndertakingNonExtensionStatusAdvice1) AddNotifyingParty() *PartyIdentification43 {
	u.NotifyingParty = new(PartyIdentification43)
	return u.NotifyingParty
}

func (u *UndertakingNonExtensionStatusAdvice1) AddUndertakingIdentification() *Undertaking7 {
	u.UndertakingIdentification = new(Undertaking7)
	return u.UndertakingIdentification
}
