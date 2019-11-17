package model

// Information about a party's account.
type AccountParties6Choice struct {

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation11 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation11 `xml:"Trstee"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation11 `xml:"CtdnForMnr"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation11 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation11 `xml:"JntOwnr"`
}

func (a *AccountParties6Choice) AddPrimaryOwner() *InvestmentAccountOwnershipInformation11 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation11)
	return a.PrimaryOwner
}

func (a *AccountParties6Choice) AddTrustee() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties6Choice) AddCustodianForMinor() *InvestmentAccountOwnershipInformation11 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation11)
	return a.CustodianForMinor
}

func (a *AccountParties6Choice) AddNominee() *InvestmentAccountOwnershipInformation11 {
	a.Nominee = new(InvestmentAccountOwnershipInformation11)
	return a.Nominee
}

func (a *AccountParties6Choice) AddJointOwner() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}
