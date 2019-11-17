package model

// Card acceptor performing the transaction.
type Organisation19 struct {

	// Identification of the card acceptor.
	Identification *GenericIdentification32 `xml:"Id"`

	// Name of the card acceptor as appearing on the receipt or the statement of account of the cardholder.
	// It correspond to the ISO 8583, field number 43.
	CommonName *Max70Text `xml:"CmonNm"`

	// Selected language of the card acceptor. Reference ISO 639-1 (alpha-2) and ISO 639-2 (alpha-3).
	SelectedLanguage *LanguageCode `xml:"SelctdLang,omitempty"`

	// Additional card acceptor data required by a card scheme.
	SchemeData *Max140Text `xml:"SchmeData,omitempty"`
}

func (o *Organisation19) AddIdentification() *GenericIdentification32 {
	o.Identification = new(GenericIdentification32)
	return o.Identification
}

func (o *Organisation19) SetCommonName(value string) {
	o.CommonName = (*Max70Text)(&value)
}

func (o *Organisation19) SetSelectedLanguage(value string) {
	o.SelectedLanguage = (*LanguageCode)(&value)
}

func (o *Organisation19) SetSchemeData(value string) {
	o.SchemeData = (*Max140Text)(&value)
}
