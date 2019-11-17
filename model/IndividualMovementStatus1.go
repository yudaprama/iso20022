package model

// Provides information about the processing status.
type IndividualMovementStatus1 struct {

	// Identification of the movement.
	MovementIdentification *Max35Text `xml:"MvmntId,omitempty"`

	// Provides information about the processing status of individual movement.
	ProcessedStatus *MovementProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *DistributionRejectionStatus1 `xml:"RjctdSts"`
}

func (i *IndividualMovementStatus1) SetMovementIdentification(value string) {
	i.MovementIdentification = (*Max35Text)(&value)
}

func (i *IndividualMovementStatus1) AddProcessedStatus() *MovementProcessingStatus1 {
	i.ProcessedStatus = new(MovementProcessingStatus1)
	return i.ProcessedStatus
}

func (i *IndividualMovementStatus1) AddRejectedStatus() *DistributionRejectionStatus1 {
	i.RejectedStatus = new(DistributionRejectionStatus1)
	return i.RejectedStatus
}
