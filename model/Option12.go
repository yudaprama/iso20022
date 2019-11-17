package model

// Provides the details for a call/put option.
type Option12 struct {

	// Identifies whether the instrument has a call option or a put option. If the instrument contains both options, i.e. a call and a put, both the call option and the put option have to be reported.
	Type *OptionType1Code `xml:"Tp"`

	// Provides the exercise date or notice period for a call/put option.
	DateOrPeriod *OptionDateOrPeriod1Choice `xml:"DtOrPrd"`
}

func (o *Option12) SetType(value string) {
	o.Type = (*OptionType1Code)(&value)
}

func (o *Option12) AddDateOrPeriod() *OptionDateOrPeriod1Choice {
	o.DateOrPeriod = new(OptionDateOrPeriod1Choice)
	return o.DateOrPeriod
}
