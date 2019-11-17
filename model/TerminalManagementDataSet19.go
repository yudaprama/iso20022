package model

// Data set containing the acceptor parameters of a point of interaction (POI).
type TerminalManagementDataSet19 struct {

	// Identification of the data set transferred.
	Identification *DataSetIdentification6 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Identification of the point of interactions involved by the configuration data set.
	POIIdentification []*GenericIdentification71 `xml:"POIId,omitempty"`

	// Scope of the configuration contained in the data set.
	ConfigurationScope *PartyType15Code `xml:"CfgtnScp,omitempty"`

	// Content of the acceptor parameters.
	Content *AcceptorConfigurationContent5 `xml:"Cntt"`
}

func (t *TerminalManagementDataSet19) AddIdentification() *DataSetIdentification6 {
	t.Identification = new(DataSetIdentification6)
	return t.Identification
}

func (t *TerminalManagementDataSet19) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet19) AddPOIIdentification() *GenericIdentification71 {
	newValue := new(GenericIdentification71)
	t.POIIdentification = append(t.POIIdentification, newValue)
	return newValue
}

func (t *TerminalManagementDataSet19) SetConfigurationScope(value string) {
	t.ConfigurationScope = (*PartyType15Code)(&value)
}

func (t *TerminalManagementDataSet19) AddContent() *AcceptorConfigurationContent5 {
	t.Content = new(AcceptorConfigurationContent5)
	return t.Content
}
