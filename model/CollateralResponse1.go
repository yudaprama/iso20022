package model

// Provides additional information on the collateral proposal(s), that is either in cash, securities or other types.
type CollateralResponse1 struct {

	// Provides details on the securities collateral proposal.
	SecuritiesCollateralResponse []*SecuritiesCollateralResponse1 `xml:"SctiesCollRspn,omitempty"`

	// Provides details on the cash collateral proposal.
	CashCollateralResponse []*CashCollateralResponse1 `xml:"CshCollRspn,omitempty"`

	// Provides details on other collateral proposal.
	OtherCollateralResponse []*OtherCollateralResponse1 `xml:"OthrCollRspn,omitempty"`
}

func (c *CollateralResponse1) AddSecuritiesCollateralResponse() *SecuritiesCollateralResponse1 {
	newValue := new(SecuritiesCollateralResponse1)
	c.SecuritiesCollateralResponse = append(c.SecuritiesCollateralResponse, newValue)
	return newValue
}

func (c *CollateralResponse1) AddCashCollateralResponse() *CashCollateralResponse1 {
	newValue := new(CashCollateralResponse1)
	c.CashCollateralResponse = append(c.CashCollateralResponse, newValue)
	return newValue
}

func (c *CollateralResponse1) AddOtherCollateralResponse() *OtherCollateralResponse1 {
	newValue := new(OtherCollateralResponse1)
	c.OtherCollateralResponse = append(c.OtherCollateralResponse, newValue)
	return newValue
}
