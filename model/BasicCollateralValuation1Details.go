package model

// Basic valuation details of a collateral position.
type BasicCollateralValuation1Details struct {

	// Haircut percentage applied to the market value of underlying assets used as collateral as a risk control measure. The institution valuating the collateral calculates the value of underlying assets based on its market value less a certain percentage (the haircut).
	ValuationHaircut *PercentageRate `xml:"ValtnHrcut"`

	// Place where the valuation haircut was calculated.
	HaircutSource *PartyIdentification15 `xml:"HrcutSrc,omitempty"`
}

func (b *BasicCollateralValuation1Details) SetValuationHaircut(value string) {
	b.ValuationHaircut = (*PercentageRate)(&value)
}

func (b *BasicCollateralValuation1Details) AddHaircutSource() *PartyIdentification15 {
	b.HaircutSource = new(PartyIdentification15)
	return b.HaircutSource
}
