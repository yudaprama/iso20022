package model

// Provides the collateral details for the secured markets.
type SecuredCollateral2Choice struct {

	// Identifies the security pledged via a single ISIN.
	SingleCollateral *CollateralValuation6 `xml:"SnglColl"`

	// Identifies all securities pledged when the transaction is collateralised with more than one security.
	// Usage:
	// In case of multi-collateral repo, the nominal amount of each collateralised security must be provided.
	MultipleCollateral []*CollateralValuation6 `xml:"MltplColl"`

	// Identifies the pooling of repos in which the collateral basket is identified by an ISIN.
	//
	// Usage: When the collateral basket is identified by an ISIN, the basket ISIN shall be reported.
	PoolCollateral *CollateralValuation6 `xml:"PoolColl"`

	// Identifies the collateral when the asset class pledged as collateral does not correspond to an ISIN.
	OtherCollateral []*CollateralValuation7 `xml:"OthrColl"`
}

func (s *SecuredCollateral2Choice) AddSingleCollateral() *CollateralValuation6 {
	s.SingleCollateral = new(CollateralValuation6)
	return s.SingleCollateral
}

func (s *SecuredCollateral2Choice) AddMultipleCollateral() *CollateralValuation6 {
	newValue := new(CollateralValuation6)
	s.MultipleCollateral = append(s.MultipleCollateral, newValue)
	return newValue
}

func (s *SecuredCollateral2Choice) AddPoolCollateral() *CollateralValuation6 {
	s.PoolCollateral = new(CollateralValuation6)
	return s.PoolCollateral
}

func (s *SecuredCollateral2Choice) AddOtherCollateral() *CollateralValuation7 {
	newValue := new(CollateralValuation7)
	s.OtherCollateral = append(s.OtherCollateral, newValue)
	return newValue
}
