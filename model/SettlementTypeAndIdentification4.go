package model

// Provides transaction type and identification information.
type SettlementTypeAndIdentification4 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp,omitempty"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt,omitempty"`
}

func (s *SettlementTypeAndIdentification4) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification4) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification4) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}
