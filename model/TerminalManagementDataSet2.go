package model

// Data related to the management plan of a point of interaction (POI).
type TerminalManagementDataSet2 struct {

	// Identification of the data set containing the management plan.
	Identification *DataSetIdentification2 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the management plan.
	Content *ManagementPlanContent1 `xml:"Cntt,omitempty"`
}

func (t *TerminalManagementDataSet2) AddIdentification() *DataSetIdentification2 {
	t.Identification = new(DataSetIdentification2)
	return t.Identification
}

func (t *TerminalManagementDataSet2) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet2) AddContent() *ManagementPlanContent1 {
	t.Content = new(ManagementPlanContent1)
	return t.Content
}
