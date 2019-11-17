package model

// Set of actions to be performed by the card acceptor.
type Action7 struct {

	// Type of action to be performed by the card acceptor.
	ActionType *ActionType6Code `xml:"ActnTp"`

	// Information to display, print or log.
	MessageToPresent *ActionMessage4 `xml:"MsgToPres,omitempty"`

	// Message to send before the completion of the transaction.
	RequestToPerform *MessageFunction11Code `xml:"ReqToPrfrm,omitempty"`
}

func (a *Action7) SetActionType(value string) {
	a.ActionType = (*ActionType6Code)(&value)
}

func (a *Action7) AddMessageToPresent() *ActionMessage4 {
	a.MessageToPresent = new(ActionMessage4)
	return a.MessageToPresent
}

func (a *Action7) SetRequestToPerform(value string) {
	a.RequestToPerform = (*MessageFunction11Code)(&value)
}
