package model

// Reason why the service is taxed.
type TaxReason1 struct {

	// Reason why the service is taxed, in a coded form.
	Code *Max10Text `xml:"Cd"`

	// Reason why the service is taxed, in a free-text form.
	Explanation *Max105Text `xml:"Expltn"`
}

func (t *TaxReason1) SetCode(value string) {
	t.Code = (*Max10Text)(&value)
}

func (t *TaxReason1) SetExplanation(value string) {
	t.Explanation = (*Max105Text)(&value)
}
