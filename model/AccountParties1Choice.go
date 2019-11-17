package model

// Party associated with the account.
type AccountParties1Choice struct {

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation6 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation6 `xml:"Trstee"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation6 `xml:"CtdnForMnr"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation6 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation6 `xml:"JntOwnr"`
}

func (a *AccountParties1Choice) AddPrimaryOwner() *InvestmentAccountOwnershipInformation6 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation6)
	return a.PrimaryOwner
}

func (a *AccountParties1Choice) AddTrustee() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties1Choice) AddCustodianForMinor() *InvestmentAccountOwnershipInformation6 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation6)
	return a.CustodianForMinor
}

func (a *AccountParties1Choice) AddNominee() *InvestmentAccountOwnershipInformation6 {
	a.Nominee = new(InvestmentAccountOwnershipInformation6)
	return a.Nominee
}

func (a *AccountParties1Choice) AddJointOwner() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}
