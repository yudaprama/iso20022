package model

// Specifies the unique transaction identifier (UTI) that was created at the time a transaction was first executed, shared with all registered entities and counterparties involved in the transaction, and used to track that particular transaction during its lifetime and optionally, the prior unique transaction identifier (PUTI). These identifiers can also be known as the Unique Swap Identifier (USI) or the Prior Unique Swap Identifier (PUSI).
type UniqueTransactionIdentifier2 struct {

	// Unique transaction identifier will be created at the time a transaction is first executed, shared with all registered entities and counterparties involved in the transaction, and used to track that particular transaction during its lifetime. This identifier can also be known as the Unique Swap Identifier (USI).
	UniqueTransactionIdentifier *Max52Text `xml:"UnqTxIdr"`

	// Prior unique transaction identifier specifies the previous unique transaction identifier (UTI) that was created at the time the transaction was executed. This identifier can also be known as the Prior Unique Swap Identifier (PUSI).
	PriorUniqueTransactionIdentifier []*Max52Text `xml:"PrrUnqTxIdr,omitempty"`
}

func (u *UniqueTransactionIdentifier2) SetUniqueTransactionIdentifier(value string) {
	u.UniqueTransactionIdentifier = (*Max52Text)(&value)
}

func (u *UniqueTransactionIdentifier2) AddPriorUniqueTransactionIdentifier(value string) {
	u.PriorUniqueTransactionIdentifier = append(u.PriorUniqueTransactionIdentifier, (*Max52Text)(&value))
}
