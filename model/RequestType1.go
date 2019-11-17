package model

// Specifies the original transaction number.
type RequestType1 struct {

	// Idetifies the transaction number.
	Number *Max35Text `xml:"Nb"`

	// Identifies the type of information request related to an original transaction number as a code.
	Type []*TransactionRequestType1Code `xml:"Tp"`

	// Additional information, in free text form, to complement the requested information.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`
}

func (r *RequestType1) SetNumber(value string) {
	r.Number = (*Max35Text)(&value)
}

func (r *RequestType1) AddType(value string) {
	r.Type = append(r.Type, (*TransactionRequestType1Code)(&value))
}

func (r *RequestType1) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max500Text)(&value)
}
