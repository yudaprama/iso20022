package model

// Definition of retry process if activation of an action fails.
type ProcessRetry2 struct {

	// Time period to wait for a retry in months, days, hours and minutes, leading zeros could be omitted.
	Delay *Max9NumericText `xml:"Dely"`

	// Maximum number of retries.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`
}

func (p *ProcessRetry2) SetDelay(value string) {
	p.Delay = (*Max9NumericText)(&value)
}

func (p *ProcessRetry2) SetMaximumNumber(value string) {
	p.MaximumNumber = (*Number)(&value)
}
