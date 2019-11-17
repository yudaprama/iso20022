package model

// Provides further details on the parties specific to the individual transaction.
type TransactionParties3 struct {

	// Party that initiated the payment that is reported in the entry.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry has been posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that plays an active role in planning and executing the transactions that create or liquidate investments of the investors assets, or that move the investor's assets from one investment to another. A trading party is a trade instructor, an investment decision-maker, a post trade administrator, or a trader. In the context of treasury, it is the party that negotiates and executes the treasury transaction.
	TradingParty *PartyIdentification43 `xml:"TradgPty,omitempty"`

	// Proprietary party related to the underlying transaction.
	Proprietary []*ProprietaryParty3 `xml:"Prtry,omitempty"`
}

func (t *TransactionParties3) AddInitiatingParty() *PartyIdentification43 {
	t.InitiatingParty = new(PartyIdentification43)
	return t.InitiatingParty
}

func (t *TransactionParties3) AddDebtor() *PartyIdentification43 {
	t.Debtor = new(PartyIdentification43)
	return t.Debtor
}

func (t *TransactionParties3) AddDebtorAccount() *CashAccount24 {
	t.DebtorAccount = new(CashAccount24)
	return t.DebtorAccount
}

func (t *TransactionParties3) AddUltimateDebtor() *PartyIdentification43 {
	t.UltimateDebtor = new(PartyIdentification43)
	return t.UltimateDebtor
}

func (t *TransactionParties3) AddCreditor() *PartyIdentification43 {
	t.Creditor = new(PartyIdentification43)
	return t.Creditor
}

func (t *TransactionParties3) AddCreditorAccount() *CashAccount24 {
	t.CreditorAccount = new(CashAccount24)
	return t.CreditorAccount
}

func (t *TransactionParties3) AddUltimateCreditor() *PartyIdentification43 {
	t.UltimateCreditor = new(PartyIdentification43)
	return t.UltimateCreditor
}

func (t *TransactionParties3) AddTradingParty() *PartyIdentification43 {
	t.TradingParty = new(PartyIdentification43)
	return t.TradingParty
}

func (t *TransactionParties3) AddProprietary() *ProprietaryParty3 {
	newValue := new(ProprietaryParty3)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
