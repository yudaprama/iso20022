package model

// Valuation information of the portfolio.
type TotalPortfolioValuation1 struct {

	// Total value of the portfolio (sum of the assets, liabilities and unrealised gain/loss) calculated according to the accounting rules.
	TotalPortfolioValue *AmountAndDirection30 `xml:"TtlPrtflVal"`

	// Previous total value of the portfolio.
	PreviousTotalPortfolioValue *AmountAndDirection30 `xml:"PrvsTtlPrtflVal,omitempty"`

	// Difference or change between the previous total portfolio value and the current total portfolio value.
	TotalPortfolioValueChange *AmountAndRate2 `xml:"TtlPrtflValChng,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *AmountAndDirection30 `xml:"TtlBookVal"`

	// Previous net asset on balance sheet.
	PreviousTotalBookValue *AmountAndDirection30 `xml:"PrvsTtlBookVal,omitempty"`

	// Difference or change between the previous net asset on balance sheet and the current net asset on balance sheet.
	TotalBookValueChange *AmountAndRate2 `xml:"TtlBookValChng,omitempty"`

	// Total receipts attributable to the portfolio.
	TotalReceipts *AmountAndDirection30 `xml:"TtlRcts,omitempty"`

	// Total disbursements attributable to the portfolio.
	TotalDisbursements *AmountAndDirection30 `xml:"TtlDsbrsmnts,omitempty"`

	// Income attributable to the portfolio.
	IncomeReceived *AmountAndDirection30 `xml:"IncmRcvd,omitempty"`

	// Expenses attributable to the portfolio
	ExpensesPaid *AmountAndDirection30 `xml:"ExpnssPd,omitempty"`

	// Difference between the holding value and the book value of the portfolio.
	UnrealisedGainOrLoss *AmountAndDirection31 `xml:"UrlsdGnOrLoss,omitempty"`

	// Difference between the realised value caused by the actual trade/re-evaluation and the book value of the portfolio.
	RealisedGainOrLoss *AmountAndDirection31 `xml:"RealsdGnOrLoss,omitempty"`

	// Accrued income.
	AccruedIncome *AmountAndDirection30 `xml:"AcrdIncm,omitempty"`

	// Valuation information of the investment fund or investment fund share class.
	InvestmentFundDetails []*InvestmentFund1 `xml:"InvstmtFndDtls,omitempty"`
}

func (t *TotalPortfolioValuation1) AddTotalPortfolioValue() *AmountAndDirection30 {
	t.TotalPortfolioValue = new(AmountAndDirection30)
	return t.TotalPortfolioValue
}

func (t *TotalPortfolioValuation1) AddPreviousTotalPortfolioValue() *AmountAndDirection30 {
	t.PreviousTotalPortfolioValue = new(AmountAndDirection30)
	return t.PreviousTotalPortfolioValue
}

func (t *TotalPortfolioValuation1) AddTotalPortfolioValueChange() *AmountAndRate2 {
	t.TotalPortfolioValueChange = new(AmountAndRate2)
	return t.TotalPortfolioValueChange
}

func (t *TotalPortfolioValuation1) AddTotalBookValue() *AmountAndDirection30 {
	t.TotalBookValue = new(AmountAndDirection30)
	return t.TotalBookValue
}

func (t *TotalPortfolioValuation1) AddPreviousTotalBookValue() *AmountAndDirection30 {
	t.PreviousTotalBookValue = new(AmountAndDirection30)
	return t.PreviousTotalBookValue
}

func (t *TotalPortfolioValuation1) AddTotalBookValueChange() *AmountAndRate2 {
	t.TotalBookValueChange = new(AmountAndRate2)
	return t.TotalBookValueChange
}

func (t *TotalPortfolioValuation1) AddTotalReceipts() *AmountAndDirection30 {
	t.TotalReceipts = new(AmountAndDirection30)
	return t.TotalReceipts
}

func (t *TotalPortfolioValuation1) AddTotalDisbursements() *AmountAndDirection30 {
	t.TotalDisbursements = new(AmountAndDirection30)
	return t.TotalDisbursements
}

func (t *TotalPortfolioValuation1) AddIncomeReceived() *AmountAndDirection30 {
	t.IncomeReceived = new(AmountAndDirection30)
	return t.IncomeReceived
}

func (t *TotalPortfolioValuation1) AddExpensesPaid() *AmountAndDirection30 {
	t.ExpensesPaid = new(AmountAndDirection30)
	return t.ExpensesPaid
}

func (t *TotalPortfolioValuation1) AddUnrealisedGainOrLoss() *AmountAndDirection31 {
	t.UnrealisedGainOrLoss = new(AmountAndDirection31)
	return t.UnrealisedGainOrLoss
}

func (t *TotalPortfolioValuation1) AddRealisedGainOrLoss() *AmountAndDirection31 {
	t.RealisedGainOrLoss = new(AmountAndDirection31)
	return t.RealisedGainOrLoss
}

func (t *TotalPortfolioValuation1) AddAccruedIncome() *AmountAndDirection30 {
	t.AccruedIncome = new(AmountAndDirection30)
	return t.AccruedIncome
}

func (t *TotalPortfolioValuation1) AddInvestmentFundDetails() *InvestmentFund1 {
	newValue := new(InvestmentFund1)
	t.InvestmentFundDetails = append(t.InvestmentFundDetails, newValue)
	return newValue
}
