package model

// Choice of format for a frequency, for example, the frequency of payment.
type Frequency36Choice struct {

	// Specifies a frequency in terms of a specified period type.
	Type *Frequency6Code `xml:"Tp"`

	// Specifies a frequency in terms of a count per period within a specified period type.
	Period *FrequencyPeriod1 `xml:"Prd"`

	// Specifies a frequency in terms of an exact point in time or moment within a specified period type.
	PointInTime *FrequencyAndMoment1 `xml:"PtInTm"`
}

func (f *Frequency36Choice) SetType(value string) {
	f.Type = (*Frequency6Code)(&value)
}

func (f *Frequency36Choice) AddPeriod() *FrequencyPeriod1 {
	f.Period = new(FrequencyPeriod1)
	return f.Period
}

func (f *Frequency36Choice) AddPointInTime() *FrequencyAndMoment1 {
	f.PointInTime = new(FrequencyAndMoment1)
	return f.PointInTime
}
