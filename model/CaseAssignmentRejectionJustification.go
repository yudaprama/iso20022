package model

// Provides the reason for rejecting the case assignment.
type CaseAssignmentRejectionJustification struct {

	// Reason for the rejection of a case assignment, in a coded form.
	RejectionReason *CaseAssignmentRejection1Code `xml:"RjctnRsn"`
}

func (c *CaseAssignmentRejectionJustification) SetRejectionReason(value string) {
	c.RejectionReason = (*CaseAssignmentRejection1Code)(&value)
}
