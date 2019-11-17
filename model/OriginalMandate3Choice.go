package model

// Specifies the mandate that is being accepted.
type OriginalMandate3Choice struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId"`

	// Provides the original mandate data.
	OriginalMandate *Mandate5 `xml:"OrgnlMndt"`
}

func (o *OriginalMandate3Choice) SetOriginalMandateIdentification(value string) {
	o.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (o *OriginalMandate3Choice) AddOriginalMandate() *Mandate5 {
	o.OriginalMandate = new(Mandate5)
	return o.OriginalMandate
}
