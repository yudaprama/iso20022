package model

// Type of tax expressed in structured or free text form.
type TaxTypeFormat2Choice struct {

	// Specifies the type of tax in free text form.
	Unstructured *Max35Text `xml:"Ustrd"`

	// Specifies the type of tax in structured form.
	Structured *TaxType2Code `xml:"Strd"`
}

func (t *TaxTypeFormat2Choice) SetUnstructured(value string) {
	t.Unstructured = (*Max35Text)(&value)
}

func (t *TaxTypeFormat2Choice) SetStructured(value string) {
	t.Structured = (*TaxType2Code)(&value)
}
