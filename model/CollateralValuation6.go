package model

// Provides the valuation of a collateral, identified through an ISIN.
type CollateralValuation6 struct {

	// Nominal amount of the security pledged as collateral. Except for tri-party repos and any other transaction in which the security pledged is not identified via a single ISIN.
	NominalAmount *ActiveCurrencyAndAmount `xml:"NmnlAmt,omitempty"`

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINOct2015Identifier `xml:"ISIN"`
}

func (c *CollateralValuation6) SetNominalAmount(value, currency string) {
	c.NominalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralValuation6) SetISIN(value string) {
	c.ISIN = (*ISINOct2015Identifier)(&value)
}
