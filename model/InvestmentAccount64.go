package model

// Information about a securities account and its characteristics.
type InvestmentAccount64 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Legal form of the fund, for example, UCITS, SICAV, OEIC, Unit Trust, and FCP.
	FundType *Max35Text `xml:"FndTp,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	SecurityDetails *FinancialInstrument55 `xml:"SctyDtls,omitempty"`

	// Owner of the account.
	AccountOwner *AccountOwner2Choice `xml:"AcctOwnr,omitempty"`

	// Intermediary or other party related to the management of the account.
	Intermediary []*Intermediary33 `xml:"Intrmy,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount64) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount64) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount64) SetFundType(value string) {
	i.FundType = (*Max35Text)(&value)
}

func (i *InvestmentAccount64) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount64) AddSecurityDetails() *FinancialInstrument55 {
	i.SecurityDetails = new(FinancialInstrument55)
	return i.SecurityDetails
}

func (i *InvestmentAccount64) AddAccountOwner() *AccountOwner2Choice {
	i.AccountOwner = new(AccountOwner2Choice)
	return i.AccountOwner
}

func (i *InvestmentAccount64) AddIntermediary() *Intermediary33 {
	newValue := new(Intermediary33)
	i.Intermediary = append(i.Intermediary, newValue)
	return newValue
}

func (i *InvestmentAccount64) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}
