package model

// Set of actions to be performed by the POI (Point Of Interaction) system.
type Action2 struct {

	// Type of action to be performed by the POI (Point Of Interaction) system.
	ActionType *ActionType2Code `xml:"ActnTp"`

	// Message to be displayed to the cardholder or the cashier.
	MessageToPresent *ActionMessage1 `xml:"MsgToPres,omitempty"`
}

func (a *Action2) SetActionType(value string) {
	a.ActionType = (*ActionType2Code)(&value)
}

func (a *Action2) AddMessageToPresent() *ActionMessage1 {
	a.MessageToPresent = new(ActionMessage1)
	return a.MessageToPresent
}
