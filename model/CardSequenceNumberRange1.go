package model

// Range of sequence numbers related to card transactions.
type CardSequenceNumberRange1 struct {

	// CardSequenceNumberRange1:FirstTransactionSequenceNumberMessage element to be finalised once feedback from Card SEG has been received.
	FirstTransaction *Max35Text `xml:"FrstTx,omitempty"`

	// CardSequenceNumberRange1:LastTransactionSequenceNumberMessage element to be finalised once feedback from Card SEG has been received.
	LastTransaction *Max35Text `xml:"LastTx,omitempty"`
}

func (c *CardSequenceNumberRange1) SetFirstTransaction(value string) {
	c.FirstTransaction = (*Max35Text)(&value)
}

func (c *CardSequenceNumberRange1) SetLastTransaction(value string) {
	c.LastTransaction = (*Max35Text)(&value)
}
