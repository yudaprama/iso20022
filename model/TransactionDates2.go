package model

// Set of elements used to provide information on the dates related to the underlying individual transaction.
type TransactionDates2 struct {

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Identifies when an amount of money should have contractually been credited or debited the account versus when the amount of money was actually settled (debited/credited) on the cash account.
	TradeActivityContractualSettlementDate *ISODate `xml:"TradActvtyCtrctlSttlmDt,omitempty"`

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
	Proprietary []*ProprietaryDate2 `xml:"Prtry,omitempty"`
}

func (t *TransactionDates2) SetAcceptanceDateTime(value string) {
	t.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (t *TransactionDates2) SetTradeActivityContractualSettlementDate(value string) {
	t.TradeActivityContractualSettlementDate = (*ISODate)(&value)
}

func (t *TransactionDates2) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TransactionDates2) SetInterbankSettlementDate(value string) {
	t.InterbankSettlementDate = (*ISODate)(&value)
}

func (t *TransactionDates2) SetStartDate(value string) {
	t.StartDate = (*ISODate)(&value)
}

func (t *TransactionDates2) SetEndDate(value string) {
	t.EndDate = (*ISODate)(&value)
}

func (t *TransactionDates2) SetTransactionDateTime(value string) {
	t.TransactionDateTime = (*ISODateTime)(&value)
}

func (t *TransactionDates2) AddProprietary() *ProprietaryDate2 {
	newValue := new(ProprietaryDate2)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
