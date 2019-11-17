package model

// Provides information about the standing instruction.
type CorporateActionStandingInstruction1 struct {

	// Identifies whether the account Holders want their income to be paid net or gross of income tax (default is gross).
	NetOrGross *StandingInstructionGrossNet1Code `xml:"NetOrGrss"`

	// Provides information about the cash distribution standing instruction.
	CashDistributionDetails *CashAccount17 `xml:"CshDstrbtnDtls"`

	// Provides information about the securities distribution standing instruction.
	SecuritiesDistributionDetails *SecuritiesAccount6 `xml:"SctiesDstrbtnDtls"`

	// Additional information about the standing instruction.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionStandingInstruction1) SetNetOrGross(value string) {
	c.NetOrGross = (*StandingInstructionGrossNet1Code)(&value)
}

func (c *CorporateActionStandingInstruction1) AddCashDistributionDetails() *CashAccount17 {
	c.CashDistributionDetails = new(CashAccount17)
	return c.CashDistributionDetails
}

func (c *CorporateActionStandingInstruction1) AddSecuritiesDistributionDetails() *SecuritiesAccount6 {
	c.SecuritiesDistributionDetails = new(SecuritiesAccount6)
	return c.SecuritiesDistributionDetails
}

func (c *CorporateActionStandingInstruction1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
