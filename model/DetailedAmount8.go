package model

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount8 struct {

	// Amount after the currency exchange.
	// It corresponds to ISO 8583 field number 6, completed by the field number 51 for the versions 87 and 93.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Exchange rate to the currency of the amount.
	// It corresponds to ISO 8583 field number 10.
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which the exchange rate has been quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Text to display on the cardholder or to print on the cardholder bank statement.
	Label *Max140Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount8) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount8) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DetailedAmount8) SetQuotationDate(value string) {
	d.QuotationDate = (*ISODateTime)(&value)
}

func (d *DetailedAmount8) SetLabel(value string) {
	d.Label = (*Max140Text)(&value)
}
