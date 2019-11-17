package model

// Amendment details and reason related to the registered contract.
type RegisteredContractAmendment1 struct {

	// Date of the amendment of the registered contract.
	AmendmentDate *ISODate `xml:"AmdmntDt"`

	// Reference of the amendment document.
	Document *DocumentIdentification28 `xml:"Doc"`

	// Date from which the amendment is applicable.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Reason for the amendment.
	AmendmentReason *Max35Text `xml:"AmdmntRsn,omitempty"`

	// Further details on the amendment.
	AdditionalInformation *Max1025Text `xml:"AddtlInf,omitempty"`
}

func (r *RegisteredContractAmendment1) SetAmendmentDate(value string) {
	r.AmendmentDate = (*ISODate)(&value)
}

func (r *RegisteredContractAmendment1) AddDocument() *DocumentIdentification28 {
	r.Document = new(DocumentIdentification28)
	return r.Document
}

func (r *RegisteredContractAmendment1) SetStartDate(value string) {
	r.StartDate = (*ISODate)(&value)
}

func (r *RegisteredContractAmendment1) SetAmendmentReason(value string) {
	r.AmendmentReason = (*Max35Text)(&value)
}

func (r *RegisteredContractAmendment1) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max1025Text)(&value)
}
