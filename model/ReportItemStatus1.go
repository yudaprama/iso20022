package model

// Status information of the report item.
type ReportItemStatus1 struct {

	// Reason for the exception status.
	Exception *ReportItemRejectionReason1Choice `xml:"Xcptn"`

	// Additional  information about the reason for the status that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`

	// Details of the report item.
	ReportItem []*ReportItem1 `xml:"RptItm,omitempty"`
}

func (r *ReportItemStatus1) AddException() *ReportItemRejectionReason1Choice {
	r.Exception = new(ReportItemRejectionReason1Choice)
	return r.Exception
}

func (r *ReportItemStatus1) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

func (r *ReportItemStatus1) AddReportItem() *ReportItem1 {
	newValue := new(ReportItem1)
	r.ReportItem = append(r.ReportItem, newValue)
	return newValue
}
