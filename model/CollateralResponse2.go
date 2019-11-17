package model

// Provides additional information on the collateral proposal(s), that is either in cash, securities or other types.
type CollateralResponse2 struct {

	// Provides details on the securities collateral proposal.
	SecuritiesCollateralResponse []*SecuritiesCollateralResponse1 `xml:"SctiesCollRspn,omitempty"`

	// Provides details on the cash collateral proposal.
	CashCollateralResponse []*CashCollateralResponse2 `xml:"CshCollRspn,omitempty"`

	// Provides details on other collateral proposal.
	OtherCollateralResponse []*OtherCollateralResponse2 `xml:"OthrCollRspn,omitempty"`
}

func (c *CollateralResponse2) AddSecuritiesCollateralResponse() *SecuritiesCollateralResponse1 {
	newValue := new(SecuritiesCollateralResponse1)
	c.SecuritiesCollateralResponse = append(c.SecuritiesCollateralResponse, newValue)
	return newValue
}

func (c *CollateralResponse2) AddCashCollateralResponse() *CashCollateralResponse2 {
	newValue := new(CashCollateralResponse2)
	c.CashCollateralResponse = append(c.CashCollateralResponse, newValue)
	return newValue
}

func (c *CollateralResponse2) AddOtherCollateralResponse() *OtherCollateralResponse2 {
	newValue := new(OtherCollateralResponse2)
	c.OtherCollateralResponse = append(c.OtherCollateralResponse, newValue)
	return newValue
}
