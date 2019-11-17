package model

// Parameters defining the timing conditions to process an action.
type ProcessTiming4 struct {

	// Date and time to start the action.
	StartTime *ISODateTime `xml:"StartTm,omitempty"`

	// Date and time after which the action cannot be processed.
	EndTime *ISODateTime `xml:"EndTm,omitempty"`

	// Period delay between cyclic action activation in months, days, hours and minutes, leading zeros could be omitted.
	Period *Max9NumericText `xml:"Prd,omitempty"`
}

func (p *ProcessTiming4) SetStartTime(value string) {
	p.StartTime = (*ISODateTime)(&value)
}

func (p *ProcessTiming4) SetEndTime(value string) {
	p.EndTime = (*ISODateTime)(&value)
}

func (p *ProcessTiming4) SetPeriod(value string) {
	p.Period = (*Max9NumericText)(&value)
}
