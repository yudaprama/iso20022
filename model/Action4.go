package model

// Set of actions to be performed by the card acceptor.
type Action4 struct {

	// Type of action to be performed by the card acceptor.
	ActionType *ActionType5Code `xml:"ActnTp"`

	// Message to be displayed to the cardholder or the cashier.
	MessageToPresent *ActionMessage2 `xml:"MsgToPres,omitempty"`
}

func (a *Action4) SetActionType(value string) {
	a.ActionType = (*ActionType5Code)(&value)
}

func (a *Action4) AddMessageToPresent() *ActionMessage2 {
	a.MessageToPresent = new(ActionMessage2)
	return a.MessageToPresent
}
