package model

// Time span defined by a start date and time, and an end date and time.
type Period5 struct {

	// Date and time at which the range starts.
	StartDate *DateFormat21Choice `xml:"StartDt"`

	// Date and time at which the range ends.
	EndDate *DateFormat21Choice `xml:"EndDt"`
}

func (p *Period5) AddStartDate() *DateFormat21Choice {
	p.StartDate = new(DateFormat21Choice)
	return p.StartDate
}

func (p *Period5) AddEndDate() *DateFormat21Choice {
	p.EndDate = new(DateFormat21Choice)
	return p.EndDate
}
