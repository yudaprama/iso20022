package model

// Defines a frequency in terms a specific moment within a specified period type.
type FrequencyAndMoment1 struct {

	// Period for which the number of instructions are to be created and processed.
	Type *Frequency6Code `xml:"Tp"`

	// Further information on the exact point in time the event should take place.
	PointInTime *Exact2NumericText `xml:"PtInTm"`
}

func (f *FrequencyAndMoment1) SetType(value string) {
	f.Type = (*Frequency6Code)(&value)
}

func (f *FrequencyAndMoment1) SetPointInTime(value string) {
	f.PointInTime = (*Exact2NumericText)(&value)
}
