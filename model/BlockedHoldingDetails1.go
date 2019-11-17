package model

// Specifies information about a blocked holding.
type BlockedHoldingDetails1 struct {

	// Specifies how the blocked account holding is defined.
	BlockedHolding *Holding1Code `xml:"BlckdHldg"`

	// When an account is blocked at the level of fund, partially, this is the number of units blocked.
	PartialHoldingUnits *DecimalNumber `xml:"PrtlHldgUnits,omitempty"`

	// When an account is blocked at the level of fund, this specifies the certificate number of the blocked units.
	HoldingCertificateNumber *Max35Text `xml:"HldgCertNb,omitempty"`
}

func (b *BlockedHoldingDetails1) SetBlockedHolding(value string) {
	b.BlockedHolding = (*Holding1Code)(&value)
}

func (b *BlockedHoldingDetails1) SetPartialHoldingUnits(value string) {
	b.PartialHoldingUnits = (*DecimalNumber)(&value)
}

func (b *BlockedHoldingDetails1) SetHoldingCertificateNumber(value string) {
	b.HoldingCertificateNumber = (*Max35Text)(&value)
}
