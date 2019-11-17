package model

// Information about a blocked holding.
type BlockedHoldingDetails2 struct {

	// Specifies how the blocked account holding is defined.
	BlockedHolding *Holding1Code `xml:"BlckdHldg"`

	// When an account is blocked at the level of fund or security, partially, this is the number of units blocked.
	PartialHoldingUnits *DecimalNumber `xml:"PrtlHldgUnits,omitempty"`

	// When an account is blocked at the level of fund or security, this specifies the certificate number of the blocked units.
	HoldingCertificateNumber *Max35Text `xml:"HldgCertNb,omitempty"`
}

func (b *BlockedHoldingDetails2) SetBlockedHolding(value string) {
	b.BlockedHolding = (*Holding1Code)(&value)
}

func (b *BlockedHoldingDetails2) SetPartialHoldingUnits(value string) {
	b.PartialHoldingUnits = (*DecimalNumber)(&value)
}

func (b *BlockedHoldingDetails2) SetHoldingCertificateNumber(value string) {
	b.HoldingCertificateNumber = (*Max35Text)(&value)
}
