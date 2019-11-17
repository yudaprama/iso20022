package model

// Data set containing the acceptor parameters of a point of interaction (POI).
type TerminalManagementDataSet14 struct {

	// Identification of the data set transferred.
	Identification *DataSetIdentification4 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the acceptor parameters.
	Content *AcceptorConfigurationContent4 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet14) AddIdentification() *DataSetIdentification4 {
	t.Identification = new(DataSetIdentification4)
	return t.Identification
}

func (t *TerminalManagementDataSet14) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet14) AddContent() *AcceptorConfigurationContent4 {
	t.Content = new(AcceptorConfigurationContent4)
	return t.Content
}
