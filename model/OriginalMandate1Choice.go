package model

// Specifies the mandate that is being accepted.
type OriginalMandate1Choice struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId"`

	// Set of elements used to provide the original mandate data.
	OriginalMandate *MandateInformation1 `xml:"OrgnlMndt"`
}

func (o *OriginalMandate1Choice) SetOriginalMandateIdentification(value string) {
	o.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (o *OriginalMandate1Choice) AddOriginalMandate() *MandateInformation1 {
	o.OriginalMandate = new(MandateInformation1)
	return o.OriginalMandate
}
