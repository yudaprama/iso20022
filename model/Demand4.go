package model

// Information about a demand.
type Demand4 struct {

	// Unique and unambiguous identifier assigned by the presenting party to the demand.
	Identification *Max35Text `xml:"Id"`

	// Amount and currency of the demand.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (d *Demand4) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *Demand4) SetAmount(value, currency string) {
	d.Amount = NewActiveCurrencyAndAmount(value, currency)
}
