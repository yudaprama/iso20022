package model

// Provides the money market statistical report instrument related header details.
type MoneyMarketReportHeader1 struct {

	// Agent which is subject to reporting requirements.
	ReportingAgent *LEIIdentifier `xml:"RptgAgt"`

	// Beginning and ending date-time to which the transaction data refers (trade date in case of new transactions and date of amendment in case of revisions).
	ReferencePeriod *DateTimePeriod1 `xml:"RefPrd"`
}

func (m *MoneyMarketReportHeader1) SetReportingAgent(value string) {
	m.ReportingAgent = (*LEIIdentifier)(&value)
}

func (m *MoneyMarketReportHeader1) AddReferencePeriod() *DateTimePeriod1 {
	m.ReferencePeriod = new(DateTimePeriod1)
	return m.ReferencePeriod
}
