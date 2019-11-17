package model

// Details related to the demand.
type Demand3 struct {

	// Unique and unambiguous identifier assigned by the presenting party to the demand.
	Identification *Max35Text `xml:"Id"`

	// Date and time of demand submission.
	SubmissionDateTime *ISODateTime `xml:"SubmissnDtTm"`

	// Amount and currency of the demand.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (d *Demand3) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *Demand3) SetSubmissionDateTime(value string) {
	d.SubmissionDateTime = (*ISODateTime)(&value)
}

func (d *Demand3) SetAmount(value, currency string) {
	d.Amount = NewActiveCurrencyAndAmount(value, currency)
}
