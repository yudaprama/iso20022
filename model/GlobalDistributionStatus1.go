package model

// Provides information about the status of the global movement.
type GlobalDistributionStatus1 struct {

	// Provides information about the processing status of the global movement
	ProcessedStatus *DistributionProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *DistributionRejectionStatus1 `xml:"RjctdSts"`
}

func (g *GlobalDistributionStatus1) AddProcessedStatus() *DistributionProcessingStatus1 {
	g.ProcessedStatus = new(DistributionProcessingStatus1)
	return g.ProcessedStatus
}

func (g *GlobalDistributionStatus1) AddRejectedStatus() *DistributionRejectionStatus1 {
	g.RejectedStatus = new(DistributionRejectionStatus1)
	return g.RejectedStatus
}
