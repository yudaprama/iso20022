package model

// Parameters defining the timing conditions to process an action.
type ProcessTiming3 struct {

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
}

func (p *ProcessTiming3) SetWaitingTime(value string) {
	p.WaitingTime = (*Max9NumericText)(&value)
}

func (p *ProcessTiming3) SetStartTime(value string) {
	p.StartTime = (*ISODateTime)(&value)
}

func (p *ProcessTiming3) SetEndTime(value string) {
	p.EndTime = (*ISODateTime)(&value)
}

func (p *ProcessTiming3) SetPeriod(value string) {
	p.Period = (*Max9NumericText)(&value)
}

func (p *ProcessTiming3) SetMaximumNumber(value string) {
	p.MaximumNumber = (*Number)(&value)
}
