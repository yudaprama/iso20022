package model

// Information about a party's account.
type AccountParties9Choice struct {

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation13 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation13 `xml:"Trstee"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation13 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation13 `xml:"JntOwnr"`
}

func (a *AccountParties9Choice) AddPrimaryOwner() *InvestmentAccountOwnershipInformation13 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation13)
	return a.PrimaryOwner
}

func (a *AccountParties9Choice) AddTrustee() *InvestmentAccountOwnershipInformation13 {
	newValue := new(InvestmentAccountOwnershipInformation13)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties9Choice) AddNominee() *InvestmentAccountOwnershipInformation13 {
	a.Nominee = new(InvestmentAccountOwnershipInformation13)
	return a.Nominee
}

func (a *AccountParties9Choice) AddJointOwner() *InvestmentAccountOwnershipInformation13 {
	newValue := new(InvestmentAccountOwnershipInformation13)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}
