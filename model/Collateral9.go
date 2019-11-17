package model

// Provides for each collateral account the report summary and the valuation of each piece of collateral.
type Collateral9 struct {

	// Provides information about the collateral account, that is the identification, the type and optionally the name.
	AccountIdentification *CollateralAccount1 `xml:"AcctId"`

	// Provides the summary of the collateral valuation report.
	ReportSummary *Summary1 `xml:"RptSummry"`

	// Provides additionnal information about the collateral valuation that has been posted.
	CollateralValuation []*CollateralValuation2 `xml:"CollValtn,omitempty"`
}

func (c *Collateral9) AddAccountIdentification() *CollateralAccount1 {
	c.AccountIdentification = new(CollateralAccount1)
	return c.AccountIdentification
}

func (c *Collateral9) AddReportSummary() *Summary1 {
	c.ReportSummary = new(Summary1)
	return c.ReportSummary
}

func (c *Collateral9) AddCollateralValuation() *CollateralValuation2 {
	newValue := new(CollateralValuation2)
	c.CollateralValuation = append(c.CollateralValuation, newValue)
	return newValue
}
