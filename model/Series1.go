package model

// Identification of a series.
type Series1 struct {

	// Issue date of the fund series. It is typically applicable to a redemption order, subscription order confirmation or redemption order confirmation, but may be specified in the subscription order, if known.
	SeriesDate *DateFormat42Choice `xml:"SrsDt,omitempty"`

	// Name of the fund series. It is typically applicable to a redemption order, subscription order confirmation or redemption order confirmation, but may be specified in the subscription, if known.
	SeriesName *Max35Text `xml:"SrsNm,omitempty"`
}

func (s *Series1) AddSeriesDate() *DateFormat42Choice {
	s.SeriesDate = new(DateFormat42Choice)
	return s.SeriesDate
}

func (s *Series1) SetSeriesName(value string) {
	s.SeriesName = (*Max35Text)(&value)
}
