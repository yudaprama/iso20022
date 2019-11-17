package model

// Response of a requested service.
type ResponseType7 struct {

	// Result of the requested transaction.
	Response *Response4Code `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *ResultDetail4Code `xml:"RspnRsn,omitempty"`

	// Additional information to be logged for further examination.
	AdditionalResponseInformation *Max140Text `xml:"AddtlRspnInf,omitempty"`
}

func (r *ResponseType7) SetResponse(value string) {
	r.Response = (*Response4Code)(&value)
}

func (r *ResponseType7) SetResponseReason(value string) {
	r.ResponseReason = (*ResultDetail4Code)(&value)
}

func (r *ResponseType7) SetAdditionalResponseInformation(value string) {
	r.AdditionalResponseInformation = (*Max140Text)(&value)
}
