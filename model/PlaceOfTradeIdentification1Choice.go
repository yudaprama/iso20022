package model

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type PlaceOfTradeIdentification1Choice struct {

	// Country in which the transaction is executed.
	Country *CountryCode `xml:"Ctry"`

	// Exchange at which the transaction is executed.
	Exchange *MICIdentifier `xml:"Xchg"`

	// Party's location at which the transaction is executed.
	Party *AnyBICIdentifier `xml:"Pty"`

	// Place at which the Over-the-Counter (OTC) transaction is executed.
	OverTheCounter *Max35Text `xml:"OverTheCntr"`
}

func (p *PlaceOfTradeIdentification1Choice) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

func (p *PlaceOfTradeIdentification1Choice) SetExchange(value string) {
	p.Exchange = (*MICIdentifier)(&value)
}

func (p *PlaceOfTradeIdentification1Choice) SetParty(value string) {
	p.Party = (*AnyBICIdentifier)(&value)
}

func (p *PlaceOfTradeIdentification1Choice) SetOverTheCounter(value string) {
	p.OverTheCounter = (*Max35Text)(&value)
}
