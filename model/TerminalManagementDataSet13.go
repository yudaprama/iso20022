package model

// Data related to the status report of a point of interaction (POI).
type TerminalManagementDataSet13 struct {

	// Identification of the data set containing the status report.
	Identification *DataSetIdentification4 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the status report.
	Content *StatusReportContent4 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet13) AddIdentification() *DataSetIdentification4 {
	t.Identification = new(DataSetIdentification4)
	return t.Identification
}

func (t *TerminalManagementDataSet13) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet13) AddContent() *StatusReportContent4 {
	t.Content = new(StatusReportContent4)
	return t.Content
}
