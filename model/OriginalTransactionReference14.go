package model

// Set of elements used to provide reference information to the original message.
type OriginalTransactionReference14 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// Specifies the original message name identifier to which the message refers.
	MessageNameIdentification *Max35Text `xml:"MsgNmId,omitempty"`

	// Date and time at which the original message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`

	// Provides reference information to the original transaction.
	OriginalTransaction []*PaymentIdentification3 `xml:"OrgnlTx,omitempty"`
}

func (o *OriginalTransactionReference14) SetMessageIdentification(value string) {
	o.MessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalTransactionReference14) SetMessageNameIdentification(value string) {
	o.MessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalTransactionReference14) SetCreationDateTime(value string) {
	o.CreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalTransactionReference14) AddOriginalTransaction() *PaymentIdentification3 {
	newValue := new(PaymentIdentification3)
	o.OriginalTransaction = append(o.OriginalTransaction, newValue)
	return newValue
}
