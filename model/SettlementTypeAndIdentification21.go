package model

// Overall process covering the trade and settlement transactions of financial instruments.
type SettlementTypeAndIdentification21 struct {

	// Specifies how the transaction is to be settled.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Reference of the transaction.
	TransactionIdentification *RestrictedFINXMax16Text `xml:"TxId"`

	// Identifies the intended settlement date.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt,omitempty"`
}

func (s *SettlementTypeAndIdentification21) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification21) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndIdentification21) AddSettlementDate() *DateAndDateTimeChoice {
	s.SettlementDate = new(DateAndDateTimeChoice)
	return s.SettlementDate
}
