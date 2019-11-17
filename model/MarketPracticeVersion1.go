package model

// Identifies the implementation and version.
type MarketPracticeVersion1 struct {

	// Market practice, for example, “UKTRANSFERS”, “FINDELSLT”.
	Name *Max35Text `xml:"Nm"`

	// Year and month, for example, 2013-06
	Date *ISOYearMonth `xml:"Dt,omitempty"`

	// Version of the market practice.
	Number *Max35Text `xml:"Nb,omitempty"`
}

func (m *MarketPracticeVersion1) SetName(value string) {
	m.Name = (*Max35Text)(&value)
}

func (m *MarketPracticeVersion1) SetDate(value string) {
	m.Date = (*ISOYearMonth)(&value)
}

func (m *MarketPracticeVersion1) SetNumber(value string) {
	m.Number = (*Max35Text)(&value)
}
