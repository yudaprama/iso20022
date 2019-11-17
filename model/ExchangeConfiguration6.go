package model

// Configuration parameters of data exchanges.
type ExchangeConfiguration6 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Definition of retry process if activation of an action fails.
	ReTry *ProcessRetry2 `xml:"ReTry,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming4 `xml:"TmCond,omitempty"`
}

func (e *ExchangeConfiguration6) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration6) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration6) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration6) AddReTry() *ProcessRetry2 {
	e.ReTry = new(ProcessRetry2)
	return e.ReTry
}

func (e *ExchangeConfiguration6) AddTimeCondition() *ProcessTiming4 {
	e.TimeCondition = new(ProcessTiming4)
	return e.TimeCondition
}
