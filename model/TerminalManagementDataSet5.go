package model

// Data related to the management plan of a point of interaction (POI).
type TerminalManagementDataSet5 struct {

	// Identification of the data set containing the management plan.
	Identification *DataSetIdentification3 `xml:"Id"`

	// Counter to identify a single data set within the whole transfer.
	SequenceCounter *Max9NumericText `xml:"SeqCntr,omitempty"`

	// Content of the management plan.
	Content *ManagementPlanContent2 `xml:"Cntt,omitempty"`
}

func (t *TerminalManagementDataSet5) AddIdentification() *DataSetIdentification3 {
	t.Identification = new(DataSetIdentification3)
	return t.Identification
}

func (t *TerminalManagementDataSet5) SetSequenceCounter(value string) {
	t.SequenceCounter = (*Max9NumericText)(&value)
}

func (t *TerminalManagementDataSet5) AddContent() *ManagementPlanContent2 {
	t.Content = new(ManagementPlanContent2)
	return t.Content
}
