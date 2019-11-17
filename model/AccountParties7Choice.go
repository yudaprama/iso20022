package model

// Information about a party's account.
type AccountParties7Choice struct {

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation10 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation10 `xml:"Trstee"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation10 `xml:"CtdnForMnr"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation10 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation9 `xml:"JntOwnr"`
}

func (a *AccountParties7Choice) AddPrimaryOwner() *InvestmentAccountOwnershipInformation10 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation10)
	return a.PrimaryOwner
}

func (a *AccountParties7Choice) AddTrustee() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties7Choice) AddCustodianForMinor() *InvestmentAccountOwnershipInformation10 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation10)
	return a.CustodianForMinor
}

func (a *AccountParties7Choice) AddNominee() *InvestmentAccountOwnershipInformation10 {
	a.Nominee = new(InvestmentAccountOwnershipInformation10)
	return a.Nominee
}

func (a *AccountParties7Choice) AddJointOwner() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}
