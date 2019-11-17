package model

// Date parameters.
type DateInformation1 struct {

	// Date on which a recurrent date will commence.
	StartDate *ISODate `xml:"StartDt"`

	// Specifies the regularity of the trigger date.
	Frequency *ExternalDateFrequency1Code `xml:"Frqcy"`

	// Maximum number of trigger date occurrence cycles.
	Number *Number `xml:"Nb"`
}

func (d *DateInformation1) SetStartDate(value string) {
	d.StartDate = (*ISODate)(&value)
}

func (d *DateInformation1) SetFrequency(value string) {
	d.Frequency = (*ExternalDateFrequency1Code)(&value)
}

func (d *DateInformation1) SetNumber(value string) {
	d.Number = (*Number)(&value)
}
