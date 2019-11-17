package model

// Set of elements providing information specific to the individual transaction(s) included in the message.
type TransactionParty1 struct {

	// Party initiating the payment to an agent. In the payment context, this can either be the debtor (in a credit transfer), the creditor (in a direct debit), or a party that initiates the payment on behalf of the debtor or creditor. In the context of treasury, the party that instructs the trading party to execute a treasury deal on its behalf.
	InitiatingParty *PartyIdentification8 `xml:"InitgPty,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification8 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor.
	DebtorAccount *CashAccount7 `xml:"DbtrAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification8 `xml:"UltmtDbtr,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification8 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor of the payment transaction.
	CreditorAccount *CashAccount7 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification8 `xml:"UltmtCdtr,omitempty"`

	// Party that plays an active role in planning and executing the transactions that create or liquidate investments of the investors assets, or that move the investor's assets from one investment to another. A trading party is a trade instructor, an investment decision-maker, a post trade administrator, or a trader. In the context of treasury, it is the party that negotiates and executes the treasury transaction.
	TradingParty *PartyIdentification8 `xml:"TradgPty,omitempty"`

	// Provides proprietary party information.
	Proprietary []*ProprietaryParty1 `xml:"Prtry,omitempty"`
}

func (t *TransactionParty1) AddInitiatingParty() *PartyIdentification8 {
	t.InitiatingParty = new(PartyIdentification8)
	return t.InitiatingParty
}

func (t *TransactionParty1) AddDebtor() *PartyIdentification8 {
	t.Debtor = new(PartyIdentification8)
	return t.Debtor
}

func (t *TransactionParty1) AddDebtorAccount() *CashAccount7 {
	t.DebtorAccount = new(CashAccount7)
	return t.DebtorAccount
}

func (t *TransactionParty1) AddUltimateDebtor() *PartyIdentification8 {
	t.UltimateDebtor = new(PartyIdentification8)
	return t.UltimateDebtor
}

func (t *TransactionParty1) AddCreditor() *PartyIdentification8 {
	t.Creditor = new(PartyIdentification8)
	return t.Creditor
}

func (t *TransactionParty1) AddCreditorAccount() *CashAccount7 {
	t.CreditorAccount = new(CashAccount7)
	return t.CreditorAccount
}

func (t *TransactionParty1) AddUltimateCreditor() *PartyIdentification8 {
	t.UltimateCreditor = new(PartyIdentification8)
	return t.UltimateCreditor
}

func (t *TransactionParty1) AddTradingParty() *PartyIdentification8 {
	t.TradingParty = new(PartyIdentification8)
	return t.TradingParty
}

func (t *TransactionParty1) AddProprietary() *ProprietaryParty1 {
	newValue := new(ProprietaryParty1)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
