package model

// Set of elements identifying the dates related to the underlying transactions.
type TransactionDates1 struct {

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent (debtor's agent in case of a credit transfer, creditor's agent in case of a direct debit). This means - amongst others - that the account servicing agent has received the payment order and has applied checks as eg, authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Date on which the trade was executed.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Start date of the underlying transaction, such as a treasury transaction, an investment plan.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// End date of the underlying transaction, such as a treasury transaction, an investment plan.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Date and time of the underlying transaction.
	TransactionDateTime *ISODateTime `xml:"TxDtTm,omitempty"`

	// Proprietary date related to the underlying transaction.
	Proprietary []*ProprietaryDate1 `xml:"Prtry,omitempty"`
}

func (t *TransactionDates1) SetAcceptanceDateTime(value string) {
	t.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (t *TransactionDates1) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TransactionDates1) SetInterbankSettlementDate(value string) {
	t.InterbankSettlementDate = (*ISODate)(&value)
}

func (t *TransactionDates1) SetStartDate(value string) {
	t.StartDate = (*ISODate)(&value)
}

func (t *TransactionDates1) SetEndDate(value string) {
	t.EndDate = (*ISODate)(&value)
}

func (t *TransactionDates1) SetTransactionDateTime(value string) {
	t.TransactionDateTime = (*ISODateTime)(&value)
}

func (t *TransactionDates1) AddProprietary() *ProprietaryDate1 {
	newValue := new(ProprietaryDate1)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
