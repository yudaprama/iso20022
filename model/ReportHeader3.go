package model

// Provides header details on the report.
type ReportHeader3 struct {

	// Identification of a report billing statement.
	ReportIdentification *Max35Text `xml:"RptId"`

	// Provides details on the page number of the message.
	//
	// Usage: The pagination of the message is only allowed when agreed between the parties.
	MessagePagination *Pagination `xml:"MsgPgntn,omitempty"`
}

func (r *ReportHeader3) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *ReportHeader3) AddMessagePagination() *Pagination {
	r.MessagePagination = new(Pagination)
	return r.MessagePagination
}
