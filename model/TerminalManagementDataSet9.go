package model

// Data related to the status report of a point of interaction (POI).
type TerminalManagementDataSet9 struct {

	// Identification of the data set containing the status report.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the status report.
	Content *StatusReportContent3 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet9) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet9) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet9) AddContent() *StatusReportContent3 {
	t.Content = new(StatusReportContent3)
	return t.Content
}
