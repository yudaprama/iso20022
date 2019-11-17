package model

// Information about the demand.
type Demand2 struct {

	// Unique and unambiguous identifier assigned by the presenting party to the demand.
	Identification *Max35Text `xml:"Id"`

	// Date and time the demand is submitted.
	SubmissionDateTime *ISODateTime `xml:"SubmissnDtTm"`

	// Amount and currency of the demand.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Additional information related to the demand.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (d *Demand2) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *Demand2) SetSubmissionDateTime(value string) {
	d.SubmissionDateTime = (*ISODateTime)(&value)
}

func (d *Demand2) SetAmount(value, currency string) {
	d.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *Demand2) AddAdditionalInformation(value string) {
	d.AdditionalInformation = append(d.AdditionalInformation, (*Max2000Text)(&value))
}
