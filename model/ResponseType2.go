package model

// Response of a requested service.
type ResponseType2 struct {

	// Result of the request message or advice message.
	Result *Response3Code `xml:"Rslt"`

	// Detail of the result.
	ResultDetails *ResultDetail1Code `xml:"RsltDtls,omitempty"`

	// Additional information to be logged for further examination.
	AdditionalResultInformation *Max140Text `xml:"AddtlRsltInf,omitempty"`
}

func (r *ResponseType2) SetResult(value string) {
	r.Result = (*Response3Code)(&value)
}

func (r *ResponseType2) SetResultDetails(value string) {
	r.ResultDetails = (*ResultDetail1Code)(&value)
}

func (r *ResponseType2) SetAdditionalResultInformation(value string) {
	r.AdditionalResultInformation = (*Max140Text)(&value)
}
