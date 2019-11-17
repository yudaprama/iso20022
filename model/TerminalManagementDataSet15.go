package model

// Data related to the management plan of a point of interaction (POI).
type TerminalManagementDataSet15 struct {

	// Identification of the data set containing the management plan.
	Identification *DataSetIdentification4 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the management plan.
	Content *ManagementPlanContent4 `xml:"Cntt,omitempty"`
}

func (t *TerminalManagementDataSet15) AddIdentification() *DataSetIdentification4 {
	t.Identification = new(DataSetIdentification4)
	return t.Identification
}

func (t *TerminalManagementDataSet15) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet15) AddContent() *ManagementPlanContent4 {
	t.Content = new(ManagementPlanContent4)
	return t.Content
}
