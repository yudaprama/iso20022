package model

// Provides information about the processing status of a global movement.
type DistributionProcessingStatus1 struct {

	// The processing status.
	Status *ProcessedStatus3FormatChoice `xml:"Sts"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (d *DistributionProcessingStatus1) AddStatus() *ProcessedStatus3FormatChoice {
	d.Status = new(ProcessedStatus3FormatChoice)
	return d.Status
}

func (d *DistributionProcessingStatus1) SetAdditionalInformation(value string) {
	d.AdditionalInformation = (*Max350Text)(&value)
}
