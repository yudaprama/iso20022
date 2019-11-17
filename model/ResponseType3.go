package model

// Response of a requested service.
type ResponseType3 struct {

	// Result of the requested transaction.
	Response *Response4Code `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *ResultDetail2Code `xml:"RspnRsn,omitempty"`

	// Additional information to be logged for further examination.
	AdditionalResponseInformation *Max140Text `xml:"AddtlRspnInf,omitempty"`
}

func (r *ResponseType3) SetResponse(value string) {
	r.Response = (*Response4Code)(&value)
}

func (r *ResponseType3) SetResponseReason(value string) {
	r.ResponseReason = (*ResultDetail2Code)(&value)
}

func (r *ResponseType3) SetAdditionalResponseInformation(value string) {
	r.AdditionalResponseInformation = (*Max140Text)(&value)
}
