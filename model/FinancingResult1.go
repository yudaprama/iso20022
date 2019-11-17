package model

// The final result of a single invoice financing request.
type FinancingResult1 struct {

	// Specifies the status of the financing request (e.g. financed. not financed).
	FinancingRequestStatus *RequestStatus1Code `xml:"FincgReqSts"`

	// Indicates the reasons that have determined the result of the single request.
	StatusReason *StatusReason4Choice `xml:"StsRsn,omitempty"`

	// Further details on the status reason.
	AdditionalStatusReasonInformation []*Max105Text `xml:"AddtlStsRsnInf,omitempty"`

	// Indicates amount financed related to the request.
	FinancedAmount *FinancingRateOrAmountChoice `xml:"FincdAmt,omitempty"`
}

func (f *FinancingResult1) SetFinancingRequestStatus(value string) {
	f.FinancingRequestStatus = (*RequestStatus1Code)(&value)
}

func (f *FinancingResult1) AddStatusReason() *StatusReason4Choice {
	f.StatusReason = new(StatusReason4Choice)
	return f.StatusReason
}

func (f *FinancingResult1) AddAdditionalStatusReasonInformation(value string) {
	f.AdditionalStatusReasonInformation = append(f.AdditionalStatusReasonInformation, (*Max105Text)(&value))
}

func (f *FinancingResult1) AddFinancedAmount() *FinancingRateOrAmountChoice {
	f.FinancedAmount = new(FinancingRateOrAmountChoice)
	return f.FinancedAmount
}
