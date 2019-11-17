package model

// Provides the reason for rejecting a RequestToCancelPayment.
type RejectedCancellationJustification struct {

	// Justification for the rejection of the cancellation.
	ReasonCode *PaymentCancellationRejection1Code `xml:"RsnCd"`

	// Free text justification for rejecting a cancellation.
	Reason *Max140Text `xml:"Rsn,omitempty"`
}

func (r *RejectedCancellationJustification) SetReasonCode(value string) {
	r.ReasonCode = (*PaymentCancellationRejection1Code)(&value)
}

func (r *RejectedCancellationJustification) SetReason(value string) {
	r.Reason = (*Max140Text)(&value)
}
