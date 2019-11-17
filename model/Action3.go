package model

// Set of actions to be performed by the POI (Point Of Interaction) system.
type Action3 struct {

	// Type of action to be performed by the POI (Point Of Interaction) system.
	ActionType *ActionType3Code `xml:"ActnTp"`

	// Message to be displayed to the cardholder or the cashier.
	MessageToPresent *ActionMessage1 `xml:"MsgToPres,omitempty"`
}

func (a *Action3) SetActionType(value string) {
	a.ActionType = (*ActionType3Code)(&value)
}

func (a *Action3) AddMessageToPresent() *ActionMessage1 {
	a.MessageToPresent = new(ActionMessage1)
	return a.MessageToPresent
}
