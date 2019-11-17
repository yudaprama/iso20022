package model

// Any party who is related to an investment account.
type AccountParties5 struct {

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation5 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation5 `xml:"Trstee"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation5 `xml:"CtdnForMnr"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation5 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation5 `xml:"JntOwnr"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation5 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation5 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation5 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation5 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation5 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator *InvestmentAccountOwnershipInformation5 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty2 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation5 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation5 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties5) AddPrimaryOwner() *InvestmentAccountOwnershipInformation5 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation5)
	return a.PrimaryOwner
}

func (a *AccountParties5) AddTrustee() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties5) AddCustodianForMinor() *InvestmentAccountOwnershipInformation5 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation5)
	return a.CustodianForMinor
}

func (a *AccountParties5) AddNominee() *InvestmentAccountOwnershipInformation5 {
	a.Nominee = new(InvestmentAccountOwnershipInformation5)
	return a.Nominee
}

func (a *AccountParties5) AddJointOwner() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}

func (a *AccountParties5) AddSecondaryOwner() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties5) AddBeneficiary() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties5) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties5) AddLegalGuardian() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties5) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties5) AddAdministrator() *InvestmentAccountOwnershipInformation5 {
	a.Administrator = new(InvestmentAccountOwnershipInformation5)
	return a.Administrator
}

func (a *AccountParties5) AddOtherParty() *ExtendedParty2 {
	newValue := new(ExtendedParty2)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties5) AddGranter() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties5) AddSettler() *InvestmentAccountOwnershipInformation5 {
	newValue := new(InvestmentAccountOwnershipInformation5)
	a.Settler = append(a.Settler, newValue)
	return newValue
}
