package model

// Provides the reason for rejecting the case assignment.
type InvestigationRejectionJustification1 struct {

	// Reason for the rejection of a case assignment, in a coded form.
	RejectionReason *InvestigationRejection1Code `xml:"RjctnRsn"`
}

func (i *InvestigationRejectionJustification1) SetRejectionReason(value string) {
	i.RejectionReason = (*InvestigationRejection1Code)(&value)
}
