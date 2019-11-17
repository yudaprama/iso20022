package model

// Specifies the mandate that is being accepted.
type OriginalMandate2Choice struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId"`

	// Provides the original mandate data.
	OriginalMandate *Mandate1 `xml:"OrgnlMndt"`
}

func (o *OriginalMandate2Choice) SetOriginalMandateIdentification(value string) {
	o.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (o *OriginalMandate2Choice) AddOriginalMandate() *Mandate1 {
	o.OriginalMandate = new(Mandate1)
	return o.OriginalMandate
}
