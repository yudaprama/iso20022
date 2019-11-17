package model

// Reason to reject the message.
type RejectionReason3 struct {

	// Reason to reject the message.
	Reason *MessageRejectedReason1Code `xml:"Rsn"`

	// Additional information about the rejection reason.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`

	// Linked previous reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessagePreviousReference *AdditionalReference3 `xml:"LkdMsgPrvsRef,omitempty"`

	// Linked other reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessageOtherReference *AdditionalReference3 `xml:"LkdMsgOthrRef,omitempty"`

	// Linked related reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessageRelatedReference *AdditionalReference3 `xml:"LkdMsgRltdRef,omitempty"`
}

func (r *RejectionReason3) SetReason(value string) {
	r.Reason = (*MessageRejectedReason1Code)(&value)
}

func (r *RejectionReason3) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max140Text)(&value)
}

func (r *RejectionReason3) AddLinkedMessagePreviousReference() *AdditionalReference3 {
	r.LinkedMessagePreviousReference = new(AdditionalReference3)
	return r.LinkedMessagePreviousReference
}

func (r *RejectionReason3) AddLinkedMessageOtherReference() *AdditionalReference3 {
	r.LinkedMessageOtherReference = new(AdditionalReference3)
	return r.LinkedMessageOtherReference
}

func (r *RejectionReason3) AddLinkedMessageRelatedReference() *AdditionalReference3 {
	r.LinkedMessageRelatedReference = new(AdditionalReference3)
	return r.LinkedMessageRelatedReference
}
