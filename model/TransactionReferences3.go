package model

// Identifies the underlying transaction.
type TransactionReferences3 struct {

	// Point to point reference, as assigned by the instructing party of the underlying message.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Unique and unambiguous identifier for a cheque as assigned by the agent.
	ChequeNumber *Max35Text `xml:"ChqNb,omitempty"`

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Unambiguous identification of the securities transaction as known by the securities account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the securities transaction as known by the securities account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a securities transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the securities transaction assigned by the processor of the instruction other than the securities account owner, the securities account servicer and the market infrastructure.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Proprietary reference related to the underlying transaction.
	Proprietary []*ProprietaryReference1 `xml:"Prtry,omitempty"`
}

func (t *TransactionReferences3) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetAccountServicerReference(value string) {
	t.AccountServicerReference = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetPaymentInformationIdentification(value string) {
	t.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetInstructionIdentification(value string) {
	t.InstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetEndToEndIdentification(value string) {
	t.EndToEndIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetTransactionIdentification(value string) {
	t.TransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetMandateIdentification(value string) {
	t.MandateIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetChequeNumber(value string) {
	t.ChequeNumber = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetClearingSystemReference(value string) {
	t.ClearingSystemReference = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences3) SetProcessingIdentification(value string) {
	t.ProcessingIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences3) AddProprietary() *ProprietaryReference1 {
	newValue := new(ProprietaryReference1)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
