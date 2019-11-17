package model

// Defines a frequency in terms on counts per period for a specific period type.
type FrequencyPeriod1 struct {

	// Period for which the number of instructions are to be created and processed.
	Type *Frequency6Code `xml:"Tp"`

	// Number of instructions to be created and processed during the specified period
	CountPerPeriod *DecimalNumber `xml:"CntPerPrd"`
}

func (f *FrequencyPeriod1) SetType(value string) {
	f.Type = (*Frequency6Code)(&value)
}

func (f *FrequencyPeriod1) SetCountPerPeriod(value string) {
	f.CountPerPeriod = (*DecimalNumber)(&value)
}
