package model

// Choice between a payment term in a coded or free format.
type PaymentCodeOrOther1Choice struct {

	// Specifies the payment period in coded form and a number of days.
	PaymentCode *PaymentPeriod3 `xml:"PmtCd"`

	// Specifies the payment date as a fixed date.
	PaymentDueDate *ISODate `xml:"PmtDueDt"`

	// Specifies payment terms not present in a code list.
	OtherPaymentTerms *Max140Text `xml:"OthrPmtTerms"`
}

func (p *PaymentCodeOrOther1Choice) AddPaymentCode() *PaymentPeriod3 {
	p.PaymentCode = new(PaymentPeriod3)
	return p.PaymentCode
}

func (p *PaymentCodeOrOther1Choice) SetPaymentDueDate(value string) {
	p.PaymentDueDate = (*ISODate)(&value)
}

func (p *PaymentCodeOrOther1Choice) SetOtherPaymentTerms(value string) {
	p.OtherPaymentTerms = (*Max140Text)(&value)
}
