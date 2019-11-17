package model

// Provides the dispute notification details for the variation margin and the segregated independent amount, or the segregated independent amount only.
type DisputeNotification1Choice struct {

	// Provides the dispute notification details for the variation margin and optionaly the segregated independent amount.
	DisputeNotificationDetails *DisputeNotification1 `xml:"DsptNtfctnDtls"`

	// Provides the dispute notification details for the segregated independent amount only.
	SegregatedIndependentAmountDisputeDetails *SegregatedIndependentAmountDispute1 `xml:"SgrtdIndpdntAmtDsptDtls"`
}

func (d *DisputeNotification1Choice) AddDisputeNotificationDetails() *DisputeNotification1 {
	d.DisputeNotificationDetails = new(DisputeNotification1)
	return d.DisputeNotificationDetails
}

func (d *DisputeNotification1Choice) AddSegregatedIndependentAmountDisputeDetails() *SegregatedIndependentAmountDispute1 {
	d.SegregatedIndependentAmountDisputeDetails = new(SegregatedIndependentAmountDispute1)
	return d.SegregatedIndependentAmountDisputeDetails
}
