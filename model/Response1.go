package model

// Provides details on the margin call, that is a description and a response type.
type Response1 struct {

	// Provides details about the response type.
	ResponseTypeDetails []*ResponseType1Choice `xml:"RspnTpDtls"`

	// Provides additional details related to the margin call response.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (r *Response1) AddResponseTypeDetails() *ResponseType1Choice {
	newValue := new(ResponseType1Choice)
	r.ResponseTypeDetails = append(r.ResponseTypeDetails, newValue)
	return newValue
}

func (r *Response1) SetDescription(value string) {
	r.Description = (*Max140Text)(&value)
}
