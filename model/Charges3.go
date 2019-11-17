package model

// Provides further details on the charges related to the payment transaction.
type Charges3 struct {

	// Total of all charges and taxes applied to the entry.
	TotalChargesAndTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty"`

	// Provides details of the individual charges record.
	Record []*ChargesRecord1 `xml:"Rcrd,omitempty"`
}

func (c *Charges3) SetTotalChargesAndTaxAmount(value, currency string) {
	c.TotalChargesAndTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charges3) AddRecord() *ChargesRecord1 {
	newValue := new(ChargesRecord1)
	c.Record = append(c.Record, newValue)
	return newValue
}
