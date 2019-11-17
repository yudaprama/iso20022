package model

// Provides the specification of the valuation of a collateral, based on the sector and the asset classification.
type CollateralValuation7 struct {

	// Specifies whether the collateral is a pool collateral or not.
	PoolStatus *CollateralPool1Code `xml:"PoolSts"`

	// Identifies the asset class pledged as collateral, expressed as an ISO 10962 Classification of Financial Instrument (CFI).
	Type *CFIOct2015Identifier `xml:"Tp"`

	// Provides the institutional sector, such as central government, central bank, etc. of the issuer of collateral.
	Sector *SNA2008SectorIdentifier `xml:"Sctr"`

	// Nominal amount of money of the security pledged as collateral, when the collateral cannot be identified through an individual or basket ISIN.
	NominalAmount *ActiveCurrencyAndAmount `xml:"NmnlAmt,omitempty"`
}

func (c *CollateralValuation7) SetPoolStatus(value string) {
	c.PoolStatus = (*CollateralPool1Code)(&value)
}

func (c *CollateralValuation7) SetType(value string) {
	c.Type = (*CFIOct2015Identifier)(&value)
}

func (c *CollateralValuation7) SetSector(value string) {
	c.Sector = (*SNA2008SectorIdentifier)(&value)
}

func (c *CollateralValuation7) SetNominalAmount(value, currency string) {
	c.NominalAmount = NewActiveCurrencyAndAmount(value, currency)
}
