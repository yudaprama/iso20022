package model

// Time span defined by a start date and time, and an end date and time.
type Period1 struct {

	// Date and time at which the range starts.
	StartDate *DateFormat4Choice `xml:"StartDt"`

	// Date and time at which the range ends.
	EndDate *DateFormat4Choice `xml:"EndDt"`
}

func (p *Period1) AddStartDate() *DateFormat4Choice {
	p.StartDate = new(DateFormat4Choice)
	return p.StartDate
}

func (p *Period1) AddEndDate() *DateFormat4Choice {
	p.EndDate = new(DateFormat4Choice)
	return p.EndDate
}
