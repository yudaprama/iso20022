package model

// Choice between formats for the frequency.
type FrequencyCodeAndDSSCode1Choice struct {

	// Frequency expressed as a code.
	FrequencyAsCode *EventFrequency1Code `xml:"FrqcyAsCd"`

	// Frequency expressed as a data source scheme and a code used within the data source scheme.
	FrequencyAsDSS *GenericIdentification7 `xml:"FrqcyAsDSS"`
}

func (f *FrequencyCodeAndDSSCode1Choice) SetFrequencyAsCode(value string) {
	f.FrequencyAsCode = (*EventFrequency1Code)(&value)
}

func (f *FrequencyCodeAndDSSCode1Choice) AddFrequencyAsDSS() *GenericIdentification7 {
	f.FrequencyAsDSS = new(GenericIdentification7)
	return f.FrequencyAsDSS
}
