package model

// Provides details about the reponse to the interest payment request.
type InterestResponse1 struct {

	// Provides the type of the response, either accepted or rejected.
	ResponseType *Status4Code `xml:"RspnTp"`

	// Provides a reason for rejection identified using a code or a proprietary format.
	RejectionReason *RejectionReason21FormatChoice `xml:"RjctnRsn,omitempty"`

	// Provides additional information on the rejection reason.
	RejectionReasonInformation *Max140Text `xml:"RjctnRsnInf,omitempty"`

	// Provides the reference to the interest payment request.
	InterestPaymentRequestIdentification *Max35Text `xml:"IntrstPmtReqId"`
}

func (i *InterestResponse1) SetResponseType(value string) {
	i.ResponseType = (*Status4Code)(&value)
}

func (i *InterestResponse1) AddRejectionReason() *RejectionReason21FormatChoice {
	i.RejectionReason = new(RejectionReason21FormatChoice)
	return i.RejectionReason
}

func (i *InterestResponse1) SetRejectionReasonInformation(value string) {
	i.RejectionReasonInformation = (*Max140Text)(&value)
}

func (i *InterestResponse1) SetInterestPaymentRequestIdentification(value string) {
	i.InterestPaymentRequestIdentification = (*Max35Text)(&value)
}
