package model

// Party associated with the account.
type AccountParties3Choice struct {

	// Single owner of the investment account or, when the ownership is split among several owners, the primary owner is the one giving its address and account details for the registration.
	PrimaryOwner *InvestmentAccountOwnershipInformation8 `xml:"PmryOwnr"`

	// Legal owners of the property. However, the beneficiary has the equitable or beneficial ownership.
	Trustee []*InvestmentAccountOwnershipInformation8 `xml:"Trstee"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation8 `xml:"CtdnForMnr"`

	// Entity named by the beneficial owner to act on its behalf, often to facilitate dealing, or to conceal the identity of the beneficiary.
	Nominee *InvestmentAccountOwnershipInformation8 `xml:"Nmnee"`

	// Co-owner of the investment account when the ownership is assigned to more than one party.
	JointOwner []*InvestmentAccountOwnershipInformation8 `xml:"JntOwnr"`
}

func (a *AccountParties3Choice) AddPrimaryOwner() *InvestmentAccountOwnershipInformation8 {
	a.PrimaryOwner = new(InvestmentAccountOwnershipInformation8)
	return a.PrimaryOwner
}

func (a *AccountParties3Choice) AddTrustee() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.Trustee = append(a.Trustee, newValue)
	return newValue
}

func (a *AccountParties3Choice) AddCustodianForMinor() *InvestmentAccountOwnershipInformation8 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation8)
	return a.CustodianForMinor
}

func (a *AccountParties3Choice) AddNominee() *InvestmentAccountOwnershipInformation8 {
	a.Nominee = new(InvestmentAccountOwnershipInformation8)
	return a.Nominee
}

func (a *AccountParties3Choice) AddJointOwner() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.JointOwner = append(a.JointOwner, newValue)
	return newValue
}
