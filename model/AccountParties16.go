package model

// Information about a party's account.
type AccountParties16 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties11Choice `xml:"PrncplAcctPty,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation15 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation15 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation15 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation15 `xml:"LglGuardn,omitempty"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation15 `xml:"CtdnForMnr,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation15 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation15 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty12 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation15 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation15 `xml:"Sttlr,omitempty"`

	// Party that makes, or participates in the making of, decisions that affect the whole, or a substantial part, of the business of a customer of a reporting entity or that has the capacity to affect significantly the financial standing of a customer of a reporting entity. Typically, this is a controlling person of a corporate (ownership type CORP).
	SeniorManagingOfficial []*InvestmentAccountOwnershipInformation15 `xml:"SnrMggOffcl,omitempty"`

	// Person appointed under the trust instrument to direct or restrain the trustees in relation to their administration of the trust. Typically, this is a controlling person of a trust (ownership type TRUS) or other non-individual organisation (ownership type ONIS).
	Protector []*InvestmentAccountOwnershipInformation15 `xml:"Prtctr,omitempty"`

	// Party for which shares are to be registered.
	RegisteredShareholderName *RegisteredShareholderName1Choice `xml:"RegdShrhldrNm,omitempty"`
}

func (a *AccountParties16) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties16) AddPrincipalAccountParty() *AccountParties11Choice {
	a.PrincipalAccountParty = new(AccountParties11Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties16) AddSecondaryOwner() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties16) AddBeneficiary() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties16) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties16) AddLegalGuardian() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties16) AddCustodianForMinor() *InvestmentAccountOwnershipInformation15 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation15)
	return a.CustodianForMinor
}

func (a *AccountParties16) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties16) AddAdministrator() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties16) AddOtherParty() *ExtendedParty12 {
	newValue := new(ExtendedParty12)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties16) AddGranter() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties16) AddSettlor() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}

func (a *AccountParties16) AddSeniorManagingOfficial() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.SeniorManagingOfficial = append(a.SeniorManagingOfficial, newValue)
	return newValue
}

func (a *AccountParties16) AddProtector() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Protector = append(a.Protector, newValue)
	return newValue
}

func (a *AccountParties16) AddRegisteredShareholderName() *RegisteredShareholderName1Choice {
	a.RegisteredShareholderName = new(RegisteredShareholderName1Choice)
	return a.RegisteredShareholderName
}
