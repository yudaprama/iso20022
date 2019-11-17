package model

// Information on the occurred settlement time(s) of the payment transaction.
type SettlementDateTimeIndication1 struct {

	// Date and time at which a payment has been debited at the transaction administrator. In the case of TARGET, the date and time at which the payment has been debited at the central bank, expressed in Central European Time (CET).
	DebitDateTime *ISODateTime `xml:"DbtDtTm,omitempty"`

	// Date and time at which a payment has been credited at the transaction administrator. In the case of TARGET, the date and time at which the payment has been credited at the receiving central bank, expressed in Central European Time (CET).
	CreditDateTime *ISODateTime `xml:"CdtDtTm,omitempty"`
}

func (s *SettlementDateTimeIndication1) SetDebitDateTime(value string) {
	s.DebitDateTime = (*ISODateTime)(&value)
}

func (s *SettlementDateTimeIndication1) SetCreditDateTime(value string) {
	s.CreditDateTime = (*ISODateTime)(&value)
}
