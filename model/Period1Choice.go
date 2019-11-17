package model

// Choice between a period or a period code.
type Period1Choice struct {

	// Time span defined by a start date and time, and an end date and time.
	Period *Period3 `xml:"Prd"`

	// Standard code to specify the type of period.
	PeriodCode *DateType6Code `xml:"PrdCd"`
}

func (p *Period1Choice) AddPeriod() *Period3 {
	p.Period = new(Period3)
	return p.Period
}

func (p *Period1Choice) SetPeriodCode(value string) {
	p.PeriodCode = (*DateType6Code)(&value)
}
