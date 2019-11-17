package model

// Provides reference information to the original message.
type OriginalTransactionReference18 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// Specifies the original message name identifier to which the message refers.
	MessageNameIdentification *Max35Text `xml:"MsgNmId,omitempty"`

	// Date and time at which the original message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`

	// Provides reference information to the original transaction.
	OriginalTransaction []*PaymentIdentification4 `xml:"OrgnlTx,omitempty"`
}

func (o *OriginalTransactionReference18) SetMessageIdentification(value string) {
	o.MessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalTransactionReference18) SetMessageNameIdentification(value string) {
	o.MessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalTransactionReference18) SetCreationDateTime(value string) {
	o.CreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalTransactionReference18) AddOriginalTransaction() *PaymentIdentification4 {
	newValue := new(PaymentIdentification4)
	o.OriginalTransaction = append(o.OriginalTransaction, newValue)
	return newValue
}
