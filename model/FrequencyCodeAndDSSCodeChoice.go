package model

// Choice between formats for the frequency.
type FrequencyCodeAndDSSCodeChoice struct {

	// Frequency expressed as a code.
	FrequencyAsCode *Frequency1Code `xml:"FrqcyAsCd"`

	// Frequency expressed as a data source scheme and a code used within the data source scheme.
	FrequencyAsDSS *GenericIdentification7 `xml:"FrqcyAsDSS"`
}

func (f *FrequencyCodeAndDSSCodeChoice) SetFrequencyAsCode(value string) {
	f.FrequencyAsCode = (*Frequency1Code)(&value)
}

func (f *FrequencyCodeAndDSSCodeChoice) AddFrequencyAsDSS() *GenericIdentification7 {
	f.FrequencyAsDSS = new(GenericIdentification7)
	return f.FrequencyAsDSS
}
