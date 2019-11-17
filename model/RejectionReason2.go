package model

// General information about the reason of the rejection.
type RejectionReason2 struct {

	// Reason of the rejection provided by the rejecting party.
	RejectingPartyReason *Max35Text `xml:"RjctgPtyRsn"`

	// Date and time at which the rejection was generated.
	RejectionDateTime *ISODateTime `xml:"RjctnDtTm,omitempty"`

	// Description of the precise location of the potential error in a message.
	ErrorLocation *Max350Text `xml:"ErrLctn,omitempty"`

	// Detailed description of the rejection reason.
	ReasonDescription *Max350Text `xml:"RsnDesc,omitempty"`

	// Additional information related to the rejection and meant to allow for the precise identification of the rejection reason. This could include a copy of the rejected message in part or in full.
	AdditionalData *Max20000Text `xml:"AddtlData,omitempty"`
}

func (r *RejectionReason2) SetRejectingPartyReason(value string) {
	r.RejectingPartyReason = (*Max35Text)(&value)
}

func (r *RejectionReason2) SetRejectionDateTime(value string) {
	r.RejectionDateTime = (*ISODateTime)(&value)
}

func (r *RejectionReason2) SetErrorLocation(value string) {
	r.ErrorLocation = (*Max350Text)(&value)
}

func (r *RejectionReason2) SetReasonDescription(value string) {
	r.ReasonDescription = (*Max350Text)(&value)
}

func (r *RejectionReason2) SetAdditionalData(value string) {
	r.AdditionalData = (*Max20000Text)(&value)
}
