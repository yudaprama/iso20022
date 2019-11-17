package model

// Details of the processing request.
type RequestDetails19 struct {

	// Type of data being requested, for example, a sub-member BIC.
	Type *Max35Text `xml:"Tp"`

	// Identificates the requestor.
	RequestorIdentification *PartyIdentification73Choice `xml:"RqstrId,omitempty"`

	// Additional information to support the processing request.
	AdditionalRequestInformation []*Max35Text `xml:"AddtlReqInf,omitempty"`
}

func (r *RequestDetails19) SetType(value string) {
	r.Type = (*Max35Text)(&value)
}

func (r *RequestDetails19) AddRequestorIdentification() *PartyIdentification73Choice {
	r.RequestorIdentification = new(PartyIdentification73Choice)
	return r.RequestorIdentification
}

func (r *RequestDetails19) AddAdditionalRequestInformation(value string) {
	r.AdditionalRequestInformation = append(r.AdditionalRequestInformation, (*Max35Text)(&value))
}
