package model

// Data related to the management plan of a point of interaction (POI).
type TerminalManagementDataSet10 struct {

	// Identification of the data set containing the management plan.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the management plan.
	Content *ManagementPlanContent3 `xml:"Cntt,omitempty"`
}

func (t *TerminalManagementDataSet10) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet10) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet10) AddContent() *ManagementPlanContent3 {
	t.Content = new(ManagementPlanContent3)
	return t.Content
}
