package model

// Provides the details of the security pledge as collateral.
type Collateral14 struct {

	// Provides the values of the security pledged as collateral.
	Valuation *SecuredCollateral2Choice `xml:"Valtn"`

	// Risk control measure applied to underlying collateral whereby the value of that underlying collateral is calculated as the market value of the assets reduced by a certain percentage.
	//
	// For reporting purposes the collateral haircut will be calculated as 100 minus the ratio between the cash lent/borrowed and the market value including accrued interest of the collateral pledged times 100.
	//
	// In the case of multi-collateral repos the haircut will be based on the ratio between the cash borrowed/lent and the market value, including accrued interest of each of the individual collateral pledged.
	//
	// Only actual values, as opposed to estimated or default values will be reported for this variable.
	//
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Identifies all repurchase agreements conducted against general collateral and those conducted against special collateral.
	// - General Collateral is a repurchase transaction in which the security lender may choose the security to pledge as collateral with the cash provider amongst a relatively wide range of securities meeting predefined criteria;
	// - Special Collateral is a repurchase transaction in which the cash provider requests a specific security (individual ISIN) to be provided by the cash borrower.
	//
	// Usage:
	// This field is optional and it should be provided only in case it is feasible for the reporting agent.
	SpecialCollateralIndicator *SpecialCollateral1Code `xml:"SpclCollInd,omitempty"`
}

func (c *Collateral14) AddValuation() *SecuredCollateral2Choice {
	c.Valuation = new(SecuredCollateral2Choice)
	return c.Valuation
}

func (c *Collateral14) SetHaircut(value string) {
	c.Haircut = (*PercentageRate)(&value)
}

func (c *Collateral14) SetSpecialCollateralIndicator(value string) {
	c.SpecialCollateralIndicator = (*SpecialCollateral1Code)(&value)
}
