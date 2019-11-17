package model

// Further detailed information on the exchange rate that has been used in the payment transaction.
type ExchangeRateInformation1 struct {

	// The factor used for conversion of an amount from one currency to another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies the type used to complete the currency exchange.
	RateType *ExchangeRateType1Code `xml:"RateTp,omitempty"`

	// Unique and unambiguous reference to the foreign exchange contract agreed between the initiating party/creditor and the debtor agent.
	ContractIdentification *Max35Text `xml:"CtrctId,omitempty"`
}

func (e *ExchangeRateInformation1) SetExchangeRate(value string) {
	e.ExchangeRate = (*BaseOneRate)(&value)
}

func (e *ExchangeRateInformation1) SetRateType(value string) {
	e.RateType = (*ExchangeRateType1Code)(&value)
}

func (e *ExchangeRateInformation1) SetContractIdentification(value string) {
	e.ContractIdentification = (*Max35Text)(&value)
}
