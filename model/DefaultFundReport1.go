package model

// Provides details on the calculation of the default fund and the collateral that has been posted by the clearing member.
type DefaultFundReport1 struct {

	// Provides details about the calculation of the clearing member contribution to the default fund.
	DefaultFundCalculation []*DefaultFund1 `xml:"DfltFndClctn"`

	// Provides details about the collateral held.
	CollateralDescription []*Collateral3 `xml:"CollDesc"`

	// Excess amount that the central counterparty will restitute to the clearing member or deficit to be provided by the member for the guarantee fund.
	NetExcessOrDeficit *AmountAndDirection21 `xml:"NetXcssOrDfcit"`
}

func (d *DefaultFundReport1) AddDefaultFundCalculation() *DefaultFund1 {
	newValue := new(DefaultFund1)
	d.DefaultFundCalculation = append(d.DefaultFundCalculation, newValue)
	return newValue
}

func (d *DefaultFundReport1) AddCollateralDescription() *Collateral3 {
	newValue := new(Collateral3)
	d.CollateralDescription = append(d.CollateralDescription, newValue)
	return newValue
}

func (d *DefaultFundReport1) AddNetExcessOrDeficit() *AmountAndDirection21 {
	d.NetExcessOrDeficit = new(AmountAndDirection21)
	return d.NetExcessOrDeficit
}
