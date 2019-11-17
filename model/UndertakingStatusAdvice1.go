package model

// Undertaking status information.
type UndertakingStatusAdvice1 struct {

	// Details related to the party that initiated the report.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty"`

	// Details related to the identification of the undertaking.
	UndertakingIdentification *Undertaking8 `xml:"UdrtkgId,omitempty"`

	// Sequence number assigned by the issuer to each amendment of the undertaking.
	AmendmentSequenceNumber *Number `xml:"AmdmntSeqNb,omitempty"`

	// Unique and unambiguous identifier assigned by the advising party to the undertaking.
	AdvisingPartyReferenceNumber *Max35Text `xml:"AdvsgPtyRefNb,omitempty"`

	// Unique and unambiguous identifier assigned by the confirmer to the undertaking.
	ConfirmerReferenceNumber *Max35Text `xml:"CnfrmrRefNb,omitempty"`

	// Specifies the category of the status.
	StatusCategory *ExternalUndertakingStatusCategory1Code `xml:"StsCtgy"`

	// Specifies the reported status.
	Status *UndertakingStatus3Code `xml:"Sts"`

	// Set of elements used to provide detailed information on the status reason.
	StatusReason []*StatusReasonInformation8 `xml:"StsRsn,omitempty"`

	// Amount reported.
	ReportedAmount []*ReportedAmount1 `xml:"RptdAmt,omitempty"`

	// Information concerning the original message to which the status report may be sent in response.
	OriginalMessageDetails *OriginalMessage1 `xml:"OrgnlMsgDtls,omitempty"`

	// Document or template enclosed in the report.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the report.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingStatusAdvice1) AddInitiatingParty() *PartyIdentification43 {
	u.InitiatingParty = new(PartyIdentification43)
	return u.InitiatingParty
}

func (u *UndertakingStatusAdvice1) AddUndertakingIdentification() *Undertaking8 {
	u.UndertakingIdentification = new(Undertaking8)
	return u.UndertakingIdentification
}

func (u *UndertakingStatusAdvice1) SetAmendmentSequenceNumber(value string) {
	u.AmendmentSequenceNumber = (*Number)(&value)
}

func (u *UndertakingStatusAdvice1) SetAdvisingPartyReferenceNumber(value string) {
	u.AdvisingPartyReferenceNumber = (*Max35Text)(&value)
}

func (u *UndertakingStatusAdvice1) SetConfirmerReferenceNumber(value string) {
	u.ConfirmerReferenceNumber = (*Max35Text)(&value)
}

func (u *UndertakingStatusAdvice1) SetStatusCategory(value string) {
	u.StatusCategory = (*ExternalUndertakingStatusCategory1Code)(&value)
}

func (u *UndertakingStatusAdvice1) SetStatus(value string) {
	u.Status = (*UndertakingStatus3Code)(&value)
}

func (u *UndertakingStatusAdvice1) AddStatusReason() *StatusReasonInformation8 {
	newValue := new(StatusReasonInformation8)
	u.StatusReason = append(u.StatusReason, newValue)
	return newValue
}

func (u *UndertakingStatusAdvice1) AddReportedAmount() *ReportedAmount1 {
	newValue := new(ReportedAmount1)
	u.ReportedAmount = append(u.ReportedAmount, newValue)
	return newValue
}

func (u *UndertakingStatusAdvice1) AddOriginalMessageDetails() *OriginalMessage1 {
	u.OriginalMessageDetails = new(OriginalMessage1)
	return u.OriginalMessageDetails
}

func (u *UndertakingStatusAdvice1) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	u.EnclosedFile = append(u.EnclosedFile, newValue)
	return newValue
}

func (u *UndertakingStatusAdvice1) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
