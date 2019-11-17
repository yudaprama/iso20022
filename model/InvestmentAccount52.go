package model

// Information about a securities account and its characteristics.
type InvestmentAccount52 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Legal form of the fund, for example, UCITS, SICAV, OEIC, Unit Trust, and FCP.
	FundType *Max35Text `xml:"FndTp,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	SecurityDetails *FinancialInstrument45 `xml:"SctyDtls,omitempty"`

	// Owner of the account.
	AccountOwner *AccountOwner1Choice `xml:"AcctOwnr,omitempty"`

	// Intermediary or other party related to the management of the account. In some markets, when this intermediary is a party acting on behalf of the investor for which it has opened an account at, for example, a central securities depository or international central securities depository, this party is known by the investor as the 'account controller'.
	Intermediary []*Intermediary33 `xml:"Intrmy,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount52) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount52) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount52) SetFundType(value string) {
	i.FundType = (*Max35Text)(&value)
}

func (i *InvestmentAccount52) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount52) AddSecurityDetails() *FinancialInstrument45 {
	i.SecurityDetails = new(FinancialInstrument45)
	return i.SecurityDetails
}

func (i *InvestmentAccount52) AddAccountOwner() *AccountOwner1Choice {
	i.AccountOwner = new(AccountOwner1Choice)
	return i.AccountOwner
}

func (i *InvestmentAccount52) AddIntermediary() *Intermediary33 {
	newValue := new(Intermediary33)
	i.Intermediary = append(i.Intermediary, newValue)
	return newValue
}

func (i *InvestmentAccount52) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}
