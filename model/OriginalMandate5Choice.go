package model

// Specifies the mandate that is being accepted.
type OriginalMandate5Choice struct {

	// Unique identification, as assigned by the responsible party (such as the creditor) or agent (such as the debtor agent), to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId"`

	// Provides the original mandate data.
	OriginalMandate *Mandate11 `xml:"OrgnlMndt"`
}

func (o *OriginalMandate5Choice) SetOriginalMandateIdentification(value string) {
	o.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (o *OriginalMandate5Choice) AddOriginalMandate() *Mandate11 {
	o.OriginalMandate = new(Mandate11)
	return o.OriginalMandate
}
