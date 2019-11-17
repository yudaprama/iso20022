package model

// Choice between a date and period.
type DateAndPeriod1Choice struct {

	// Date of the statement.
	StatementDate *DateAndDateTimeChoice `xml:"StmtDt"`

	// Period for the statement.
	StatementPeriod *Period2Choice `xml:"StmtPrd"`
}

func (d *DateAndPeriod1Choice) AddStatementDate() *DateAndDateTimeChoice {
	d.StatementDate = new(DateAndDateTimeChoice)
	return d.StatementDate
}

func (d *DateAndPeriod1Choice) AddStatementPeriod() *Period2Choice {
	d.StatementPeriod = new(Period2Choice)
	return d.StatementPeriod
}
