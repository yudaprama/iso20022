package model

// Any party who is related to an investment account.
type AccountParties4 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation4 `xml:"PmryOwnr,omitempty"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation4 `xml:"Trstee,omitempty"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation4 `xml:"CtdnForMnr,omitempty"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation4 `xml:"Nmnee,omitempty"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation4 `xml:"JntOwnr,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner *InvestmentAccountOwnershipInformation4 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary *InvestmentAccountOwnershipInformation4 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney *InvestmentAccountOwnershipInformation4 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian *InvestmentAccountOwnershipInformation4 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath *InvestmentAccountOwnershipInformation4 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator *InvestmentAccountOwnershipInformation4 `xml:"Admstr,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation4 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation4 `xml:"Sttlr,omitempty"`

	// An other type of party.
	OtherParty []*ExtendedParty1 `xml:"OthrPty,omitempty"`
}

func (a *AccountParties4) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties4) AddPrimaryOwner() *InvestmentAccountOwnershipInformation4 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation4)
	return a.PrimaryOwner
}

func (a *AccountParties4) AddTrustee() *InvestmentAccountOwnershipInformation4 {
	newValue := new(InvestmentAccountOwnershipInformation4)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties4) AddCustodianForMinor() *InvestmentAccountOwnershipInformation4 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation4)
	return a.CustodianForMinor
}

func (a *AccountParties4) AddNominee() *InvestmentAccountOwnershipInformation4 {
	a.Nominee = new(InvestmentAccountOwnershipInformation4)
	return a.Nominee
}

func (a *AccountParties4) AddJointOwner() *InvestmentAccountOwnershipInformation4 {
	newValue := new(InvestmentAccountOwnershipInformation4)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}

func (a *AccountParties4) AddSecondaryOwner() *InvestmentAccountOwnershipInformation4 {
	a.SecondaryOwner = new(InvestmentAccountOwnershipInformation4)
	return a.SecondaryOwner
}

func (a *AccountParties4) AddBeneficiary() *InvestmentAccountOwnershipInformation4 {
	a.Beneficiary = new(InvestmentAccountOwnershipInformation4)
	return a.Beneficiary
}

func (a *AccountParties4) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation4 {
	a.PowerOfAttorney = new(InvestmentAccountOwnershipInformation4)
	return a.PowerOfAttorney
}

func (a *AccountParties4) AddLegalGuardian() *InvestmentAccountOwnershipInformation4 {
	a.LegalGuardian = new(InvestmentAccountOwnershipInformation4)
	return a.LegalGuardian
}

func (a *AccountParties4) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation4 {
	a.SuccessorOnDeath = new(InvestmentAccountOwnershipInformation4)
	return a.SuccessorOnDeath
}

func (a *AccountParties4) AddAdministrator() *InvestmentAccountOwnershipInformation4 {
	a.Administrator = new(InvestmentAccountOwnershipInformation4)
	return a.Administrator
}

func (a *AccountParties4) AddGranter() *InvestmentAccountOwnershipInformation4 {
	newValue := new(InvestmentAccountOwnershipInformation4)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties4) AddSettler() *InvestmentAccountOwnershipInformation4 {
	newValue := new(InvestmentAccountOwnershipInformation4)
	a.Settler = append(a.Settler, newValue)
	return newValue
}

func (a *AccountParties4) AddOtherParty() *ExtendedParty1 {
	newValue := new(ExtendedParty1)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}
