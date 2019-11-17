package model

// Choice between a reason for no activity and the foreign exchange swaps segment transaction details.
type ForeignExchangeSwap2Choice struct {

	// Provides the reason why no transactions are being reported for a money market reporting period.
	DataSetAction *ReportPeriodActivity1Code `xml:"DataSetActn"`

	// Provides the details of the foreign exchange transaction as reported by the reporting agent.
	Transaction []*ForeignExchangeSwapTransaction2 `xml:"Tx"`
}

func (f *ForeignExchangeSwap2Choice) SetDataSetAction(value string) {
	f.DataSetAction = (*ReportPeriodActivity1Code)(&value)
}

func (f *ForeignExchangeSwap2Choice) AddTransaction() *ForeignExchangeSwapTransaction2 {
	newValue := new(ForeignExchangeSwapTransaction2)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}
