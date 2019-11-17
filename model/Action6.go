package model

// Set of actions to be performed by the POI (Point Of Interaction) system.
type Action6 struct {

	// Type of action to be performed by the POI (Point Of Interaction) system.
	ActionType *ActionType3Code `xml:"ActnTp"`

	// Message to be displayed to the cardholder or the cashier.
	MessageToPresent *ActionMessage2 `xml:"MsgToPres,omitempty"`
}

func (a *Action6) SetActionType(value string) {
	a.ActionType = (*ActionType3Code)(&value)
}

func (a *Action6) AddMessageToPresent() *ActionMessage2 {
	a.MessageToPresent = new(ActionMessage2)
	return a.MessageToPresent
}
