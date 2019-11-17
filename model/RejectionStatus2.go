package model

// Provides a rejection reason and additional information.
type RejectionStatus2 struct {

	// Provides the rejection reason using a code.
	RejectedReason *RejectionReasonV021Code `xml:"RjctdRsn"`

	// Allows to provides additional information to the rejection reason code.
	AdditionalInformation *Max35Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectionStatus2) SetRejectedReason(value string) {
	r.RejectedReason = (*RejectionReasonV021Code)(&value)
}

func (r *RejectionStatus2) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max35Text)(&value)
}
