package model

// Details of the settlement condition modification request
type RequestDetails15 struct {

	// References of the transaction for which the securities settlement condition modification is requested.
	Reference *References18 `xml:"Ref"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing7Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether the instruction due to expire is confirmed for settlement
	RetainIndicator *YesNoIndicator `xml:"RtnInd,omitempty"`

	// Specifies the type of linkage requested.
	Linkage *LinkageType3Choice `xml:"Lkg,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Specifies another type of processing change request
	OtherProcessing []*GenericIdentification30 `xml:"OthrPrcg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Specifies the matching processing change requested.
	MatchingDenial *MatchingDenied3Choice `xml:"MtchgDnl,omitempty"`

	// Specifies that the transaction is requested to be unilaterally split.
	UnilateralSplit *UnilateralSplit3Choice `xml:"UnltrlSplt,omitempty"`

	// Information regarding the linkage requested.
	Linkages []*Linkages39 `xml:"Lnkgs,omitempty"`
}

func (r *RequestDetails15) AddReference() *References18 {
	r.Reference = new(References18)
	return r.Reference
}

func (r *RequestDetails15) AddAutomaticBorrowing() *AutomaticBorrowing7Choice {
	r.AutomaticBorrowing = new(AutomaticBorrowing7Choice)
	return r.AutomaticBorrowing
}

func (r *RequestDetails15) SetRetainIndicator(value string) {
	r.RetainIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails15) AddLinkage() *LinkageType3Choice {
	r.Linkage = new(LinkageType3Choice)
	return r.Linkage
}

func (r *RequestDetails15) AddPriority() *PriorityNumeric4Choice {
	r.Priority = new(PriorityNumeric4Choice)
	return r.Priority
}

func (r *RequestDetails15) AddOtherProcessing() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	r.OtherProcessing = append(r.OtherProcessing, newValue)
	return newValue
}

func (r *RequestDetails15) SetPartialSettlementIndicator(value string) {
	r.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (r *RequestDetails15) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	r.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return r.SecuritiesRTGS
}

func (r *RequestDetails15) AddHoldIndicator() *HoldIndicator6 {
	r.HoldIndicator = new(HoldIndicator6)
	return r.HoldIndicator
}

func (r *RequestDetails15) AddMatchingDenial() *MatchingDenied3Choice {
	r.MatchingDenial = new(MatchingDenied3Choice)
	return r.MatchingDenial
}

func (r *RequestDetails15) AddUnilateralSplit() *UnilateralSplit3Choice {
	r.UnilateralSplit = new(UnilateralSplit3Choice)
	return r.UnilateralSplit
}

func (r *RequestDetails15) AddLinkages() *Linkages39 {
	newValue := new(Linkages39)
	r.Linkages = append(r.Linkages, newValue)
	return newValue
}
