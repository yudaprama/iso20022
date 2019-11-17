package model

// Processing characteristics linked to the instrument, ie, not to  the market.
type Forms struct {

	// Physical application form is required.
	ApplicationForm *YesNoIndicator `xml:"ApplForm"`

	// Type of signature.
	SignatureType *SignatureType1Code `xml:"SgntrTp"`
}

func (f *Forms) SetApplicationForm(value string) {
	f.ApplicationForm = (*YesNoIndicator)(&value)
}

func (f *Forms) SetSignatureType(value string) {
	f.SignatureType = (*SignatureType1Code)(&value)
}
