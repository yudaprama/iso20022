package model

// Rules which apply to the BPO (Bank Payment Obligation).
type BPOApplicableRules1Choice struct {

	// URBPO are rules that apply to a BPO when the Payment Obligation Segment within an Established Baseline expressly states that it is subject to these rules or when each Involved Bank agrees in a separate agreement that a BPO is subject to these rules. If an Established Baseline or separate agreement does not indicate the applicable version of URBPO, the BPO will be subject to the latest version in effect when the Baseline is established in accordance with sub-article 9 (d).
	URBPOVersion *DecimalNumber `xml:"URBPOVrsn"`

	// Applicable rules are not URBPO and are specified here with version.
	OtherRulesAndVersion *Max35Text `xml:"OthrRulesAndVrsn"`
}

func (b *BPOApplicableRules1Choice) SetURBPOVersion(value string) {
	b.URBPOVersion = (*DecimalNumber)(&value)
}

func (b *BPOApplicableRules1Choice) SetOtherRulesAndVersion(value string) {
	b.OtherRulesAndVersion = (*Max35Text)(&value)
}
