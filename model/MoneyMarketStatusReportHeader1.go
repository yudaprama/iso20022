package model

// Provides the money market statistical status report header details.
type MoneyMarketStatusReportHeader1 struct {

	// Agent which is subject to reporting requirements.
	ReportingAgent *LEIIdentifier `xml:"RptgAgt"`

	// For daily reporting this is the day to which the transaction data in the status message refers (trade date or amendment date if there are corrections).
	// For periodic reporting this is the first and the last day to which the transaction data in the status message refers (trade date or amendment date in case of corrections).
	ReportingPeriod *DateTimePeriod1 `xml:"RptgPrd"`

	// Provides the status for the full report.
	ReportStatus *StatisticalReportingStatus1Code `xml:"RptSts"`

	// Provides the details of the rule which could not be validated.
	ValidationRule []*GenericValidationRuleIdentification1 `xml:"VldtnRule,omitempty"`
}

func (m *MoneyMarketStatusReportHeader1) SetReportingAgent(value string) {
	m.ReportingAgent = (*LEIIdentifier)(&value)
}

func (m *MoneyMarketStatusReportHeader1) AddReportingPeriod() *DateTimePeriod1 {
	m.ReportingPeriod = new(DateTimePeriod1)
	return m.ReportingPeriod
}

func (m *MoneyMarketStatusReportHeader1) SetReportStatus(value string) {
	m.ReportStatus = (*StatisticalReportingStatus1Code)(&value)
}

func (m *MoneyMarketStatusReportHeader1) AddValidationRule() *GenericValidationRuleIdentification1 {
	newValue := new(GenericValidationRuleIdentification1)
	m.ValidationRule = append(m.ValidationRule, newValue)
	return newValue
}
