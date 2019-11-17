package model

// Provides details about the collateral that will be substituted.
type CollateralSubstitution4 struct {

	// Indicates whether the collateral substitution request is new or updated.
	CollateralSubstitutionSequence *CollateralSubstitutionSequence1Code `xml:"CollSbstitnSeq"`

	// Cash value of the requested collateral substitution transfer in the base currency.
	SubstitutionRequirement *ActiveCurrencyAndAmount `xml:"SbstitnRqrmnt"`

	// Specifies if the collateral that is substituted was posted against the variation margin or the independent amount.
	CollateralSubstitutionType *CollateralSubstitutionType1Code `xml:"CollSbstitnTp"`

	// Identifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral5 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral3 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral5 `xml:"OthrColl,omitempty"`

	// Provides details on the identification of previously sent and/or received message(s), in case of updated substitution request.
	LinkedReferences *Reference17 `xml:"LkdRefs,omitempty"`
}

func (c *CollateralSubstitution4) SetCollateralSubstitutionSequence(value string) {
	c.CollateralSubstitutionSequence = (*CollateralSubstitutionSequence1Code)(&value)
}

func (c *CollateralSubstitution4) SetSubstitutionRequirement(value, currency string) {
	c.SubstitutionRequirement = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralSubstitution4) SetCollateralSubstitutionType(value string) {
	c.CollateralSubstitutionType = (*CollateralSubstitutionType1Code)(&value)
}

func (c *CollateralSubstitution4) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *CollateralSubstitution4) AddSecuritiesCollateral() *SecuritiesCollateral5 {
	newValue := new(SecuritiesCollateral5)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution4) AddCashCollateral() *CashCollateral3 {
	newValue := new(CashCollateral3)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution4) AddOtherCollateral() *OtherCollateral5 {
	newValue := new(OtherCollateral5)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution4) AddLinkedReferences() *Reference17 {
	c.LinkedReferences = new(Reference17)
	return c.LinkedReferences
}
