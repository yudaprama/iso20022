package model

// Details of the amendment.
type Amendment6 struct {

	// Contents of the related proposed Undertaking Amendment message.
	UndertakingAmendmentMessage *UndertakingAmendmentMessage1 `xml:"UdrtkgAmdmntMsg"`

	// Unique and unambiguous identifier assigned by the applicant to the undertaking.
	ApplicantReferenceNumber *Max35Text `xml:"ApplcntRefNb,omitempty"`

	// Additional information related to the notification.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (a *Amendment6) AddUndertakingAmendmentMessage() *UndertakingAmendmentMessage1 {
	a.UndertakingAmendmentMessage = new(UndertakingAmendmentMessage1)
	return a.UndertakingAmendmentMessage
}

func (a *Amendment6) SetApplicantReferenceNumber(value string) {
	a.ApplicantReferenceNumber = (*Max35Text)(&value)
}

func (a *Amendment6) AddAdditionalInformation(value string) {
	a.AdditionalInformation = append(a.AdditionalInformation, (*Max2000Text)(&value))
}
