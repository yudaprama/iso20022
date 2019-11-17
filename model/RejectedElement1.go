package model

// Provides information on the rejection reason of an individual element.
type RejectedElement1 struct {

	// Sequence number that allows to easily identify the element that is rejected.
	ElementSequenceNumber *Number `xml:"ElmtSeqNb"`

	// Reason for rejecting an individual element.
	IndividualRejectionReason *Max140Text `xml:"IndvRjctnRsn"`
}

func (r *RejectedElement1) SetElementSequenceNumber(value string) {
	r.ElementSequenceNumber = (*Number)(&value)
}

func (r *RejectedElement1) SetIndividualRejectionReason(value string) {
	r.IndividualRejectionReason = (*Max140Text)(&value)
}
