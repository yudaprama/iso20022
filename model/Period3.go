package model

// Time span defined by a start date and time, and an end date and time.
type Period3 struct {

	// Date and time at which the range starts.
	StartDate *DateFormat12Choice `xml:"StartDt"`

	// Date and time at which the range ends.
	EndDate *DateFormat12Choice `xml:"EndDt"`
}

func (p *Period3) AddStartDate() *DateFormat12Choice {
	p.StartDate = new(DateFormat12Choice)
	return p.StartDate
}

func (p *Period3) AddEndDate() *DateFormat12Choice {
	p.EndDate = new(DateFormat12Choice)
	return p.EndDate
}
