package model

// Provides the related report identification and its status. If the status is rejected, a reason for this status must be given.
type ReportStatusAndReason1 struct {

	// Provides the identification of the RegulatoryTransactionReport document that was previously sent by the reporting institution.
	RelatedReference *Max35Text `xml:"RltdRef"`

	// Indicates the status of a report message.
	Status *Status2Code `xml:"Sts"`

	// Indicates that the report is rejected and provides a reason why.
	Rejected []*RejectedStatusReason9Choice `xml:"Rjctd"`
}

func (r *ReportStatusAndReason1) SetRelatedReference(value string) {
	r.RelatedReference = (*Max35Text)(&value)
}

func (r *ReportStatusAndReason1) SetStatus(value string) {
	r.Status = (*Status2Code)(&value)
}

func (r *ReportStatusAndReason1) AddRejected() *RejectedStatusReason9Choice {
	newValue := new(RejectedStatusReason9Choice)
	r.Rejected = append(r.Rejected, newValue)
	return newValue
}
