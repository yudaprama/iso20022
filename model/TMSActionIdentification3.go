package model

// Result of an individual terminal management action by the point of interaction.
type TMSActionIdentification3 struct {

	// Types of terminal management action performed by a point of interaction.
	ActionType *TerminalManagementAction1Code `xml:"ActnTp"`

	// Data set on which the action has been performed.
	DataSetIdentification *DataSetIdentification4 `xml:"DataSetId,omitempty"`
}

func (t *TMSActionIdentification3) SetActionType(value string) {
	t.ActionType = (*TerminalManagementAction1Code)(&value)
}

func (t *TMSActionIdentification3) AddDataSetIdentification() *DataSetIdentification4 {
	t.DataSetIdentification = new(DataSetIdentification4)
	return t.DataSetIdentification
}
