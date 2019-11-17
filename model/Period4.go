package model

// Time span defined by a start date and time, and an end date and time.
type Period4 struct {

	// Date and time at which the range starts.
	StartDate *DateFormat18Choice `xml:"StartDt"`

	// Date and time at which the range ends.
	EndDate *DateFormat18Choice `xml:"EndDt"`
}

func (p *Period4) AddStartDate() *DateFormat18Choice {
	p.StartDate = new(DateFormat18Choice)
	return p.StartDate
}

func (p *Period4) AddEndDate() *DateFormat18Choice {
	p.EndDate = new(DateFormat18Choice)
	return p.EndDate
}
