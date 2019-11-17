package model

// Response of a requested service.
type ResponseType6 struct {

	// Response of the terminal manager.
	Response *Response2Code `xml:"Rspn"`

	// Detail of the response.
	ResponseDetail *ResultDetail3Code `xml:"RspnDtl,omitempty"`

	// Additional information on the response for further examination.
	AdditionalResponse *Max140Text `xml:"AddtlRspn,omitempty"`
}

func (r *ResponseType6) SetResponse(value string) {
	r.Response = (*Response2Code)(&value)
}

func (r *ResponseType6) SetResponseDetail(value string) {
	r.ResponseDetail = (*ResultDetail3Code)(&value)
}

func (r *ResponseType6) SetAdditionalResponse(value string) {
	r.AdditionalResponse = (*Max140Text)(&value)
}
