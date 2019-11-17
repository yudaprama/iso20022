package model

// Information about a party's account.
type AccountParties10Choice struct {

	// Single owner of the account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation14 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation14 `xml:"Trstee"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation14 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation14 `xml:"JntOwnr"`
}

func (a *AccountParties10Choice) AddPrimaryOwner() *InvestmentAccountOwnershipInformation14 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation14)
	return a.PrimaryOwner
}

func (a *AccountParties10Choice) AddTrustee() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties10Choice) AddNominee() *InvestmentAccountOwnershipInformation14 {
	a.Nominee = new(InvestmentAccountOwnershipInformation14)
	return a.Nominee
}

func (a *AccountParties10Choice) AddJointOwner() *InvestmentAccountOwnershipInformation14 {
	newValue := new(InvestmentAccountOwnershipInformation14)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}
