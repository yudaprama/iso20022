package model

// Action to perform in case of error on the related action in progress.
type ErrorAction2 struct {

	// List of error action result codes.
	ActionResult []*TerminalManagementActionResult1Code `xml:"ActnRslt"`

	// Action to be processed for the related errors.
	ActionToProcess *TerminalManagementErrorAction2Code `xml:"ActnToPrc"`
}

func (e *ErrorAction2) AddActionResult(value string) {
	e.ActionResult = append(e.ActionResult, (*TerminalManagementActionResult1Code)(&value))
}

func (e *ErrorAction2) SetActionToProcess(value string) {
	e.ActionToProcess = (*TerminalManagementErrorAction2Code)(&value)
}
