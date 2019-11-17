package model

// Provide further details on transaction specific interest information that applies to the underlying transaction.
type TransactionInterest3 struct {

	// Total amount of interests and taxes included in the entry amount.
	TotalInterestAndTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlIntrstAndTaxAmt,omitempty"`

	// Individual interest record.
	Record []*InterestRecord1 `xml:"Rcrd,omitempty"`
}

func (t *TransactionInterest3) SetTotalInterestAndTaxAmount(value, currency string) {
	t.TotalInterestAndTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TransactionInterest3) AddRecord() *InterestRecord1 {
	newValue := new(InterestRecord1)
	t.Record = append(t.Record, newValue)
	return newValue
}
