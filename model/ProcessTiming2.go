package model

// Parameters defining the timing conditions to process an action.
type ProcessTiming2 struct {

	// Waiting time after the previous action in months, days, hours and minutes, leading zeros could be omitted.
	WaitingTime *Max9NumericText `xml:"WtgTm,omitempty"`

	// Date and time to start the action.
	StartTime *ISODateTime `xml:"StartTm,omitempty"`

	// Date and time after which the action cannot be processed.
	EndTime *ISODateTime `xml:"EndTm,omitempty"`

	// Period delay between cyclic action activation in months, days, hours and minutes, leading zeros could be omitted.
	Period *Max9NumericText `xml:"Prd,omitempty"`

	// Maximum number of cyclic calls.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Definition of retry process if activation of the action fails.
	ReTry *ProcessRetry2 `xml:"ReTry,omitempty"`
}

func (p *ProcessTiming2) SetWaitingTime(value string) {
	p.WaitingTime = (*Max9NumericText)(&value)
}

func (p *ProcessTiming2) SetStartTime(value string) {
	p.StartTime = (*ISODateTime)(&value)
}

func (p *ProcessTiming2) SetEndTime(value string) {
	p.EndTime = (*ISODateTime)(&value)
}

func (p *ProcessTiming2) SetPeriod(value string) {
	p.Period = (*Max9NumericText)(&value)
}

func (p *ProcessTiming2) SetMaximumNumber(value string) {
	p.MaximumNumber = (*Number)(&value)
}

func (p *ProcessTiming2) AddReTry() *ProcessRetry2 {
	p.ReTry = new(ProcessRetry2)
	return p.ReTry
}
