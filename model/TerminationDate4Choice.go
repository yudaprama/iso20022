package model

// Choice of format for the termination date.
type TerminationDate4Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Closing date/time or maturity date/time of the transaction expressed as an ISO 20022 code.
	Code *DateCode18Choice `xml:"Cd"`
}

func (t *TerminationDate4Choice) AddDate() *DateAndDateTimeChoice {
	t.Date = new(DateAndDateTimeChoice)
	return t.Date
}

func (t *TerminationDate4Choice) AddCode() *DateCode18Choice {
	t.Code = new(DateCode18Choice)
	return t.Code
}
