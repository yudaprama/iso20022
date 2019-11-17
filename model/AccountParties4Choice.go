package model

// Party associated with the account.
type AccountParties4Choice struct {

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation9 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation9 `xml:"Trstee"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation9 `xml:"CtdnForMnr"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation9 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation9 `xml:"JntOwnr"`
}

func (a *AccountParties4Choice) AddPrimaryOwner() *InvestmentAccountOwnershipInformation9 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation9)
	return a.PrimaryOwner
}

func (a *AccountParties4Choice) AddTrustee() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties4Choice) AddCustodianForMinor() *InvestmentAccountOwnershipInformation9 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation9)
	return a.CustodianForMinor
}

func (a *AccountParties4Choice) AddNominee() *InvestmentAccountOwnershipInformation9 {
	a.Nominee = new(InvestmentAccountOwnershipInformation9)
	return a.Nominee
}

func (a *AccountParties4Choice) AddJointOwner() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}
