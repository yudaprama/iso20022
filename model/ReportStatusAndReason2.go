package model

// Provides the related report identification and its status. If the status is rejected, a reason for this status must be given.
type ReportStatusAndReason2 struct {

	// Provides the identification of the RegulatoryTransactionReportCancellationRequest document that was previously sent by the reporting institution.
	RelatedReference *Max35Text `xml:"RltdRef"`

	// Indicates the status of a report cancellation request message.
	Status *Status2Code `xml:"Sts"`

	// Indicates that the cancellation is rejected and provides a reason why.
	Rejected []*RejectedCancellationStatusReason1Choice `xml:"Rjctd"`
}

func (r *ReportStatusAndReason2) SetRelatedReference(value string) {
	r.RelatedReference = (*Max35Text)(&value)
}

func (r *ReportStatusAndReason2) SetStatus(value string) {
	r.Status = (*Status2Code)(&value)
}

func (r *ReportStatusAndReason2) AddRejected() *RejectedCancellationStatusReason1Choice {
	newValue := new(RejectedCancellationStatusReason1Choice)
	r.Rejected = append(r.Rejected, newValue)
	return newValue
}
