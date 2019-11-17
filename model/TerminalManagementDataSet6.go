package model

// Data set containing the acceptor parameters of a point of interaction (POI).
type TerminalManagementDataSet6 struct {

	// Identification of the data set transferred.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the acceptor parameters.
	Content *AcceptorConfigurationContent2 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet6) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet6) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet6) AddContent() *AcceptorConfigurationContent2 {
	t.Content = new(AcceptorConfigurationContent2)
	return t.Content
}
