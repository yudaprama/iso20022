package model

// Specifies the type of tax as a code or free text.
type TaxType2Choice struct {

	// Type of tax applied.
	Type *TaxType9Code `xml:"Tp"`

	// Specifies types of tax not present in a code list.
	OtherTaxType *Max35Text `xml:"OthrTaxTp"`
}

func (t *TaxType2Choice) SetType(value string) {
	t.Type = (*TaxType9Code)(&value)
}

func (t *TaxType2Choice) SetOtherTaxType(value string) {
	t.OtherTaxType = (*Max35Text)(&value)
}
