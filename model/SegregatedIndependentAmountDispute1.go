package model

// Provides for the segregated independent amount, the dispute details like the dispute amount or the dispute date and the resolution type details.
type SegregatedIndependentAmountDispute1 struct {

	// Details of the disputed instruction.
	DisputeDetails *Dispute1 `xml:"DsptDtls"`

	// Specifies the type of dispute that is to be resolved regarding the disputed collateral amount.
	DisputeResolutionType1Choice []*DisputeResolutionType1Choice `xml:"DsptRsltnTp1Chc,omitempty"`
}

func (s *SegregatedIndependentAmountDispute1) AddDisputeDetails() *Dispute1 {
	s.DisputeDetails = new(Dispute1)
	return s.DisputeDetails
}

func (s *SegregatedIndependentAmountDispute1) AddDisputeResolutionType1Choice() *DisputeResolutionType1Choice {
	newValue := new(DisputeResolutionType1Choice)
	s.DisputeResolutionType1Choice = append(s.DisputeResolutionType1Choice, newValue)
	return newValue
}
