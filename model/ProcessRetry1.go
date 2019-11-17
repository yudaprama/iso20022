package model

// Definition of retry process if activation of an action fails.
type ProcessRetry1 struct {

	// Time period to wait for a retry in months, days, hours and minutes, leading zeros could be omitted.
	Delay *Max9NumericText `xml:"Dely"`

	// Maximum number of retries.
	MaximumNumber *Number `xml:"MaxNb,omitempty"`

	// Time of the last retry.
	LastReTryTime *ISOTime `xml:"LastReTryTm,omitempty"`
}

func (p *ProcessRetry1) SetDelay(value string) {
	p.Delay = (*Max9NumericText)(&value)
}

func (p *ProcessRetry1) SetMaximumNumber(value string) {
	p.MaximumNumber = (*Number)(&value)
}

func (p *ProcessRetry1) SetLastReTryTime(value string) {
	p.LastReTryTime = (*ISOTime)(&value)
}
