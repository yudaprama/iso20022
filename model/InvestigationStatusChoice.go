package model

// Choice between different statuses of an investigation case.
type InvestigationStatusChoice struct {

	// Indicates the status of an investigation.
	Confirmation *InvestigationExecutionConfirmation1Code `xml:"Conf"`

	// Reason for the rejection of a modification request, in a coded form.
	RejectedModification []*PaymentModificationRejection1Code `xml:"RjctdMod"`

	// Explains the reason for rejecting a payment cancellation request.
	RejectedCancellation *RejectedCancellationJustification `xml:"RjctdCxl"`

	// Identifies a duplicated case. When present, the case identified in the message must be closed. The case identified as duplicated (in this component) will be pursued.
	DuplicateOf *Case `xml:"DplctOf"`

	// If yes, it means the cancellation of the assignment is confirmed.
	// If no, it means the cancellation of the assignment is rejected and the investigation process will continue.
	AssignmentCancellationConfirmation *YesNoIndicator `xml:"AssgnmtCxlConf"`
}

func (i *InvestigationStatusChoice) SetConfirmation(value string) {
	i.Confirmation = (*InvestigationExecutionConfirmation1Code)(&value)
}

func (i *InvestigationStatusChoice) AddRejectedModification(value string) {
	i.RejectedModification = append(i.RejectedModification, (*PaymentModificationRejection1Code)(&value))
}

func (i *InvestigationStatusChoice) AddRejectedCancellation() *RejectedCancellationJustification {
	i.RejectedCancellation = new(RejectedCancellationJustification)
	return i.RejectedCancellation
}

func (i *InvestigationStatusChoice) AddDuplicateOf() *Case {
	i.DuplicateOf = new(Case)
	return i.DuplicateOf
}

func (i *InvestigationStatusChoice) SetAssignmentCancellationConfirmation(value string) {
	i.AssignmentCancellationConfirmation = (*YesNoIndicator)(&value)
}
