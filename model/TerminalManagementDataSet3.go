package model

// Data set containing the acceptor parameters of a point of interaction (POI).
type TerminalManagementDataSet3 struct {

	// Identification of the data set transferred.
	Identification *DataSetIdentification2 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the acceptor parameters.
	Content *AcceptorConfigurationContent1 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet3) AddIdentification() *DataSetIdentification2 {
	t.Identification = new(DataSetIdentification2)
	return t.Identification
}

func (t *TerminalManagementDataSet3) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet3) AddContent() *AcceptorConfigurationContent1 {
	t.Content = new(AcceptorConfigurationContent1)
	return t.Content
}
