package model

// Set of elements that identify the underlying transaction.
type TransactionReferences1 struct {

	// Point to point reference assigned by the instructing party of the underlying message.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// The account servicing institution's reference for the transaction.
	AccountServicerReference *Max35Text `xml:"AcctSvcrRef,omitempty"`

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: the  instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification assigned by the initiating party to unumbiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to carry on multiple references, the end-to-end identification must be carried on throughout the entire end-to-end chain.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Unique identification assigned by the first instructing agent to unambiguously identify the transaction and passed on, unchanged, throughout the entire interbank chain.
	//
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Reference of the direct debit mandate that has been signed between by the debtor and the creditor.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`

	// Identifies the cheque number.
	ChequeNumber *Max35Text `xml:"ChqNb,omitempty"`

	// Unique and unambiguous identifier for a payment instruction,  assigned by the clearing system.
	ClearingSystemReference *Max35Text `xml:"ClrSysRef,omitempty"`

	// Proprietary reference of an underlying transaction.
	Proprietary *ProprietaryReference1 `xml:"Prtry,omitempty"`
}

func (t *TransactionReferences1) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences1) SetAccountServicerReference(value string) {
	t.AccountServicerReference = (*Max35Text)(&value)
}

func (t *TransactionReferences1) SetInstructionIdentification(value string) {
	t.InstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences1) SetEndToEndIdentification(value string) {
	t.EndToEndIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences1) SetTransactionIdentification(value string) {
	t.TransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences1) SetMandateIdentification(value string) {
	t.MandateIdentification = (*Max35Text)(&value)
}

func (t *TransactionReferences1) SetChequeNumber(value string) {
	t.ChequeNumber = (*Max35Text)(&value)
}

func (t *TransactionReferences1) SetClearingSystemReference(value string) {
	t.ClearingSystemReference = (*Max35Text)(&value)
}

func (t *TransactionReferences1) AddProprietary() *ProprietaryReference1 {
	t.Proprietary = new(ProprietaryReference1)
	return t.Proprietary
}
