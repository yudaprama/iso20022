package model

// Details of the settlement condition modification request
type RequestDetails1 struct {

	// References of the transaction for which the securities settlement condition modification is requested.
	Reference *References1 `xml:"Ref"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing2Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether the instruction due to expire is confirmed for settlement
	RetainIndicator *YesNoIndicator `xml:"RtnInd,omitempty"`

	// Specifies the type of linkage requested.
	Linkage *LinkageType1Choice `xml:"Lkg,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Specifies another type of processing change request
	OtherProcessing []*GenericIdentification20 `xml:"OthrPrcg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies the matching processing change requested.
	MatchingDenial *MatchingDenied1Choice `xml:"MtchgDnl,omitempty"`

	// Specifies that the transaction is requested to be unilaterally split.
	UnilateralSplit *UnilateralSplit1Choice `xml:"UnltrlSplt,omitempty"`

	// Information regarding the linkage requested.
	Linkages []*Linkages3 `xml:"Lnkgs,omitempty"`
}

func (r *RequestDetails1) AddReference() *References1 {
	r.Reference = new(References1)
	return r.Reference
}

func (r *RequestDetails1) AddAutomaticBorrowing() *AutomaticBorrowing2Choice {
	r.AutomaticBorrowing = new(AutomaticBorrowing2Choice)
	return r.AutomaticBorrowing
}

func (r *RequestDetails1) SetRetainIndicator(value string) {
	r.RetainIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails1) AddLinkage() *LinkageType1Choice {
	r.Linkage = new(LinkageType1Choice)
	return r.Linkage
}

func (r *RequestDetails1) AddPriority() *PriorityNumeric1Choice {
	r.Priority = new(PriorityNumeric1Choice)
	return r.Priority
}

func (r *RequestDetails1) AddOtherProcessing() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	r.OtherProcessing = append(r.OtherProcessing, newValue)
	return newValue
}

func (r *RequestDetails1) SetPartialSettlementIndicator(value string) {
	r.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails1) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	r.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return r.SecuritiesRTGS
}

func (r *RequestDetails1) SetHoldIndicator(value string) {
	r.HoldIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails1) AddMatchingDenial() *MatchingDenied1Choice {
	r.MatchingDenial = new(MatchingDenied1Choice)
	return r.MatchingDenial
}

func (r *RequestDetails1) AddUnilateralSplit() *UnilateralSplit1Choice {
	r.UnilateralSplit = new(UnilateralSplit1Choice)
	return r.UnilateralSplit
}

func (r *RequestDetails1) AddLinkages() *Linkages3 {
	newValue := new(Linkages3)
	r.Linkages = append(r.Linkages, newValue)
	return newValue
}
