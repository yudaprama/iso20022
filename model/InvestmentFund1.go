package model

// Details about the investment fund class.
type InvestmentFund1 struct {

	// Identification of the investment fund or investment fund class.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, for example, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Number of shares outstanding for the investment fund or investment fund share class.
	TotalUnitsOutstanding *DecimalNumber `xml:"TtlUnitsOutsdng,omitempty"`

	// Total transactional units (subscriptions and redemptions) which are applied to the investment fund or  investment fund share class for the report period.
	TransactionalUnits *DecimalNumber `xml:"TxnlUnits,omitempty"`

	// Total value of the investment fund or  investment fund share class units
	TotalValue *AmountAndDirection30 `xml:"TtlVal,omitempty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	Price []*PriceInformation10 `xml:"Pric,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InvestmentFund1) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return i.FinancialInstrumentIdentification
}

func (i *InvestmentFund1) SetClassType(value string) {
	i.ClassType = (*Max35Text)(&value)
}

func (i *InvestmentFund1) SetTotalUnitsOutstanding(value string) {
	i.TotalUnitsOutstanding = (*DecimalNumber)(&value)
}

func (i *InvestmentFund1) SetTransactionalUnits(value string) {
	i.TransactionalUnits = (*DecimalNumber)(&value)
}

func (i *InvestmentFund1) AddTotalValue() *AmountAndDirection30 {
	i.TotalValue = new(AmountAndDirection30)
	return i.TotalValue
}

func (i *InvestmentFund1) AddPrice() *PriceInformation10 {
	newValue := new(PriceInformation10)
	i.Price = append(i.Price, newValue)
	return newValue
}

func (i *InvestmentFund1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
