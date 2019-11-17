package model

// Response of a requested service.
type ResponseType1 struct {

	// Result of the transaction.
	Response *Response1Code `xml:"Rspn"`

	// Detailed result of the transaction.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`
}

func (r *ResponseType1) SetResponse(value string) {
	r.Response = (*Response1Code)(&value)
}

func (r *ResponseType1) SetResponseReason(value string) {
	r.ResponseReason = (*Max35Text)(&value)
}
