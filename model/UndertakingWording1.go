package model

// Information about the wording for a demand guarantee, standby letter of credit or other undertaking.
type UndertakingWording1 struct {

	// Wording template for the undertaking content made available for use with certain governance rules or made available by particular institutions.
	ModelForm *ModelFormIdentification1 `xml:"MdlForm,omitempty"`

	// Language of the standard wording provided by the issuer.
	RequestedWordingLanguage *ISO2ALanguageCode `xml:"ReqdWrdgLang,omitempty"`

	// Terms and conditions of the undertaking.
	UndertakingTermsAndConditions []*Narrative1 `xml:"UdrtkgTermsAndConds,omitempty"`
}

func (u *UndertakingWording1) AddModelForm() *ModelFormIdentification1 {
	u.ModelForm = new(ModelFormIdentification1)
	return u.ModelForm
}

func (u *UndertakingWording1) SetRequestedWordingLanguage(value string) {
	u.RequestedWordingLanguage = (*ISO2ALanguageCode)(&value)
}

func (u *UndertakingWording1) AddUndertakingTermsAndConditions() *Narrative1 {
	newValue := new(Narrative1)
	u.UndertakingTermsAndConditions = append(u.UndertakingTermsAndConditions, newValue)
	return newValue
}
