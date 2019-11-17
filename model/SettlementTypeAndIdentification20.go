package model

// Overall process covering the trade and settlement transactions of financial instruments.
type SettlementTypeAndIdentification20 struct {

	// Specifies how the transaction is to be settled.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Reference of the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Identifies the intended settlement date.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt,omitempty"`
}

func (s *SettlementTypeAndIdentification20) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification20) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification20) AddSettlementDate() *DateAndDateTimeChoice {
	s.SettlementDate = new(DateAndDateTimeChoice)
	return s.SettlementDate
}
