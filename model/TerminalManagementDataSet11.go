package model

// Data set containing the acceptor parameters of a point of interaction (POI).
type TerminalManagementDataSet11 struct {

	// Identification of the data set transferred.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the acceptor parameters.
	Content *AcceptorConfigurationContent3 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet11) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet11) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet11) AddContent() *AcceptorConfigurationContent3 {
	t.Content = new(AcceptorConfigurationContent3)
	return t.Content
}
