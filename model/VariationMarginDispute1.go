package model

// Provides for the variation margin, the dispute details like the dispute amount or the dispute date and the resolution type details.
type VariationMarginDispute1 struct {

	// Details of the disputed instruction.
	DisputeDetails *Dispute1 `xml:"DsptDtls"`

	// Specifies the type of dispute that is to be resolved regarding the disputed collateral amount.
	ResolutionTypeDetails []*DisputeResolutionType2Choice `xml:"RsltnTpDtls,omitempty"`
}

func (v *VariationMarginDispute1) AddDisputeDetails() *Dispute1 {
	v.DisputeDetails = new(Dispute1)
	return v.DisputeDetails
}

func (v *VariationMarginDispute1) AddResolutionTypeDetails() *DisputeResolutionType2Choice {
	newValue := new(DisputeResolutionType2Choice)
	v.ResolutionTypeDetails = append(v.ResolutionTypeDetails, newValue)
	return newValue
}
