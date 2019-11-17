package model

// Details of the advice for the issuance of an undertaking.
type UndertakingAdvice2 struct {

	// Unique and unambiguous identifier assigned by the applicant to the undertaking.
	ApplicantReferenceNumber *Max35Text `xml:"ApplcntRefNb"`

	// Party obligated to reimburse the issuer.
	Obligor *PartyIdentification43 `xml:"Oblgr,omitempty"`

	// Contents of the related UndertakingIssuance message.
	UndertakingIssuanceMessage *UndertakingIssuanceMessage `xml:"UdrtkgIssncMsg"`

	// Medium used to issue the original undertaking.
	OriginalIssuedMedium *PresentationMedium1Code `xml:"OrgnlIssdMdm,omitempty"`

	// Document or template enclosed in the notification.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the undertaking notification.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingAdvice2) SetApplicantReferenceNumber(value string) {
	u.ApplicantReferenceNumber = (*Max35Text)(&value)
}

func (u *UndertakingAdvice2) AddObligor() *PartyIdentification43 {
	u.Obligor = new(PartyIdentification43)
	return u.Obligor
}

func (u *UndertakingAdvice2) AddUndertakingIssuanceMessage() *UndertakingIssuanceMessage {
	u.UndertakingIssuanceMessage = new(UndertakingIssuanceMessage)
	return u.UndertakingIssuanceMessage
}

func (u *UndertakingAdvice2) SetOriginalIssuedMedium(value string) {
	u.OriginalIssuedMedium = (*PresentationMedium1Code)(&value)
}

func (u *UndertakingAdvice2) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	u.EnclosedFile = append(u.EnclosedFile, newValue)
	return newValue
}

func (u *UndertakingAdvice2) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
