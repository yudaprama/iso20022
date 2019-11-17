package model

// Choice of format for the termination date.
type TerminationDate5Choice struct {

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Closing date/time or maturity date/time of the transaction expressed as an ISO 20022 code.
	Code *DateCode32Choice `xml:"Cd"`
}

func (t *TerminationDate5Choice) AddDate() *DateAndDateTimeChoice {
	t.Date = new(DateAndDateTimeChoice)
	return t.Date
}

func (t *TerminationDate5Choice) AddCode() *DateCode32Choice {
	t.Code = new(DateCode32Choice)
	return t.Code
}
