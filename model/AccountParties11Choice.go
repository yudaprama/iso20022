package model

// Information about a party's account.
type AccountParties11Choice struct {

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation15 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation15 `xml:"Trstee"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation15 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation15 `xml:"JntOwnr"`
}

func (a *AccountParties11Choice) AddPrimaryOwner() *InvestmentAccountOwnershipInformation15 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation15)
	return a.PrimaryOwner
}

func (a *AccountParties11Choice) AddTrustee() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties11Choice) AddNominee() *InvestmentAccountOwnershipInformation15 {
	a.Nominee = new(InvestmentAccountOwnershipInformation15)
	return a.Nominee
}

func (a *AccountParties11Choice) AddJointOwner() *InvestmentAccountOwnershipInformation15 {
	newValue := new(InvestmentAccountOwnershipInformation15)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}
