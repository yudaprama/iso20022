package model

// Party related to an investment account.
type AccountParties7 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties2Choice `xml:"PrncplAcctPty,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner *InvestmentAccountOwnershipInformation7 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary *InvestmentAccountOwnershipInformation7 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney *InvestmentAccountOwnershipInformation7 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian *InvestmentAccountOwnershipInformation7 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath *InvestmentAccountOwnershipInformation7 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator *InvestmentAccountOwnershipInformation7 `xml:"Admstr,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation7 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation7 `xml:"Sttlr,omitempty"`

	// An other type of party.
	OtherParty []*ExtendedParty4 `xml:"OthrPty,omitempty"`
}

func (a *AccountParties7) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties7) AddPrincipalAccountParty() *AccountParties2Choice {
	a.PrincipalAccountParty = new(AccountParties2Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties7) AddSecondaryOwner() *InvestmentAccountOwnershipInformation7 {
	a.SecondaryOwner = new(InvestmentAccountOwnershipInformation7)
	return a.SecondaryOwner
}

func (a *AccountParties7) AddBeneficiary() *InvestmentAccountOwnershipInformation7 {
	a.Beneficiary = new(InvestmentAccountOwnershipInformation7)
	return a.Beneficiary
}

func (a *AccountParties7) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation7 {
	a.PowerOfAttorney = new(InvestmentAccountOwnershipInformation7)
	return a.PowerOfAttorney
}

func (a *AccountParties7) AddLegalGuardian() *InvestmentAccountOwnershipInformation7 {
	a.LegalGuardian = new(InvestmentAccountOwnershipInformation7)
	return a.LegalGuardian
}

func (a *AccountParties7) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation7 {
	a.SuccessorOnDeath = new(InvestmentAccountOwnershipInformation7)
	return a.SuccessorOnDeath
}

func (a *AccountParties7) AddAdministrator() *InvestmentAccountOwnershipInformation7 {
	a.Administrator = new(InvestmentAccountOwnershipInformation7)
	return a.Administrator
}

func (a *AccountParties7) AddGranter() *InvestmentAccountOwnershipInformation7 {
	newValue := new(InvestmentAccountOwnershipInformation7)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties7) AddSettler() *InvestmentAccountOwnershipInformation7 {
	newValue := new(InvestmentAccountOwnershipInformation7)
	a.Settler = append(a.Settler, newValue)
	return newValue
}

func (a *AccountParties7) AddOtherParty() *ExtendedParty4 {
	newValue := new(ExtendedParty4)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}
