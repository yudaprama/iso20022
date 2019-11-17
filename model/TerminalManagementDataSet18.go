package model

// Data related to the management plan of a point of interaction (POI).
type TerminalManagementDataSet18 struct {

	// Identification of the data set containing the management plan.
	Identification *DataSetIdentification6 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the management plan.
	Content *ManagementPlanContent5 `xml:"Cntt,omitempty"`
}

func (t *TerminalManagementDataSet18) AddIdentification() *DataSetIdentification6 {
	t.Identification = new(DataSetIdentification6)
	return t.Identification
}

func (t *TerminalManagementDataSet18) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet18) AddContent() *ManagementPlanContent5 {
	t.Content = new(ManagementPlanContent5)
	return t.Content
}
