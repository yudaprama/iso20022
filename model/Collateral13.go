package model

// Provides for each collateral account the report summary and the valuation of each piece of collateral.
type Collateral13 struct {

	// Provides information about the collateral account, that is the identification, the type and optionally the name.
	AccountIdentification *CollateralAccount2 `xml:"AcctId"`

	// Provides the summary of the collateral valuation report.
	ReportSummary *Summary1 `xml:"RptSummry"`

	// Provides additionnal information about the collateral valuation that has been posted.
	CollateralValuation []*CollateralValuation5 `xml:"CollValtn,omitempty"`
}

func (c *Collateral13) AddAccountIdentification() *CollateralAccount2 {
	c.AccountIdentification = new(CollateralAccount2)
	return c.AccountIdentification
}

func (c *Collateral13) AddReportSummary() *Summary1 {
	c.ReportSummary = new(Summary1)
	return c.ReportSummary
}

func (c *Collateral13) AddCollateralValuation() *CollateralValuation5 {
	newValue := new(CollateralValuation5)
	c.CollateralValuation = append(c.CollateralValuation, newValue)
	return newValue
}
