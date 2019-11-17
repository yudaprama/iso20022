package model

// Defines a business letter containing identifications of involved entities and their roles, references to documents, free form text and signatures.
// The semantics of this information are defined by usual business practices for the exchange and tracing of business letters. The described references and party identifiers permit to establish a linked informal trace of sequences of letters.
// This message component contains three types of elements that can be referenced using IDREF:
// (1) - all elements defining qualified parties,
// (2) - all elements defining qualified documents or references to them,
// (3) - the LegalContext element.
type BusinessLetter1 struct {

	// Application context defined by users. This is typically the name of a product.
	ApplicationContext *Max35Text `xml:"ApplCntxt,omitempty"`

	// Unambiguous identifier for this letter.
	LetterIdentifier *QualifiedDocumentInformation1 `xml:"LttrIdr"`

	// Purported creation date of the document.
	Date *ISODate `xml:"Dt"`

	// Identifier of a related letter.
	RelatedLetter []*QualifiedDocumentInformation1 `xml:"RltdLttr,omitempty"`

	// Identifier of a related message.
	RelatedMessage []*QualifiedDocumentInformation1 `xml:"RltdMsg,omitempty"`

	// Cross references the lists that are associated to this letter inside a message. The identifiers are relative to the Originator.
	ContentIdentifier []*Max35Text `xml:"CnttIdr,omitempty"`

	// Urgency or order of importance that the originator would like the recipient of the business letter to apply to the processing of the letter.
	InstructionPriority *Priority3Code `xml:"InstrPrty,omitempty"`

	// Identification of the originating party of this letter.
	Originator *QualifiedPartyIdentification1 `xml:"Orgtr"`

	// Primary recipient of the business letter. The exact meaning is given by the users.
	PrimaryRecipient []*QualifiedPartyIdentification1 `xml:"PmryRcpt"`

	// Sender of the business letter. The exact meaning is given by the users.
	Sender []*QualifiedPartyIdentification1 `xml:"Sndr,omitempty"`

	// User who, either individually or in concert with others, authorises the origination of a message.
	AuthorisationUser []*QualifiedPartyIdentification1 `xml:"AuthstnUsr"`

	// Party to receive a reply to this letter.
	ResponseRecipient []*QualifiedPartyIdentification1 `xml:"RspnRcpt,omitempty"`

	// Party to receive a copy of the message.
	CopyRecipient []*QualifiedPartyIdentification1 `xml:"CpyRcpt,omitempty"`

	// Other party involved. This element is usable as a target for IDREFs.
	OtherParty []*QualifiedPartyIdentification1 `xml:"OthrPty,omitempty"`

	// Associated free form document.
	AssociatedDocument []*QualifiedDocumentInformation1 `xml:"AssoctdDoc,omitempty"`

	// Governing contract.
	GoverningContract []*QualifiedDocumentInformation1 `xml:"GovngCtrct,omitempty"`

	// Rules and laws governing the letter.
	LegalContext []*GovernanceRules2 `xml:"LglCntxt,omitempty"`

	// Free form information about this message.
	AdditionalInformation *Max2000Text `xml:"AddtlInf,omitempty"`

	// Free form information unrelated to the message for example advertising or a service notice.
	Notice *Max350Text `xml:"Ntce,omitempty"`

	// Status of referenced messages or letters.
	ValidationStatusInformation *ValidationStatusInformation1 `xml:"VldtnStsInf,omitempty"`

	// Digital signatures and signing parties of this letter or parts of it.
	DigitalSignature []*QualifiedPartyAndXMLSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (b *BusinessLetter1) SetApplicationContext(value string) {
	b.ApplicationContext = (*Max35Text)(&value)
}

func (b *BusinessLetter1) AddLetterIdentifier() *QualifiedDocumentInformation1 {
	b.LetterIdentifier = new(QualifiedDocumentInformation1)
	return b.LetterIdentifier
}

func (b *BusinessLetter1) SetDate(value string) {
	b.Date = (*ISODate)(&value)
}

func (b *BusinessLetter1) AddRelatedLetter() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	b.RelatedLetter = append(b.RelatedLetter, newValue)
	return newValue
}

func (b *BusinessLetter1) AddRelatedMessage() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	b.RelatedMessage = append(b.RelatedMessage, newValue)
	return newValue
}

func (b *BusinessLetter1) AddContentIdentifier(value string) {
	b.ContentIdentifier = append(b.ContentIdentifier, (*Max35Text)(&value))
}

func (b *BusinessLetter1) SetInstructionPriority(value string) {
	b.InstructionPriority = (*Priority3Code)(&value)
}

func (b *BusinessLetter1) AddOriginator() *QualifiedPartyIdentification1 {
	b.Originator = new(QualifiedPartyIdentification1)
	return b.Originator
}

func (b *BusinessLetter1) AddPrimaryRecipient() *QualifiedPartyIdentification1 {
	newValue := new(QualifiedPartyIdentification1)
	b.PrimaryRecipient = append(b.PrimaryRecipient, newValue)
	return newValue
}

func (b *BusinessLetter1) AddSender() *QualifiedPartyIdentification1 {
	newValue := new(QualifiedPartyIdentification1)
	b.Sender = append(b.Sender, newValue)
	return newValue
}

func (b *BusinessLetter1) AddAuthorisationUser() *QualifiedPartyIdentification1 {
	newValue := new(QualifiedPartyIdentification1)
	b.AuthorisationUser = append(b.AuthorisationUser, newValue)
	return newValue
}

func (b *BusinessLetter1) AddResponseRecipient() *QualifiedPartyIdentification1 {
	newValue := new(QualifiedPartyIdentification1)
	b.ResponseRecipient = append(b.ResponseRecipient, newValue)
	return newValue
}

func (b *BusinessLetter1) AddCopyRecipient() *QualifiedPartyIdentification1 {
	newValue := new(QualifiedPartyIdentification1)
	b.CopyRecipient = append(b.CopyRecipient, newValue)
	return newValue
}

func (b *BusinessLetter1) AddOtherParty() *QualifiedPartyIdentification1 {
	newValue := new(QualifiedPartyIdentification1)
	b.OtherParty = append(b.OtherParty, newValue)
	return newValue
}

func (b *BusinessLetter1) AddAssociatedDocument() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	b.AssociatedDocument = append(b.AssociatedDocument, newValue)
	return newValue
}

func (b *BusinessLetter1) AddGoverningContract() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	b.GoverningContract = append(b.GoverningContract, newValue)
	return newValue
}

func (b *BusinessLetter1) AddLegalContext() *GovernanceRules2 {
	newValue := new(GovernanceRules2)
	b.LegalContext = append(b.LegalContext, newValue)
	return newValue
}

func (b *BusinessLetter1) SetAdditionalInformation(value string) {
	b.AdditionalInformation = (*Max2000Text)(&value)
}

func (b *BusinessLetter1) SetNotice(value string) {
	b.Notice = (*Max350Text)(&value)
}

func (b *BusinessLetter1) AddValidationStatusInformation() *ValidationStatusInformation1 {
	b.ValidationStatusInformation = new(ValidationStatusInformation1)
	return b.ValidationStatusInformation
}

func (b *BusinessLetter1) AddDigitalSignature() *QualifiedPartyAndXMLSignature1 {
	newValue := new(QualifiedPartyAndXMLSignature1)
	b.DigitalSignature = append(b.DigitalSignature, newValue)
	return newValue
}
