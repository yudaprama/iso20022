package model

// Description of the mis-matched situation between two baselines or between a baseline and a data set.
type MisMatchReport3 struct {

	// Total number of mismatches between two baselines or between one baseline and one data set.
	NumberOfMisMatches *Number `xml:"NbOfMisMtchs"`

	// Details of each mismatch occurrence.
	MisMatchInformation []*ValidationResult5 `xml:"MisMtchInf,omitempty"`
}

func (m *MisMatchReport3) SetNumberOfMisMatches(value string) {
	m.NumberOfMisMatches = (*Number)(&value)
}

func (m *MisMatchReport3) AddMisMatchInformation() *ValidationResult5 {
	newValue := new(ValidationResult5)
	m.MisMatchInformation = append(m.MisMatchInformation, newValue)
	return newValue
}
