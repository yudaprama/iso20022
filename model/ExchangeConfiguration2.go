package model

// Configuration parameters of data exchanges.
type ExchangeConfiguration2 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming2 `xml:"TmCond,omitempty"`
}

func (e *ExchangeConfiguration2) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration2) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration2) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration2) AddTimeCondition() *ProcessTiming2 {
	e.TimeCondition = new(ProcessTiming2)
	return e.TimeCondition
}
