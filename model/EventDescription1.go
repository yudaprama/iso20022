package model

// Describes an event not covered by other formal messages, for example a trace after a telephone call.
type EventDescription1 struct {

	// Identification of the event.
	Identifier *Max35Text `xml:"Idr"`

	// Date when event occurred.
	Date *ISODateTime `xml:"Dt,omitempty"`

	// Party to be advised.
	Recipient *QualifiedPartyIdentification1 `xml:"Rcpt"`

	// Advising party.
	Advisor *QualifiedPartyIdentification1 `xml:"Advsr"`

	// Parties involved in the event.
	OtherParty []*QualifiedPartyIdentification1 `xml:"OthrPty,omitempty"`

	// Identifier for a language used for the description.
	LanguageCode *LanguageCode `xml:"LangCd"`

	// Free form description of event.
	Description *Max2000Text `xml:"Desc"`

	// Reference to related document.
	RelatedDocument []*QualifiedDocumentInformation1 `xml:"RltdDoc,omitempty"`

	// Identifier of related letter.
	RelatedLetter []*QualifiedDocumentInformation1 `xml:"RltdLttr,omitempty"`

	// Identifier of related message.
	RelatedMessage []*QualifiedDocumentInformation1 `xml:"RltdMsg,omitempty"`

	// Associated free form document.
	AssociatedDocument []*QualifiedDocumentInformation1 `xml:"AssoctdDoc,omitempty"`

	// Reference to the contractual context.
	GoverningContract []*QualifiedDocumentInformation1 `xml:"GovngCtrct,omitempty"`

	// Rules and laws governing the event.
	LegalContext *GovernanceRules2 `xml:"LglCntxt,omitempty"`
}

func (e *EventDescription1) SetIdentifier(value string) {
	e.Identifier = (*Max35Text)(&value)
}

func (e *EventDescription1) SetDate(value string) {
	e.Date = (*ISODateTime)(&value)
}

func (e *EventDescription1) AddRecipient() *QualifiedPartyIdentification1 {
	e.Recipient = new(QualifiedPartyIdentification1)
	return e.Recipient
}

func (e *EventDescription1) AddAdvisor() *QualifiedPartyIdentification1 {
	e.Advisor = new(QualifiedPartyIdentification1)
	return e.Advisor
}

func (e *EventDescription1) AddOtherParty() *QualifiedPartyIdentification1 {
	newValue := new(QualifiedPartyIdentification1)
	e.OtherParty = append(e.OtherParty, newValue)
	return newValue
}

func (e *EventDescription1) SetLanguageCode(value string) {
	e.LanguageCode = (*LanguageCode)(&value)
}

func (e *EventDescription1) SetDescription(value string) {
	e.Description = (*Max2000Text)(&value)
}

func (e *EventDescription1) AddRelatedDocument() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	e.RelatedDocument = append(e.RelatedDocument, newValue)
	return newValue
}

func (e *EventDescription1) AddRelatedLetter() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	e.RelatedLetter = append(e.RelatedLetter, newValue)
	return newValue
}

func (e *EventDescription1) AddRelatedMessage() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	e.RelatedMessage = append(e.RelatedMessage, newValue)
	return newValue
}

func (e *EventDescription1) AddAssociatedDocument() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	e.AssociatedDocument = append(e.AssociatedDocument, newValue)
	return newValue
}

func (e *EventDescription1) AddGoverningContract() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	e.GoverningContract = append(e.GoverningContract, newValue)
	return newValue
}

func (e *EventDescription1) AddLegalContext() *GovernanceRules2 {
	e.LegalContext = new(GovernanceRules2)
	return e.LegalContext
}
