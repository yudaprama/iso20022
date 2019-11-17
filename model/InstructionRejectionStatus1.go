package model

// Status advising on the rejection of the instruction or cancellation request.
type InstructionRejectionStatus1 struct {

	// Reason advising the rejection of the instruction.
	Reason *RejectionReason1Code `xml:"Rsn"`

	// This code can be used in case another reason is required.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Additional information about the reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (i *InstructionRejectionStatus1) SetReason(value string) {
	i.Reason = (*RejectionReason1Code)(&value)
}

func (i *InstructionRejectionStatus1) SetExtendedReason(value string) {
	i.ExtendedReason = (*Extended350Code)(&value)
}

func (i *InstructionRejectionStatus1) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max350Text)(&value)
}
