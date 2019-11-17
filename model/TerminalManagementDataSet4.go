package model

// Data related to the status report of a point of interaction (POI).
type TerminalManagementDataSet4 struct {

	// Identification of the data set containing the status report.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the status report.
	Content *StatusReportContent2 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet4) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet4) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet4) AddContent() *StatusReportContent2 {
	t.Content = new(StatusReportContent2)
	return t.Content
}
