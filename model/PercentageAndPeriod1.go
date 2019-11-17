package model

// Specifies a percentage together with a period of time. For overlapping periods, the maximum of all applicable elements at a given date is the result.
type PercentageAndPeriod1 struct {

	// Covered percentage (max 100%).
	Percentage *PercentageBoundedRate `xml:"Pctg"`

	// Start of period or immediate if not specified.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// End of period or indefinite if not specified.
	EndDate *ISODate `xml:"EndDt,omitempty"`
}

func (p *PercentageAndPeriod1) SetPercentage(value string) {
	p.Percentage = (*PercentageBoundedRate)(&value)
}

func (p *PercentageAndPeriod1) SetStartDate(value string) {
	p.StartDate = (*ISODate)(&value)
}

func (p *PercentageAndPeriod1) SetEndDate(value string) {
	p.EndDate = (*ISODate)(&value)
}
