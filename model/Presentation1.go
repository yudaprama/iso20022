package model

// Information for the presentation of documents.
type Presentation1 struct {

	// Medium through which the presentation can be submitted such as paper, electronic or both.
	Medium *PresentationMedium1Choice `xml:"Mdm,omitempty"`

	// Choice of representation for the place of presentation.
	PlaceOfPresentationOrUnderConfirmationChoice *PlaceOrUnderConfirmationChoice1 `xml:"PlcOfPresntnOrUdrConfChc,omitempty"`

	// Document required to be presented.
	Document []*Document8 `xml:"Doc,omitempty"`

	// Additional information related to the presentation.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (p *Presentation1) AddMedium() *PresentationMedium1Choice {
	p.Medium = new(PresentationMedium1Choice)
	return p.Medium
}

func (p *Presentation1) AddPlaceOfPresentationOrUnderConfirmationChoice() *PlaceOrUnderConfirmationChoice1 {
	p.PlaceOfPresentationOrUnderConfirmationChoice = new(PlaceOrUnderConfirmationChoice1)
	return p.PlaceOfPresentationOrUnderConfirmationChoice
}

func (p *Presentation1) AddDocument() *Document8 {
	newValue := new(Document8)
	p.Document = append(p.Document, newValue)
	return newValue
}

func (p *Presentation1) AddAdditionalInformation(value string) {
	p.AdditionalInformation = append(p.AdditionalInformation, (*Max2000Text)(&value))
}
