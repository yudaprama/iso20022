package model

// Information about a party's account.
type AccountParties8Choice struct {

	// Single owner of the account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation12 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation12 `xml:"Trstee"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation12 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation12 `xml:"JntOwnr"`
}

func (a *AccountParties8Choice) AddPrimaryOwner() *InvestmentAccountOwnershipInformation12 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation12)
	return a.PrimaryOwner
}

func (a *AccountParties8Choice) AddTrustee() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties8Choice) AddNominee() *InvestmentAccountOwnershipInformation12 {
	a.Nominee = new(InvestmentAccountOwnershipInformation12)
	return a.Nominee
}

func (a *AccountParties8Choice) AddJointOwner() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}
