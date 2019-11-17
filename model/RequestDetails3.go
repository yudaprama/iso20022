package model

// Details of data request.
type RequestDetails3 struct {

	// Type of data being requested, for example, a sub-member BIC.
	Type *Max35Text `xml:"Tp"`

	// Specific report data requested, for example, a  BIC.
	Key *Max35Text `xml:"Key,omitempty"`
}

func (r *RequestDetails3) SetType(value string) {
	r.Type = (*Max35Text)(&value)
}

func (r *RequestDetails3) SetKey(value string) {
	r.Key = (*Max35Text)(&value)
}
