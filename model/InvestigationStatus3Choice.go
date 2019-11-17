package model

// Specifies the status of an investigation case.
type InvestigationStatus3Choice struct {

	// Specifies the status of the investigation, in a coded form.
	Confirmation *InvestigationExecutionConfirmation3Code `xml:"Conf"`

	// Reason for the rejection of a modification request, in a coded form.
	RejectedModification []*ModificationRejection2Code `xml:"RjctdMod"`

	// Indicates a duplicated case.
	// Usage: When present, the case identified in the message must be closed. The case identified as duplicated (in this component) will be pursued.
	DuplicateOf *Case3 `xml:"DplctOf"`

	// Indicates whether the cancellation of the assignment is confirmed or rejected.
	// Usage: If yes, the cancellation of the assignment is confirmed.
	// If no, the cancellation of the assignment is rejected and the investigation process will continue.
	AssignmentCancellationConfirmation *YesNoIndicator `xml:"AssgnmtCxlConf"`
}

func (i *InvestigationStatus3Choice) SetConfirmation(value string) {
	i.Confirmation = (*InvestigationExecutionConfirmation3Code)(&value)
}

func (i *InvestigationStatus3Choice) AddRejectedModification(value string) {
	i.RejectedModification = append(i.RejectedModification, (*ModificationRejection2Code)(&value))
}

func (i *InvestigationStatus3Choice) AddDuplicateOf() *Case3 {
	i.DuplicateOf = new(Case3)
	return i.DuplicateOf
}

func (i *InvestigationStatus3Choice) SetAssignmentCancellationConfirmation(value string) {
	i.AssignmentCancellationConfirmation = (*YesNoIndicator)(&value)
}
