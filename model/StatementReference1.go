package model

// Identifies the reference of the statment  by a unique identifier and the date (and time).
type StatementReference1 struct {

	// Reference common to all pages of the statement for which the status advice is sent.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Date and time the statement was created.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Page number of the related message (within the statement) and continuation indicator to indicate that the statement is to continue or that the related message is the last page of the statement.
	Pagination *Pagination `xml:"Pgntn,omitempty"`
}

func (s *StatementReference1) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *StatementReference1) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *StatementReference1) AddPagination() *Pagination {
	s.Pagination = new(Pagination)
	return s.Pagination
}
