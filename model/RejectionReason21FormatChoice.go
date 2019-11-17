package model

// Choice of formats to express the reason of a rejection of the interest request.
type RejectionReason21FormatChoice struct {

	// Provides the interest rejection reason using an ISO 20022 code.
	Code *InterestRejectionReason1Code `xml:"Cd"`

	// Provides the interest rejection reason using a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionReason21FormatChoice) SetCode(value string) {
	r.Code = (*InterestRejectionReason1Code)(&value)
}

func (r *RejectionReason21FormatChoice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
