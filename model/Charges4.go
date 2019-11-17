package model

// Provides further details on the charges related to the payment transaction.
type Charges4 struct {

	// Total of all charges and taxes applied to the entry.
	TotalChargesAndTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty"`

	// Provides details of the individual charges record.
	Record []*ChargesRecord2 `xml:"Rcrd,omitempty"`
}

func (c *Charges4) SetTotalChargesAndTaxAmount(value, currency string) {
	c.TotalChargesAndTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charges4) AddRecord() *ChargesRecord2 {
	newValue := new(ChargesRecord2)
	c.Record = append(c.Record, newValue)
	return newValue
}
