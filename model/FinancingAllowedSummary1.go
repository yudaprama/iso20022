package model

// Summary information about amount financed.
type FinancingAllowedSummary1 struct {

	// Number of invoices/instalments financed.
	FinancedItemNumber *Number `xml:"FincdItmNb"`

	// Sum of the original total amounts of the invoices accepted for financing.
	TotalAcceptedItemsAmount *ActiveCurrencyAndAmount `xml:"TtlAccptdItmsAmt"`

	// Percentage rate applied to calculate the total amount financed related to the total amounts of the invoices accepted for financing. It represents the average percentage rate applied to all single invoice requests financed. It can be calculated as result of "TotalFinancedAmount" divided by "TotalAcceptedItemsAmount".
	AppliedPercentage *PercentageRate `xml:"ApldPctg,omitempty"`

	// Total amount financed, defined as the entire financed amount of the  requests.
	TotalFinancedAmount *ActiveCurrencyAndAmount `xml:"TtlFincdAmt"`

	// Set of dates (eg book date, credit date) related to the crediting of the financed amount.
	FinancingDateDetails *FinancingDateDetails1 `xml:"FincgDtDtls,omitempty"`

	// Unambiguous identification of the account, held by Financing Requestor, actually used for crediting the amount financed.
	CreditAccount *CashAccount7 `xml:"CdtAcct,omitempty"`

	// Unambiguous identification of the internal bank account actually used by First Agent to manage the line of credit granted to Financing Requestor.
	FinancingAccount *CashAccount7 `xml:"FincgAcct,omitempty"`
}

func (f *FinancingAllowedSummary1) SetFinancedItemNumber(value string) {
	f.FinancedItemNumber = (*Number)(&value)
}

func (f *FinancingAllowedSummary1) SetTotalAcceptedItemsAmount(value, currency string) {
	f.TotalAcceptedItemsAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancingAllowedSummary1) SetAppliedPercentage(value string) {
	f.AppliedPercentage = (*PercentageRate)(&value)
}

func (f *FinancingAllowedSummary1) SetTotalFinancedAmount(value, currency string) {
	f.TotalFinancedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancingAllowedSummary1) AddFinancingDateDetails() *FinancingDateDetails1 {
	f.FinancingDateDetails = new(FinancingDateDetails1)
	return f.FinancingDateDetails
}

func (f *FinancingAllowedSummary1) AddCreditAccount() *CashAccount7 {
	f.CreditAccount = new(CashAccount7)
	return f.CreditAccount
}

func (f *FinancingAllowedSummary1) AddFinancingAccount() *CashAccount7 {
	f.FinancingAccount = new(CashAccount7)
	return f.FinancingAccount
}
