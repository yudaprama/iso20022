package model

// Card acceptor performing the transaction.
type Organisation18 struct {

	// Identification of the card acceptor.
	Identification *GenericIdentification32 `xml:"Id"`

	// Name of the card acceptor as appearing on the receipt or the statement of account of the cardholder.
	// It correspond to the ISO 8583 field number 43.
	CommonName *Max70Text `xml:"CmonNm"`

	// Location of the card acceptor.
	// It correspond to the ISO 8583 field number 43.
	Location *CommunicationAddress5 `xml:"Lctn"`

	// Selected language of the card acceptor. Reference ISO 639-1 (alpha-2) andISO 639-2 (alpha-3).
	SelectedLanguage *LanguageCode `xml:"SelctdLang,omitempty"`

	// Additional card acceptor data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation18) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation18) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation18) AddLocation() *CommunicationAddress5 {
	o.Location = new(CommunicationAddress5)
	return o.Location
}

func (o *Organisation18) SetSelectedLanguage(value string) {
	o.SelectedLanguage = (*LanguageCode)(&value)
}

func (o *Organisation18) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}
