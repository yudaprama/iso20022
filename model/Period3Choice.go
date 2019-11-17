package model

// Choice between a period or a period code.
type Period3Choice struct {

	// Time span defined by a start date and time, and an end date and time.
	Period *Period4 `xml:"Prd"`

	// Standard code to specify the type of period.
	PeriodCode *DateType8Code `xml:"PrdCd"`
}

func (p *Period3Choice) AddPeriod() *Period4 {
	p.Period = new(Period4)
	return p.Period
}

func (p *Period3Choice) SetPeriodCode(value string) {
	p.PeriodCode = (*DateType8Code)(&value)
}
