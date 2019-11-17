package model

// Data related to the status report of a point of interaction (POI).
type TerminalManagementDataSet1 struct {

	// Identification of the data set containing the status report.
	Identification *DataSetIdentification2 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the status report.
	Content *StatusReportContent1 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet1) AddIdentification() *DataSetIdentification2 {
	t.Identification = new(DataSetIdentification2)
	return t.Identification
}

func (t *TerminalManagementDataSet1) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet1) AddContent() *StatusReportContent1 {
	t.Content = new(StatusReportContent1)
	return t.Content
}
