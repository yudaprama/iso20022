package model

// Provides more details on the response such as the response type, the collateral identification, and optionally further details in case of rejection.
type CashCollateralResponse1 struct {

	// Specifies the status of the collateral proposal.
	ResponseType *Status4Code `xml:"RspnTp"`

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	CashAccountIdentification *AccountIdentification4Choice `xml:"CshAcctId,omitempty"`

	// Specifies the reason why the instruction/cancellation request has a rejected status.
	RejectionReason *RejectionReasonV021Code `xml:"RjctnRsn,omitempty"`

	// Additional information regarding why the collateral proposal has a rejected status.
	RejectionInformation *Max35Text `xml:"RjctnInf,omitempty"`
}

func (c *CashCollateralResponse1) SetResponseType(value string) {
	c.ResponseType = (*Status4Code)(&value)
}

func (c *CashCollateralResponse1) SetCollateralIdentification(value string) {
	c.CollateralIdentification = (*Max35Text)(&value)
}

func (c *CashCollateralResponse1) SetAssetNumber(value string) {
	c.AssetNumber = (*Max35Text)(&value)
}

func (c *CashCollateralResponse1) AddCashAccountIdentification() *AccountIdentification4Choice {
	c.CashAccountIdentification = new(AccountIdentification4Choice)
	return c.CashAccountIdentification
}

func (c *CashCollateralResponse1) SetRejectionReason(value string) {
	c.RejectionReason = (*RejectionReasonV021Code)(&value)
}

func (c *CashCollateralResponse1) SetRejectionInformation(value string) {
	c.RejectionInformation = (*Max35Text)(&value)
}
