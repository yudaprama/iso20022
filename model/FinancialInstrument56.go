package model

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument56 struct {

	// Identification of the security by an ISIN.
	Identification *SecurityIdentification25Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Additional information about the financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by the fund. For example, a unit may have a specific load structure, for example, front end or back end, an income policy,  for example, pay out or accumulate, or a trailer policy,  for example, with or without. Fund classes are typically denoted by a single character,  for example, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form of ownership, that is registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to the class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`

	// When an account at fund or security level is blocked, this specifies details on how the holding is blocked.
	BlockedHoldingDetails *BlockedHoldingDetails2 `xml:"BlckdHldgDtls,omitempty"`

	// Specifies whether the holdings in the account are eligible for pledging.
	Pledging *Eligible1Code `xml:"Pldgg,omitempty"`

	// Specifies whether the holdings in the account are used as collateral.
	Collateral *Collateral1Code `xml:"Coll,omitempty"`

	// Details of third party rights.
	ThirdPartyRights *ThirdPartyRights1 `xml:"ThrdPtyRghts,omitempty"`

	// Specifies if all the shares are owned exclusively by the fund company.
	FundOwnership *FundOwnership1Code `xml:"FndOwnrsh,omitempty"`

	// Specifies if the fund is intended for qualified investors.
	FundIntention *FundIntention1Code `xml:"FndIntntn,omitempty"`

	// Operational status of the fund.
	OperationalStatus *OperationalStatus1Code `xml:"OprlSts,omitempty"`
}

func (f *FinancialInstrument56) AddIdentification() *SecurityIdentification25Choice {
	f.Identification = new(SecurityIdentification25Choice)
	return f.Identification
}

func (f *FinancialInstrument56) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument56) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument56) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument56) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument56) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument56) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument56) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

func (f *FinancialInstrument56) AddBlockedHoldingDetails() *BlockedHoldingDetails2 {
	f.BlockedHoldingDetails = new(BlockedHoldingDetails2)
	return f.BlockedHoldingDetails
}

func (f *FinancialInstrument56) SetPledging(value string) {
	f.Pledging = (*Eligible1Code)(&value)
}

func (f *FinancialInstrument56) SetCollateral(value string) {
	f.Collateral = (*Collateral1Code)(&value)
}

func (f *FinancialInstrument56) AddThirdPartyRights() *ThirdPartyRights1 {
	f.ThirdPartyRights = new(ThirdPartyRights1)
	return f.ThirdPartyRights
}

func (f *FinancialInstrument56) SetFundOwnership(value string) {
	f.FundOwnership = (*FundOwnership1Code)(&value)
}

func (f *FinancialInstrument56) SetFundIntention(value string) {
	f.FundIntention = (*FundIntention1Code)(&value)
}

func (f *FinancialInstrument56) SetOperationalStatus(value string) {
	f.OperationalStatus = (*OperationalStatus1Code)(&value)
}
