package model

// Choice between various statuses for an corporate action election cancellation.
type ElectionCancellationStatus1Choice struct {

	// Provides information about the processing status of the request.
	ProcessedStatus *CorporateActionCancellationProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *CorporateActionCancellationRejectionStatus1 `xml:"RjctdSts"`
}

func (e *ElectionCancellationStatus1Choice) AddProcessedStatus() *CorporateActionCancellationProcessingStatus1 {
	e.ProcessedStatus = new(CorporateActionCancellationProcessingStatus1)
	return e.ProcessedStatus
}

func (e *ElectionCancellationStatus1Choice) AddRejectedStatus() *CorporateActionCancellationRejectionStatus1 {
	e.RejectedStatus = new(CorporateActionCancellationRejectionStatus1)
	return e.RejectedStatus
}
