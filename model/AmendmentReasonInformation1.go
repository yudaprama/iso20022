package model

// Set of elements used to provide information on the reason of the amendment of the mandate.
type AmendmentReasonInformation1 struct {

	// Party that issues the amendment request.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the amendment request.
	Reason *MandateReason1Choice `xml:"Rsn"`

	// Further details on the amendment request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (a *AmendmentReasonInformation1) AddOriginator() *PartyIdentification32 {
	a.Originator = new(PartyIdentification32)
	return a.Originator
}

func (a *AmendmentReasonInformation1) AddReason() *MandateReason1Choice {
	a.Reason = new(MandateReason1Choice)
	return a.Reason
}

func (a *AmendmentReasonInformation1) AddAdditionalInformation(value string) {
	a.AdditionalInformation = append(a.AdditionalInformation, (*Max105Text)(&value))
}
