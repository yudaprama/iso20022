package model

// Set of elements used to identify a proprietary date.
type ProprietaryDate2 struct {

	// Specifies the type of date.
	Type *Max35Text `xml:"Tp"`

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`
}

func (p *ProprietaryDate2) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryDate2) AddDate() *DateAndDateTimeChoice {
	p.Date = new(DateAndDateTimeChoice)
	return p.Date
}
