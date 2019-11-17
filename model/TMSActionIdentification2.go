package model

// Result of an individual terminal management action by the point of interaction.
type TMSActionIdentification2 struct {

	// Types of terminal management action performed by a point of interaction.
	ActionType *TerminalManagementAction1Code `xml:"ActnTp"`

	// Data set on which the action has been performed.
	DataSetIdentification *DataSetIdentification3 `xml:"DataSetId,omitempty"`
}

func (t *TMSActionIdentification2) SetActionType(value string) {
	t.ActionType = (*TerminalManagementAction1Code)(&value)
}

func (t *TMSActionIdentification2) AddDataSetIdentification() *DataSetIdentification3 {
	t.DataSetIdentification = new(DataSetIdentification3)
	return t.DataSetIdentification
}
