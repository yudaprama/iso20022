package model

// Result of an individual terminal management action by the point of interaction.
type TMSActionIdentification1 struct {

	// Types of terminal management action performed by a point of interaction.
	ActionType *TerminalManagementAction1Code `xml:"ActnTp"`

	// Data set on which the action has been performed.
	DataSetIdentification *DataSetIdentification2 `xml:"DataSetId,omitempty"`
}

func (t *TMSActionIdentification1) SetActionType(value string) {
	t.ActionType = (*TerminalManagementAction1Code)(&value)
}

func (t *TMSActionIdentification1) AddDataSetIdentification() *DataSetIdentification2 {
	t.DataSetIdentification = new(DataSetIdentification2)
	return t.DataSetIdentification
}
