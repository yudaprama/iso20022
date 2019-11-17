package model

// Configuration parameters of data exchanges.
type ExchangeConfiguration4 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Definition of retry process if activation of an action fails.
	ReTry *ProcessRetry2 `xml:"ReTry,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming3 `xml:"TmCond,omitempty"`
}

func (e *ExchangeConfiguration4) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration4) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration4) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration4) AddReTry() *ProcessRetry2 {
	e.ReTry = new(ProcessRetry2)
	return e.ReTry
}

func (e *ExchangeConfiguration4) AddTimeCondition() *ProcessTiming3 {
	e.TimeCondition = new(ProcessTiming3)
	return e.TimeCondition
}
