package model

// Provides the return indicators and the investigation result.
type ReturnIndicator1 struct {

	// Specifies the dates between which period the response results relate to.
	ResponsePeriod *DateOrDateTimePeriodChoice `xml:"RspnPrd,omitempty"`

	// Identifies the authority request type as a code.
	AuthorityRequestType *AuthorityRequestType1 `xml:"AuthrtyReqTp"`

	// Provides the investigation result.
	InvestigationResult *InvestigationResult1Choice `xml:"InvstgtnRslt"`

	// Additional information, in free text form, to complement the investigation result.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`
}

func (r *ReturnIndicator1) AddResponsePeriod() *DateOrDateTimePeriodChoice {
	r.ResponsePeriod = new(DateOrDateTimePeriodChoice)
	return r.ResponsePeriod
}

func (r *ReturnIndicator1) AddAuthorityRequestType() *AuthorityRequestType1 {
	r.AuthorityRequestType = new(AuthorityRequestType1)
	return r.AuthorityRequestType
}

func (r *ReturnIndicator1) AddInvestigationResult() *InvestigationResult1Choice {
	r.InvestigationResult = new(InvestigationResult1Choice)
	return r.InvestigationResult
}

func (r *ReturnIndicator1) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max500Text)(&value)
}
