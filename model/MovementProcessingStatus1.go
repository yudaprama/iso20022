package model

// Provides information about the processing status of an individual movement.
type MovementProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus3FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (m *MovementProcessingStatus1) AddStatus() *ProcessedStatus3FormatChoice {
	m.Status = new(ProcessedStatus3FormatChoice)
	return m.Status
}

func (m *MovementProcessingStatus1) SetAdditionalInformation(value string) {
	m.AdditionalInformation = (*Max350Text)(&value)
}
