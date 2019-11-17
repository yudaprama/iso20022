package model

// Specifies the calculation and the resulting margin and independent amount needed to cover the risk exposure of one party versus another.
type MarginCall2 struct {

	// Provides additional information on the collateral account of the party delivering/receiving the collateral.
	CollateralAccountIdentification *CollateralAccount2 `xml:"CollAcctId,omitempty"`

	// Summation of the call amounts per margin type. It is provided for the purposes of carrying forward for future messages that are used to compare the margin call results to determine whether a call is agreed or full/partially disputed.
	MarginCallResult *MarginCallResult3 `xml:"MrgnCallRslt"`

	// Provides details about the margin calculation that would be due to party A.
	MarginDetailDueToA *MarginCall1 `xml:"MrgnDtlDueToA,omitempty"`

	// Provides details about the margin calculation that would be due to party B.
	MarginDetailDueToB *MarginCall1 `xml:"MrgnDtlDueToB,omitempty"`

	// Amount of expected margin that will be either delivered to party A by party B or recalled to party A from party B.
	RequirementDetailsDueToA *MarginRequirement1Choice `xml:"RqrmntDtlsDueToA,omitempty"`

	// Amount of expected margin that will be either delivered to party B by party A or recalled to party B from party A.
	RequirementDetailsDueToB *MarginRequirement1Choice `xml:"RqrmntDtlsDueToB,omitempty"`

	// Provides details about the type of collateral that will be either delivered to party A by party B or recalled to party A from party B.
	ExpectedCollateralDueToA *ExpectedCollateral2Choice `xml:"XpctdCollDueToA,omitempty"`

	// Provides details about the type of collateral that will be either delivered to party B by party A or recalled to party B from party A.
	ExpectedCollateralDueToB *ExpectedCollateral2Choice `xml:"XpctdCollDueToB,omitempty"`
}

func (m *MarginCall2) AddCollateralAccountIdentification() *CollateralAccount2 {
	m.CollateralAccountIdentification = new(CollateralAccount2)
	return m.CollateralAccountIdentification
}

func (m *MarginCall2) AddMarginCallResult() *MarginCallResult3 {
	m.MarginCallResult = new(MarginCallResult3)
	return m.MarginCallResult
}

func (m *MarginCall2) AddMarginDetailDueToA() *MarginCall1 {
	m.MarginDetailDueToA = new(MarginCall1)
	return m.MarginDetailDueToA
}

func (m *MarginCall2) AddMarginDetailDueToB() *MarginCall1 {
	m.MarginDetailDueToB = new(MarginCall1)
	return m.MarginDetailDueToB
}

func (m *MarginCall2) AddRequirementDetailsDueToA() *MarginRequirement1Choice {
	m.RequirementDetailsDueToA = new(MarginRequirement1Choice)
	return m.RequirementDetailsDueToA
}

func (m *MarginCall2) AddRequirementDetailsDueToB() *MarginRequirement1Choice {
	m.RequirementDetailsDueToB = new(MarginRequirement1Choice)
	return m.RequirementDetailsDueToB
}

func (m *MarginCall2) AddExpectedCollateralDueToA() *ExpectedCollateral2Choice {
	m.ExpectedCollateralDueToA = new(ExpectedCollateral2Choice)
	return m.ExpectedCollateralDueToA
}

func (m *MarginCall2) AddExpectedCollateralDueToB() *ExpectedCollateral2Choice {
	m.ExpectedCollateralDueToB = new(ExpectedCollateral2Choice)
	return m.ExpectedCollateralDueToB
}
