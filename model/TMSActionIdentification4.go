package model

// Result of an individual terminal management action by the point of interaction.
type TMSActionIdentification4 struct {

	// Types of terminal management action performed by a point of interaction.
	ActionType *TerminalManagementAction2Code `xml:"ActnTp"`

	// Data set on which the action has been performed.
	DataSetIdentification *DataSetIdentification6 `xml:"DataSetId,omitempty"`
}

func (t *TMSActionIdentification4) SetActionType(value string) {
	t.ActionType = (*TerminalManagementAction2Code)(&value)
}

func (t *TMSActionIdentification4) AddDataSetIdentification() *DataSetIdentification6 {
	t.DataSetIdentification = new(DataSetIdentification6)
	return t.DataSetIdentification
}
