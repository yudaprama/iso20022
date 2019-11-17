package model

// Information about a party's account.
type AccountParties15 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties10Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation14 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation14 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation14 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation14 `xml:"LglGuardn,omitempty"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation14 `xml:"CtdnForMnr,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation14 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation14 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty11 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation14 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation14 `xml:"Sttlr,omitempty"`

	// Party that makes, or participates in the making of, decisions that affect the whole, or a substantial part, of the business of a customer of a reporting entity or that has the capacity to affect significantly the financial standing of a customer of a reporting entity. Typically, this is a controlling person of a corporate (ownership type CORP).
	SeniorManagingOfficial []*InvestmentAccountOwnershipInformation14 `xml:"SnrMggOffcl,omitempty"`

	// Person appointed under the trust instrument to direct or restrain the trustees in relation to their administration of the trust. Typically, this is a controlling person of a trust (ownership type TRUS) or other non-individual organisation (ownership type ONIS).
	Protector []*InvestmentAccountOwnershipInformation14 `xml:"Prtctr,omitempty"`

	// Party that registers its name with the issuer and the name used for the registration.
	RegisteredShareholderName *RegisteredShareholderName1Choice `xml:"RegdShrhldrNm,omitempty"`
}

func (a *AccountParties15) AddPrincipalAccountParty() *AccountParties10Choice {
	a.PrincipalAccountParty = new(AccountParties10Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties15) AddSecondaryOwner() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties15) AddBeneficiary() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties15) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties15) AddLegalGuardian() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties15) AddCustodianForMinor() *InvestmentAccountOwnershipInformation14 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation14)
	return a.CustodianForMinor
}

func (a *AccountParties15) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties15) AddAdministrator() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties15) AddOtherParty() *ExtendedParty11 {
	newValue := new(ExtendedParty11)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties15) AddGranter() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties15) AddSettlor() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}

func (a *AccountParties15) AddSeniorManagingOfficial() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.SeniorManagingOfficial = append(a.SeniorManagingOfficial, newValue)
	return newValue
}

func (a *AccountParties15) AddProtector() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Protector = append(a.Protector, newValue)
	return newValue
}

func (a *AccountParties15) AddRegisteredShareholderName() *RegisteredShareholderName1Choice {
	a.RegisteredShareholderName = new(RegisteredShareholderName1Choice)
	return a.RegisteredShareholderName
}
