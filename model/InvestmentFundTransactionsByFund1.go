package model

// Investment fund transactions for a specific financial instrument.
type InvestmentFundTransactionsByFund1 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification1Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Specifies the form, ie, ownership, of the security.
	Form *FormOfSecurity1Code `xml:"Form,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Process of buying, selling, switching or transferring fund units.
	TransactionDetails []*InvestmentFundTransaction2 `xml:"TxDtls"`

	// Balance of the financial instrument for this specific statement page.
	BalanceByPage *PaginationBalance1 `xml:"BalByPg,omitempty"`
}

func (i *InvestmentFundTransactionsByFund1) AddIdentification() *SecurityIdentification1Choice {
	i.Identification = new(SecurityIdentification1Choice)
	return i.Identification
}

func (i *InvestmentFundTransactionsByFund1) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *InvestmentFundTransactionsByFund1) SetSupplementaryIdentification(value string) {
	i.SupplementaryIdentification = (*Max35Text)(&value)
}

func (i *InvestmentFundTransactionsByFund1) SetForm(value string) {
	i.Form = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentFundTransactionsByFund1) SetClassType(value string) {
	i.ClassType = (*Max35Text)(&value)
}

func (i *InvestmentFundTransactionsByFund1) SetDistributionPolicy(value string) {
	i.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (i *InvestmentFundTransactionsByFund1) AddTransactionDetails() *InvestmentFundTransaction2 {
	newValue := new(InvestmentFundTransaction2)
	i.TransactionDetails = append(i.TransactionDetails, newValue)
	return newValue
}

func (i *InvestmentFundTransactionsByFund1) AddBalanceByPage() *PaginationBalance1 {
	i.BalanceByPage = new(PaginationBalance1)
	return i.BalanceByPage
}
