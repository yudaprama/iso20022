package model

// Choice of format for the termination date.
type TerminationDate2Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Closing date/time or maturity date/time of the transaction expressed as an ISO 20022 code.
	Code *DateCode1Choice `xml:"Cd"`
}

func (t *TerminationDate2Choice) AddDate() *DateAndDateTimeChoice {
	t.Date = new(DateAndDateTimeChoice)
	return t.Date
}

func (t *TerminationDate2Choice) AddCode() *DateCode1Choice {
	t.Code = new(DateCode1Choice)
	return t.Code
}
