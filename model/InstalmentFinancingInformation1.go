package model

// Information about result of a single instalment (financed or not) within an invoice.
type InstalmentFinancingInformation1 struct {

	// Progressive number of the single instalment related to an invoice.
	InstalmentSequenceIdentification *Max70Text `xml:"InstlmtSeqId"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstalmentTotalAmount *ActiveCurrencyAndAmount `xml:"InstlmtTtlAmt"`

	// Information about the financing result of one instalment.
	InstalmentFinancingResult *FinancingResult1 `xml:"InstlmtFincgRslt"`
}

func (i *InstalmentFinancingInformation1) SetInstalmentSequenceIdentification(value string) {
	i.InstalmentSequenceIdentification = (*Max70Text)(&value)
}

func (i *InstalmentFinancingInformation1) SetInstalmentTotalAmount(value, currency string) {
	i.InstalmentTotalAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InstalmentFinancingInformation1) AddInstalmentFinancingResult() *FinancingResult1 {
	i.InstalmentFinancingResult = new(FinancingResult1)
	return i.InstalmentFinancingResult
}
