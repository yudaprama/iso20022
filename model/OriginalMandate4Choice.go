package model

// Specifies the mandate that is being accepted.
type OriginalMandate4Choice struct {

	// Unique identification, as assigned by the responsible party (such as the creditor) or agent (such as the debtor agent), to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId"`

	// Provides the original mandate data.
	OriginalMandate *Mandate9 `xml:"OrgnlMndt"`
}

func (o *OriginalMandate4Choice) SetOriginalMandateIdentification(value string) {
	o.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (o *OriginalMandate4Choice) AddOriginalMandate() *Mandate9 {
	o.OriginalMandate = new(Mandate9)
	return o.OriginalMandate
}
