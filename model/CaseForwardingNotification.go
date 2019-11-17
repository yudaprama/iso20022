package model

// Status of a case resulting from a case assignment.
type CaseForwardingNotification struct {

	// Justification for the forward action.
	Justification *CaseForwardingNotification1Code `xml:"Justfn"`
}

func (c *CaseForwardingNotification) SetJustification(value string) {
	c.Justification = (*CaseForwardingNotification1Code)(&value)
}
