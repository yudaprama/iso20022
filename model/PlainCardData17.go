package model

// Sensitive data associated with a payment card.
type PlainCardData17 struct {

	// Primary Account Number (PAN) of the card.
	PAN *Min8Max28NumericText `xml:"PAN,omitempty"`

	// ISO track 1 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The format is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track1 *Max76Text `xml:"Trck1,omitempty"`

	// ISO track 2 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 7813, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track2 *Max37Text `xml:"Trck2,omitempty"`

	// ISO track 3 issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The content is conform to ISO 4909, removing beginning and ending sentinels and longitudinal redundancy check characters.
	Track3 *Max104Text `xml:"Trck3,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData []*Max35Text `xml:"AddtlCardData,omitempty"`

	// Entry mode of the card.
	EntryMode *CardDataReading5Code `xml:"NtryMd,omitempty"`
}

func (p *PlainCardData17) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData17) SetTrack1(value string) {
	p.Track1 = (*Max76Text)(&value)
}

func (p *PlainCardData17) SetTrack2(value string) {
	p.Track2 = (*Max37Text)(&value)
}

func (p *PlainCardData17) SetTrack3(value string) {
	p.Track3 = (*Max104Text)(&value)
}

func (p *PlainCardData17) AddAdditionalCardData(value string) {
	p.AdditionalCardData = append(p.AdditionalCardData, (*Max35Text)(&value))
}

func (p *PlainCardData17) SetEntryMode(value string) {
	p.EntryMode = (*CardDataReading5Code)(&value)
}
