package model

// Provides transaction type and identification information.
type SettlementTypeAndIdentification22 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *RestrictedFINXMax16Text `xml:"TxId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`
}

func (s *SettlementTypeAndIdentification22) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndIdentification22) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification22) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}
