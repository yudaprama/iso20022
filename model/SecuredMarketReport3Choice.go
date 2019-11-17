package model

// Choice between a reason for no activity and the secured market segment transaction details.
type SecuredMarketReport3Choice struct {

	// Provides the reason why no transactions are being reported for a money market reporting period.
	DataSetAction *ReportPeriodActivity1Code `xml:"DataSetActn"`

	// Provides the details of the secured market transaction as reported by the reporting agent
	Transaction []*SecuredMarketTransaction3 `xml:"Tx"`
}

func (s *SecuredMarketReport3Choice) SetDataSetAction(value string) {
	s.DataSetAction = (*ReportPeriodActivity1Code)(&value)
}

func (s *SecuredMarketReport3Choice) AddTransaction() *SecuredMarketTransaction3 {
	newValue := new(SecuredMarketTransaction3)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
