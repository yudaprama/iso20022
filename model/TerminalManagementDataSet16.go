package model

// Data related to the status report of a point of interaction (POI).
type TerminalManagementDataSet16 struct {

	// Identification of the data set containing the status report.
	Identification *DataSetIdentification6 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the status report.
	Content *StatusReportContent5 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet16) AddIdentification() *DataSetIdentification6 {
	t.Identification = new(DataSetIdentification6)
	return t.Identification
}

func (t *TerminalManagementDataSet16) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet16) AddContent() *StatusReportContent5 {
	t.Content = new(StatusReportContent5)
	return t.Content
}
