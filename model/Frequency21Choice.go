package model

// Choice of format for a frequency, for example, the frequency of payment.
type Frequency21Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Type *Frequency6Code `xml:"Tp"`

	// Frequency expressed as a proprietary code.
	Period *FrequencyPeriod1 `xml:"Prd"`
}

func (f *Frequency21Choice) SetType(value string) {
	f.Type = (*Frequency6Code)(&value)
}

func (f *Frequency21Choice) AddPeriod() *FrequencyPeriod1 {
	f.Period = new(FrequencyPeriod1)
	return f.Period
}
