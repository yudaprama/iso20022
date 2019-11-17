package model

// Provides the exercise date or notice period for a call/put option.
type OptionDateOrPeriod1Choice struct {

	// First date on which the call option or the put option can be exercised.
	EarliestExerciseDate *ISODate `xml:"EarlstExrcDt"`

	// Number of calendar days that the holder of the instrument/issuer of the instrument will give to the issuer/holder of the instrument before exercising the put/call option.
	NoticePeriod *Number `xml:"NtcePrd"`
}

func (o *OptionDateOrPeriod1Choice) SetEarliestExerciseDate(value string) {
	o.EarliestExerciseDate = (*ISODate)(&value)
}

func (o *OptionDateOrPeriod1Choice) SetNoticePeriod(value string) {
	o.NoticePeriod = (*Number)(&value)
}
