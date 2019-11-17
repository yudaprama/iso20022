package model

// Choice between a location and codified form.
type PlaceOrUnderConfirmationChoice1 struct {

	// Party to which the presentation is to be made.
	PlaceOfPresentation *PlaceOfPresentation1 `xml:"PlcOfPresntn"`

	// Place of presentation when there is a confirmation.
	PresentationUnderConfirmation *PresentationParty1Code `xml:"PresntnUdrConf"`
}

func (p *PlaceOrUnderConfirmationChoice1) AddPlaceOfPresentation() *PlaceOfPresentation1 {
	p.PlaceOfPresentation = new(PlaceOfPresentation1)
	return p.PlaceOfPresentation
}

func (p *PlaceOrUnderConfirmationChoice1) SetPresentationUnderConfirmation(value string) {
	p.PresentationUnderConfirmation = (*PresentationParty1Code)(&value)
}
