package model

// Assessment of securities credit and investment risk.
type Rating1 struct {

	// Information regarding the entity that assigns the rating.
	RatingScheme *Max35Text `xml:"RatgSchme"`

	// Date/time as from which the rating is valid.
	ValueDate *ISODateTime `xml:"ValDt"`

	// Specifies the rating, which has been assigned to a security by a rating agency.
	ValueIdentification *RatingValueIdentifier `xml:"ValId"`
}

func (r *Rating1) SetRatingScheme(value string) {
	r.RatingScheme = (*Max35Text)(&value)
}

func (r *Rating1) SetValueDate(value string) {
	r.ValueDate = (*ISODateTime)(&value)
}

func (r *Rating1) SetValueIdentification(value string) {
	r.ValueIdentification = (*RatingValueIdentifier)(&value)
}
