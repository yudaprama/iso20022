package model

// Party associated with the account.
type AccountParties2Choice struct {

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation7 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation7 `xml:"Trstee"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation7 `xml:"CtdnForMnr"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation7 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation7 `xml:"JntOwnr"`
}

func (a *AccountParties2Choice) AddPrimaryOwner() *InvestmentAccountOwnershipInformation7 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation7)
	return a.PrimaryOwner
}

func (a *AccountParties2Choice) AddTrustee() *InvestmentAccountOwnershipInformation7 {
	newValue := new(InvestmentAccountOwnershipInformation7)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties2Choice) AddCustodianForMinor() *InvestmentAccountOwnershipInformation7 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation7)
	return a.CustodianForMinor
}

func (a *AccountParties2Choice) AddNominee() *InvestmentAccountOwnershipInformation7 {
	a.Nominee = new(InvestmentAccountOwnershipInformation7)
	return a.Nominee
}

func (a *AccountParties2Choice) AddJointOwner() *InvestmentAccountOwnershipInformation7 {
	newValue := new(InvestmentAccountOwnershipInformation7)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}
