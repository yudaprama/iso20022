package model

// Rate name specifies the reference rate or basis rate on which a variable rate is based (ex: EONIA, EURIBOR, LIBOR, FEFUND, EURREPO).
type RateName1 struct {

	// Entity that assigns the identification.
	Issuer *Max8Text `xml:"Issr,omitempty"`

	// Rate Name specifies the reference rate or basis rate on which a variable rate is based (ex: EONIA, EURIBOR, LIBOR, FEFUND, EURREPO).
	RateName *Max35Text `xml:"RateNm"`
}

func (r *RateName1) SetIssuer(value string) {
	r.Issuer = (*Max8Text)(&value)
}

func (r *RateName1) SetRateName(value string) {
	r.RateName = (*Max35Text)(&value)
}
