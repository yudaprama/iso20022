package model

// Information about the notification of the termination of an undertaking.
type UndertakingTerminationNotice1 struct {

	// Details related to the identification of the undertaking.
	UndertakingIdentification *Undertaking9 `xml:"UdrtkgId"`

	// Details related to the termination of the undertaking.
	TerminationDetails *UndertakingTermination3 `xml:"TermntnDtls"`

	// Document or template enclosed in the termination notification.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the notification.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingTerminationNotice1) AddUndertakingIdentification() *Undertaking9 {
	u.UndertakingIdentification = new(Undertaking9)
	return u.UndertakingIdentification
}

func (u *UndertakingTerminationNotice1) AddTerminationDetails() *UndertakingTermination3 {
	u.TerminationDetails = new(UndertakingTermination3)
	return u.TerminationDetails
}

func (u *UndertakingTerminationNotice1) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	u.EnclosedFile = append(u.EnclosedFile, newValue)
	return newValue
}

func (u *UndertakingTerminationNotice1) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
