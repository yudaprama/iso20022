package model

// Configuration parameters of data exchanges.
type ExchangeConfiguration3 struct {

	// Exchange policy between parties.
	ExchangePolicy []*ExchangePolicy1Code `xml:"XchgPlcy"`

	// Maximum number of transactions without exchange.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Maximum cumulative amount of the transactions without exchange.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Timing condition for periodic exchanges.
	TimeCondition *ProcessTiming2 `xml:"TmCond,omitempty"`

	// Failed transaction must be exchanged.
	//
	ExchangeFailed *TrueFalseIndicator `xml:"XchgFaild,omitempty"`

	// Indicates that declined transaction must be exchanged
	ExchangeDeclined *TrueFalseIndicator `xml:"XchgDclnd,omitempty"`
}

func (e *ExchangeConfiguration3) AddExchangePolicy(value string) {
	e.ExchangePolicy = append(e.ExchangePolicy, (*ExchangePolicy1Code)(&value))
}

func (e *ExchangeConfiguration3) SetMaximumNumber(value string) {
	e.MaximumNumber = (*Number)(&value)
}

func (e *ExchangeConfiguration3) SetMaximumAmount(value, currency string) {
	e.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (e *ExchangeConfiguration3) AddTimeCondition() *ProcessTiming2 {
	e.TimeCondition = new(ProcessTiming2)
	return e.TimeCondition
}

func (e *ExchangeConfiguration3) SetExchangeFailed(value string) {
	e.ExchangeFailed = (*TrueFalseIndicator)(&value)
}

func (e *ExchangeConfiguration3) SetExchangeDeclined(value string) {
	e.ExchangeDeclined = (*TrueFalseIndicator)(&value)
}
