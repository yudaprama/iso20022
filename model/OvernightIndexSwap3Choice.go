package model

// Choice between a reason for no activity and the overnight index swaps segment transaction details.
type OvernightIndexSwap3Choice struct {

	// Provides the reason why no transactions are being reported for a money market reporting period.
	DataSetAction *ReportPeriodActivity1Code `xml:"DataSetActn"`

	// Provides the details of the secured market transaction as reported by the reporting agent
	Transaction []*OvernightIndexSwapTransaction3 `xml:"Tx"`
}

func (o *OvernightIndexSwap3Choice) SetDataSetAction(value string) {
	o.DataSetAction = (*ReportPeriodActivity1Code)(&value)
}

func (o *OvernightIndexSwap3Choice) AddTransaction() *OvernightIndexSwapTransaction3 {
	newValue := new(OvernightIndexSwapTransaction3)
	o.Transaction = append(o.Transaction, newValue)
	return newValue
}
