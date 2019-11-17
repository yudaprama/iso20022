package model

// Rate name specifies the reference rate or basis rate on which a variable rate is based (ex: EONIA, EURIBOR, LIBOR, FEFUND, EURREPO).
type RateName2 struct {

	// Entity that assigns the identification.
	Issuer *RestrictedFINXMax8Text `xml:"Issr,omitempty"`

	// Rate Name specifies the reference rate or basis rate on which a variable rate is based (ex: EONIA, EURIBOR, LIBOR, FEFUND, EURREPO).
	RateName *RestrictedFINXMax24Text `xml:"RateNm"`
}

func (r *RateName2) SetIssuer(value string) {
	r.Issuer = (*RestrictedFINXMax8Text)(&value)
}

func (r *RateName2) SetRateName(value string) {
	r.RateName = (*RestrictedFINXMax24Text)(&value)
}
