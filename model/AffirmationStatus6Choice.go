package model

// Choice of status for the affirmation.
type AffirmationStatus6Choice struct {

	// Status of affirmation of a trade.
	Affirmed *ProprietaryReason1 `xml:"Affrmd"`

	// Trade has been unaffirmed.
	Unaffirmed *AffirmationReason1Choice `xml:"Uaffrmd"`

	// Provides a proprietary status and a proprietary reason of the affirmation of the trade.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts,omitempty"`
}

func (a *AffirmationStatus6Choice) AddAffirmed() *ProprietaryReason1 {
	a.Affirmed = new(ProprietaryReason1)
	return a.Affirmed
}

func (a *AffirmationStatus6Choice) AddUnaffirmed() *AffirmationReason1Choice {
	a.Unaffirmed = new(AffirmationReason1Choice)
	return a.Unaffirmed
}

func (a *AffirmationStatus6Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	a.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return a.ProprietaryStatus
}
