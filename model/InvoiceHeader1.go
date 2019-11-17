package model

// Collection of data for that is exchanged between two or more parties in written, printed or electronic form.
type InvoiceHeader1 struct {

	// Unique identification of this invoice document.
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of the document, for example commercial invoice.
	TypeCode *ExternalDocumentType1Code `xml:"TpCd"`

	// Name of invoice document or transaction, for example, tax invoice.
	Name []*Max35Text `xml:"Nm,omitempty"`

	// Issue date of the document.
	IssueDateTime *ISODateTime `xml:"IsseDtTm"`

	// Party that issued this invoice document.
	Issuer *TradeParty1 `xml:"Issr,omitempty"`

	// Unique identifier for a language used in this invoice document.
	LanguageCode *LanguageCode `xml:"LangCd,omitempty"`

	// Indicator that the document is a copy of the
	// original document.
	CopyIndicator *YesNoIndicator `xml:"CpyInd,omitempty"`

	// Specifies the function of the document.
	DocumentPurpose *ExternalDocumentPurpose1Code `xml:"DocPurp,omitempty"`

	// Note included in this invoice document.
	IncludedNote []*AdditionalInformation6 `xml:"InclNote,omitempty"`
}

func (i *InvoiceHeader1) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

func (i *InvoiceHeader1) SetTypeCode(value string) {
	i.TypeCode = (*ExternalDocumentType1Code)(&value)
}

func (i *InvoiceHeader1) AddName(value string) {
	i.Name = append(i.Name, (*Max35Text)(&value))
}

func (i *InvoiceHeader1) SetIssueDateTime(value string) {
	i.IssueDateTime = (*ISODateTime)(&value)
}

func (i *InvoiceHeader1) AddIssuer() *TradeParty1 {
	i.Issuer = new(TradeParty1)
	return i.Issuer
}

func (i *InvoiceHeader1) SetLanguageCode(value string) {
	i.LanguageCode = (*LanguageCode)(&value)
}

func (i *InvoiceHeader1) SetCopyIndicator(value string) {
	i.CopyIndicator = (*YesNoIndicator)(&value)
}

func (i *InvoiceHeader1) SetDocumentPurpose(value string) {
	i.DocumentPurpose = (*ExternalDocumentPurpose1Code)(&value)
}

func (i *InvoiceHeader1) AddIncludedNote() *AdditionalInformation6 {
	newValue := new(AdditionalInformation6)
	i.IncludedNote = append(i.IncludedNote, newValue)
	return newValue
}
