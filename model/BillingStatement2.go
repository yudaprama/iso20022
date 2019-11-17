package model

// Details of the statement reporting the bank services billing.
type BillingStatement2 struct {

	// Identification of the customer billing statement.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Date range between the start date and the end date for which the statement is issued.
	FromToDate *DatePeriod1 `xml:"FrToDt"`

	// Date the statement message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Defines the status of the statement.
	Status *BillingStatementStatus1Code `xml:"Sts"`

	// Specifies the details of the account characteristics.
	AccountCharacteristics *CashAccountCharacteristics2 `xml:"AcctChrtcs"`

	// Identifies the non tax per annum rate and factor values used within the statement along with any time dependent charge basis.
	RateData []*BillingRate1 `xml:"RateData,omitempty"`

	// Specifies details related to currency exchange data.
	CurrencyExchange []*CurrencyExchange6 `xml:"CcyXchg,omitempty"`

	// Identifies the average value of balances held within the statement period.
	Balance []*BillingBalance1 `xml:"Bal,omitempty"`

	// Identifies the set of values and totals that are used to provide compensation information, service and tax totals.
	Compensation []*BillingCompensation1 `xml:"Compstn,omitempty"`

	// Specifies the values used for every line item service in the statement.
	Service []*BillingService2 `xml:"Svc,omitempty"`

	// Tax region(s) that levy a tax on the services within this statement.
	TaxRegion []*BillingTaxRegion1 `xml:"TaxRgn,omitempty"`

	// One or more sections that identify balance or float adjustments to the account. They can reflect either adjustments to the current statement or adjustments to statements from prior reporting periods.
	BalanceAdjustment []*BalanceAdjustment1 `xml:"BalAdjstmnt,omitempty"`

	// One or more sections that identify line item service adjustments to the account. They reflect adjustments to statements from prior reporting periods.
	ServiceAdjustment []*BillingServiceAdjustment1 `xml:"SvcAdjstmnt,omitempty"`
}

func (b *BillingStatement2) SetStatementIdentification(value string) {
	b.StatementIdentification = (*Max35Text)(&value)
}

func (b *BillingStatement2) AddFromToDate() *DatePeriod1 {
	b.FromToDate = new(DatePeriod1)
	return b.FromToDate
}

func (b *BillingStatement2) SetCreationDateTime(value string) {
	b.CreationDateTime = (*ISODateTime)(&value)
}

func (b *BillingStatement2) SetStatus(value string) {
	b.Status = (*BillingStatementStatus1Code)(&value)
}

func (b *BillingStatement2) AddAccountCharacteristics() *CashAccountCharacteristics2 {
	b.AccountCharacteristics = new(CashAccountCharacteristics2)
	return b.AccountCharacteristics
}

func (b *BillingStatement2) AddRateData() *BillingRate1 {
	newValue := new(BillingRate1)
	b.RateData = append(b.RateData, newValue)
	return newValue
}

func (b *BillingStatement2) AddCurrencyExchange() *CurrencyExchange6 {
	newValue := new(CurrencyExchange6)
	b.CurrencyExchange = append(b.CurrencyExchange, newValue)
	return newValue
}

func (b *BillingStatement2) AddBalance() *BillingBalance1 {
	newValue := new(BillingBalance1)
	b.Balance = append(b.Balance, newValue)
	return newValue
}

func (b *BillingStatement2) AddCompensation() *BillingCompensation1 {
	newValue := new(BillingCompensation1)
	b.Compensation = append(b.Compensation, newValue)
	return newValue
}

func (b *BillingStatement2) AddService() *BillingService2 {
	newValue := new(BillingService2)
	b.Service = append(b.Service, newValue)
	return newValue
}

func (b *BillingStatement2) AddTaxRegion() *BillingTaxRegion1 {
	newValue := new(BillingTaxRegion1)
	b.TaxRegion = append(b.TaxRegion, newValue)
	return newValue
}

func (b *BillingStatement2) AddBalanceAdjustment() *BalanceAdjustment1 {
	newValue := new(BalanceAdjustment1)
	b.BalanceAdjustment = append(b.BalanceAdjustment, newValue)
	return newValue
}

func (b *BillingStatement2) AddServiceAdjustment() *BillingServiceAdjustment1 {
	newValue := new(BillingServiceAdjustment1)
	b.ServiceAdjustment = append(b.ServiceAdjustment, newValue)
	return newValue
}
