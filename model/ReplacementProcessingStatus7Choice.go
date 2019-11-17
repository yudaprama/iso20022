package model

// Choice of status for the replacement processing.
type ReplacementProcessingStatus7Choice struct {

	// Replacement of the trade is accepted.
	Accepted *ProprietaryReason1 `xml:"Accptd"`

	// Replacement of the trade is completed.
	Completed *ProprietaryReason1 `xml:"Cmpltd"`

	// Replacement of the trade is denied.
	Denied *ProprietaryReason1 `xml:"Dnd"`

	// Replacement of the trade is In repair.
	InRepair *ProprietaryReason1 `xml:"InRpr"`

	// Replacement of the trade is PartialReplacementAccepted.
	PartialReplacementAccepted *ProprietaryReason1 `xml:"PrtlRplcmntAccptd"`

	// Replacement of the trade is pending.
	Pending *ProprietaryReason1 `xml:"Pdg"`

	// Replacement of the trade is ReceivedAtIntermediary.
	ReceivedAtIntermediary *ProprietaryReason1 `xml:"RcvdAtIntrmy"`

	// Replacement of the trade is ReceivedAtStockExchange.
	ReceivedAtStockExchange *ProprietaryReason1 `xml:"RcvdAtStockXchg"`

	// Replacement of the trade is rejected.
	Rejected *ProprietaryReason1 `xml:"Rjctd"`

	// Replacement of the trade is Modification Requested.
	ModificationRequested *ProprietaryReason1 `xml:"ModReqd"`

	// Provides a proprietary status and a proprietary reason of the processing status of the trade.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts,omitempty"`
}

func (r *ReplacementProcessingStatus7Choice) AddAccepted() *ProprietaryReason1 {
	r.Accepted = new(ProprietaryReason1)
	return r.Accepted
}

func (r *ReplacementProcessingStatus7Choice) AddCompleted() *ProprietaryReason1 {
	r.Completed = new(ProprietaryReason1)
	return r.Completed
}

func (r *ReplacementProcessingStatus7Choice) AddDenied() *ProprietaryReason1 {
	r.Denied = new(ProprietaryReason1)
	return r.Denied
}

func (r *ReplacementProcessingStatus7Choice) AddInRepair() *ProprietaryReason1 {
	r.InRepair = new(ProprietaryReason1)
	return r.InRepair
}

func (r *ReplacementProcessingStatus7Choice) AddPartialReplacementAccepted() *ProprietaryReason1 {
	r.PartialReplacementAccepted = new(ProprietaryReason1)
	return r.PartialReplacementAccepted
}

func (r *ReplacementProcessingStatus7Choice) AddPending() *ProprietaryReason1 {
	r.Pending = new(ProprietaryReason1)
	return r.Pending
}

func (r *ReplacementProcessingStatus7Choice) AddReceivedAtIntermediary() *ProprietaryReason1 {
	r.ReceivedAtIntermediary = new(ProprietaryReason1)
	return r.ReceivedAtIntermediary
}

func (r *ReplacementProcessingStatus7Choice) AddReceivedAtStockExchange() *ProprietaryReason1 {
	r.ReceivedAtStockExchange = new(ProprietaryReason1)
	return r.ReceivedAtStockExchange
}

func (r *ReplacementProcessingStatus7Choice) AddRejected() *ProprietaryReason1 {
	r.Rejected = new(ProprietaryReason1)
	return r.Rejected
}

func (r *ReplacementProcessingStatus7Choice) AddModificationRequested() *ProprietaryReason1 {
	r.ModificationRequested = new(ProprietaryReason1)
	return r.ModificationRequested
}

func (r *ReplacementProcessingStatus7Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	r.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return r.ProprietaryStatus
}
