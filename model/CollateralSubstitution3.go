package model

// Provides details about the collateral that will be substituted.
type CollateralSubstitution3 struct {

	// Indicates whether the collateral substitution request is new or updated.
	CollateralSubstitutionSequence *CollateralSubstitutionSequence1Code `xml:"CollSbstitnSeq"`

	// Cash value of the requested collateral substitution transfer in the base currency.
	SubstitutionRequirement *ActiveCurrencyAndAmount `xml:"SbstitnRqrmnt"`

	// Specifies if the collateral that is substituted was posted against the variation margin or the independent amount.
	CollateralSubstitutionType *CollateralSubstitutionType1Code `xml:"CollSbstitnTp"`

	// Identifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral3 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral3 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral2 `xml:"OthrColl,omitempty"`

	// Provides details on the identification of previously sent and/or received message(s), in case of updated substitution request.
	LinkedReferences *Reference17 `xml:"LkdRefs,omitempty"`
}

func (c *CollateralSubstitution3) SetCollateralSubstitutionSequence(value string) {
	c.CollateralSubstitutionSequence = (*CollateralSubstitutionSequence1Code)(&value)
}

func (c *CollateralSubstitution3) SetSubstitutionRequirement(value, currency string) {
	c.SubstitutionRequirement = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralSubstitution3) SetCollateralSubstitutionType(value string) {
	c.CollateralSubstitutionType = (*CollateralSubstitutionType1Code)(&value)
}

func (c *CollateralSubstitution3) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *CollateralSubstitution3) AddSecuritiesCollateral() *SecuritiesCollateral3 {
	newValue := new(SecuritiesCollateral3)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution3) AddCashCollateral() *CashCollateral3 {
	newValue := new(CashCollateral3)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution3) AddOtherCollateral() *OtherCollateral2 {
	newValue := new(OtherCollateral2)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution3) AddLinkedReferences() *Reference17 {
	c.LinkedReferences = new(Reference17)
	return c.LinkedReferences
}
