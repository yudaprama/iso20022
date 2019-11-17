package model

// Set of elements providing the total sum of entries per bank transaction code.
type NumberAndSumOfTransactionsPerBankTransactionCode1 struct {

	// Number of individual entries contained in the report.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`

	// Resulting amount of the netted amounts for all debit and credit entries per bank transaction code.
	TotalNetEntryAmount *DecimalNumber `xml:"TtlNetNtryAmt,omitempty"`

	// Indicates whether the total net entry amount is a credit or a debit amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Set of elements to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure1 `xml:"BkTxCd"`

	// Set of elements used to indicate when the booked amount of money will become available, ie can be accessed and start generating interest.
	Availability []*CashBalanceAvailability1 `xml:"Avlbty,omitempty"`
}

func (n *NumberAndSumOfTransactionsPerBankTransactionCode1) SetNumberOfEntries(value string) {
	n.NumberOfEntries = (*Max15NumericText)(&value)
}

func (n *NumberAndSumOfTransactionsPerBankTransactionCode1) SetSum(value string) {
	n.Sum = (*DecimalNumber)(&value)
}

func (n *NumberAndSumOfTransactionsPerBankTransactionCode1) SetTotalNetEntryAmount(value string) {
	n.TotalNetEntryAmount = (*DecimalNumber)(&value)
}

func (n *NumberAndSumOfTransactionsPerBankTransactionCode1) SetCreditDebitIndicator(value string) {
	n.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (n *NumberAndSumOfTransactionsPerBankTransactionCode1) AddBankTransactionCode() *BankTransactionCodeStructure1 {
	n.BankTransactionCode = new(BankTransactionCodeStructure1)
	return n.BankTransactionCode
}

func (n *NumberAndSumOfTransactionsPerBankTransactionCode1) AddAvailability() *CashBalanceAvailability1 {
	newValue := new(CashBalanceAvailability1)
	n.Availability = append(n.Availability, newValue)
	return newValue
}
