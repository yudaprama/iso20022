package model

// Additional references linked to the updated interest request such the original InterestRequest identification, and optionaly the InterestResponse identification.
type Reference20 struct {

	// Provides the reference to the interest payment request.
	InterestPaymentRequestIdentification *Max35Text `xml:"IntrstPmtReqId"`

	// Provides the reference to the interest payment response.
	InterestPaymentResponseIdentification *Max35Text `xml:"IntrstPmtRspnId,omitempty"`
}

func (r *Reference20) SetInterestPaymentRequestIdentification(value string) {
	r.InterestPaymentRequestIdentification = (*Max35Text)(&value)
}

func (r *Reference20) SetInterestPaymentResponseIdentification(value string) {
	r.InterestPaymentResponseIdentification = (*Max35Text)(&value)
}
