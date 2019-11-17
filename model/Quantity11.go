package model

// Details on the quantity, account and other related information involved in a transaction.
type Quantity11 struct {

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Denomination of the security to be received or delivered.
	DenominationChoice *Max210Text `xml:"DnmtnChc,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber []*SecuritiesCertificate4 `xml:"CertNb,omitempty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown30 `xml:"QtyBrkdwn,omitempty"`
}

func (q *Quantity11) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	q.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return q.SettlementQuantity
}

func (q *Quantity11) SetDenominationChoice(value string) {
	q.DenominationChoice = (*Max210Text)(&value)
}

func (q *Quantity11) AddCertificateNumber() *SecuritiesCertificate4 {
	newValue := new(SecuritiesCertificate4)
	q.CertificateNumber = append(q.CertificateNumber, newValue)
	return newValue
}

func (q *Quantity11) AddQuantityBreakdown() *QuantityBreakdown30 {
	newValue := new(QuantityBreakdown30)
	q.QuantityBreakdown = append(q.QuantityBreakdown, newValue)
	return newValue
}
