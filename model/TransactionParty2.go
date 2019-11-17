package model

// Set of elements used to provide information on the parties specific to the individual transaction.
type TransactionParty2 struct {

	// Party that initiated the payment that is reported in the entry.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification32 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification32 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry has been posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Party that plays an active role in planning and executing the transactions that create or liquidate investments of the investors assets, or that move the investor's assets from one investment to another. A trading party is a trade instructor, an investment decision-maker, a post trade administrator, or a trader. In the context of treasury, it is the party that negotiates and executes the treasury transaction.
	TradingParty *PartyIdentification32 `xml:"TradgPty,omitempty"`

	// Proprietary party related to the underlying transaction.
	Proprietary []*ProprietaryParty2 `xml:"Prtry,omitempty"`
}

func (t *TransactionParty2) AddInitiatingParty() *PartyIdentification32 {
	t.InitiatingParty = new(PartyIdentification32)
	return t.InitiatingParty
}

func (t *TransactionParty2) AddDebtor() *PartyIdentification32 {
	t.Debtor = new(PartyIdentification32)
	return t.Debtor
}

func (t *TransactionParty2) AddDebtorAccount() *CashAccount16 {
	t.DebtorAccount = new(CashAccount16)
	return t.DebtorAccount
}

func (t *TransactionParty2) AddUltimateDebtor() *PartyIdentification32 {
	t.UltimateDebtor = new(PartyIdentification32)
	return t.UltimateDebtor
}

func (t *TransactionParty2) AddCreditor() *PartyIdentification32 {
	t.Creditor = new(PartyIdentification32)
	return t.Creditor
}

func (t *TransactionParty2) AddCreditorAccount() *CashAccount16 {
	t.CreditorAccount = new(CashAccount16)
	return t.CreditorAccount
}

func (t *TransactionParty2) AddUltimateCreditor() *PartyIdentification32 {
	t.UltimateCreditor = new(PartyIdentification32)
	return t.UltimateCreditor
}

func (t *TransactionParty2) AddTradingParty() *PartyIdentification32 {
	t.TradingParty = new(PartyIdentification32)
	return t.TradingParty
}

func (t *TransactionParty2) AddProprietary() *ProprietaryParty2 {
	newValue := new(ProprietaryParty2)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
