package model

// Identifies the details of the transaction.
type TransactionDetails101 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *SettlementTypeAndIdentification22 `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *RestrictedFINXMax16Text `xml:"OthrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionDetails *TransactionDetails83 `xml:"TxDtls,omitempty"`
}

func (t *TransactionDetails101) AddAccountOwnerTransactionIdentification() *SettlementTypeAndIdentification22 {
	t.AccountOwnerTransactionIdentification = new(SettlementTypeAndIdentification22)
	return t.AccountOwnerTransactionIdentification
}

func (t *TransactionDetails101) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails101) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails101) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails101) SetOtherTransactionIdentification(value string) {
	t.OtherTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails101) AddAccountOwner() *PartyIdentification119 {
	t.AccountOwner = new(PartyIdentification119)
	return t.AccountOwner
}

func (t *TransactionDetails101) AddSafekeepingAccount() *SecuritiesAccount30 {
	t.SafekeepingAccount = new(SecuritiesAccount30)
	return t.SafekeepingAccount
}

func (t *TransactionDetails101) AddTransactionDetails() *TransactionDetails83 {
	t.TransactionDetails = new(TransactionDetails83)
	return t.TransactionDetails
}
