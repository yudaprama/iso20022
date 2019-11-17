package model

// Response of a requested service.
type ResponseType5 struct {

	// Result of the transaction.
	Response *Response4Code `xml:"Rspn"`

	// Detailed result of the transaction.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`

	// Additional information on the response for further examination.
	AdditionalResponseInformation *Max140Text `xml:"AddtlRspnInf,omitempty"`
}

func (r *ResponseType5) SetResponse(value string) {
	r.Response = (*Response4Code)(&value)
}

func (r *ResponseType5) SetResponseReason(value string) {
	r.ResponseReason = (*Max35Text)(&value)
}

func (r *ResponseType5) SetAdditionalResponseInformation(value string) {
	r.AdditionalResponseInformation = (*Max140Text)(&value)
}
