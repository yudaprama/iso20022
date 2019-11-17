package model

// Magnetic track or equivalent payment card data.
type TrackData1 struct {

	// Track number of the card.
	TrackNumber *Exact1NumericText `xml:"TrckNb,omitempty"`

	// Card track content or equivalent.
	TrackValue *Max140Text `xml:"TrckVal"`
}

func (t *TrackData1) SetTrackNumber(value string) {
	t.TrackNumber = (*Exact1NumericText)(&value)
}

func (t *TrackData1) SetTrackValue(value string) {
	t.TrackValue = (*Max140Text)(&value)
}
