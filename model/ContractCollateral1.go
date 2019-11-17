package model

// Further details on the contract collateral.
type ContractCollateral1 struct {

	// Total amount of the collateral as defined in the contract.
	TotalAmount *ActiveCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed description of the collateral.
	CollateralDescription []*CashCollateral5 `xml:"CollDesc,omitempty"`

	// Further information on the contract collateral.
	AdditionalInformation *Max1025Text `xml:"AddtlInf,omitempty"`
}

func (c *ContractCollateral1) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *ContractCollateral1) AddCollateralDescription() *CashCollateral5 {
	newValue := new(CashCollateral5)
	c.CollateralDescription = append(c.CollateralDescription, newValue)
	return newValue
}

func (c *ContractCollateral1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max1025Text)(&value)
}
