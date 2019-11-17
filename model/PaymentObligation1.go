package model

// Payment obligation contracted between two financial institutions related to the financing of a commercial transaction.
type PaymentObligation1 struct {

	// Bank that has to pay under the obligation.
	ObligorBank *BICIdentification1 `xml:"OblgrBk"`

	// Bank that will be paid under the obligation.
	RecipientBank *BICIdentification1 `xml:"RcptBk"`

	// Maximum amount that will be paid under the obligation.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Maximum amount that will be paid under the obligation, expressed as a percentage of the purchase order net amount.
	Percentage *PercentageRate `xml:"Pctg"`

	// Amount of the charges taken by the obligor bank.
	ChargesAmount *CurrencyAndAmount `xml:"ChrgsAmt,omitempty"`

	// Amount of the charges expressed as a percentage of the amount paid by the obligor bank.
	ChargesPercentage *PercentageRate `xml:"ChrgsPctg,omitempty"`

	// Date at which the obligation will expire.
	ExpiryDate *ISODate `xml:"XpryDt"`

	// Country of which the law governs the bank payment obligation.
	ApplicableLaw *CountryCode `xml:"AplblLaw,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	PaymentTerms []*PaymentTerms2 `xml:"PmtTerms,omitempty"`

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SettlementTerms *SettlementTerms2 `xml:"SttlmTerms,omitempty"`
}

func (p *PaymentObligation1) AddObligorBank() *BICIdentification1 {
	p.ObligorBank = new(BICIdentification1)
	return p.ObligorBank
}

func (p *PaymentObligation1) AddRecipientBank() *BICIdentification1 {
	p.RecipientBank = new(BICIdentification1)
	return p.RecipientBank
}

func (p *PaymentObligation1) SetAmount(value, currency string) {
	p.Amount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentObligation1) SetPercentage(value string) {
	p.Percentage = (*PercentageRate)(&value)
}

func (p *PaymentObligation1) SetChargesAmount(value, currency string) {
	p.ChargesAmount = NewCurrencyAndAmount(value, currency)
}

func (p *PaymentObligation1) SetChargesPercentage(value string) {
	p.ChargesPercentage = (*PercentageRate)(&value)
}

func (p *PaymentObligation1) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISODate)(&value)
}

func (p *PaymentObligation1) SetApplicableLaw(value string) {
	p.ApplicableLaw = (*CountryCode)(&value)
}

func (p *PaymentObligation1) AddPaymentTerms() *PaymentTerms2 {
	newValue := new(PaymentTerms2)
	p.PaymentTerms = append(p.PaymentTerms, newValue)
	return newValue
}

func (p *PaymentObligation1) AddSettlementTerms() *SettlementTerms2 {
	p.SettlementTerms = new(SettlementTerms2)
	return p.SettlementTerms
}
