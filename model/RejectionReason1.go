package model

// Reason to reject the message.
type RejectionReason1 struct {

	// Reason to reject the message.
	Reason *MessageRejectedReason1Code `xml:"Rsn"`

	// Additional information about the rejection reason.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`

	// Linked previous reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessagePreviousReference *AdditionalReference2 `xml:"LkdMsgPrvsRef,omitempty"`

	// Linked other reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessageOtherReference *AdditionalReference2 `xml:"LkdMsgOthrRef,omitempty"`

	// Linked related reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessageRelatedReference *AdditionalReference2 `xml:"LkdMsgRltdRef,omitempty"`
}

func (r *RejectionReason1) SetReason(value string) {
	r.Reason = (*MessageRejectedReason1Code)(&value)
}

func (r *RejectionReason1) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max140Text)(&value)
}

func (r *RejectionReason1) AddLinkedMessagePreviousReference() *AdditionalReference2 {
	r.LinkedMessagePreviousReference = new(AdditionalReference2)
	return r.LinkedMessagePreviousReference
}

func (r *RejectionReason1) AddLinkedMessageOtherReference() *AdditionalReference2 {
	r.LinkedMessageOtherReference = new(AdditionalReference2)
	return r.LinkedMessageOtherReference
}

func (r *RejectionReason1) AddLinkedMessageRelatedReference() *AdditionalReference2 {
	r.LinkedMessageRelatedReference = new(AdditionalReference2)
	return r.LinkedMessageRelatedReference
}
