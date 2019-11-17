package model

// Choice between a reason for no activity and the unsecured market segment transaction details.
type UnsecuredMarketReport3Choice struct {

	// Provides the reason why no transactions are being reported for a money market reporting period.
	DataSetAction *ReportPeriodActivity1Code `xml:"DataSetActn"`

	// Provides the details of the unsecured market transaction as reported by the reporting agent.
	Transaction []*UnsecuredMarketTransaction3 `xml:"Tx"`
}

func (u *UnsecuredMarketReport3Choice) SetDataSetAction(value string) {
	u.DataSetAction = (*ReportPeriodActivity1Code)(&value)
}

func (u *UnsecuredMarketReport3Choice) AddTransaction() *UnsecuredMarketTransaction3 {
	newValue := new(UnsecuredMarketTransaction3)
	u.Transaction = append(u.Transaction, newValue)
	return newValue
}
