package model

// Electronic presentation information.
type Presentation3 struct {

	// Format for presentation documents that are submitted electronically.
	Format *DocumentFormat1Choice `xml:"Frmt,omitempty"`

	// Channel through which presentation documents are submitted electronically, such as SWIFT,  Web upload, or secure email.
	Channel *Channel1Choice `xml:"Chanl,omitempty"`

	// Uniform Resource Identifier (URI), such as a web or an email address, specifying where the presentation can be addressed.
	Address *Max256Text `xml:"Adr,omitempty"`
}

func (p *Presentation3) AddFormat() *DocumentFormat1Choice {
	p.Format = new(DocumentFormat1Choice)
	return p.Format
}

func (p *Presentation3) AddChannel() *Channel1Choice {
	p.Channel = new(Channel1Choice)
	return p.Channel
}

func (p *Presentation3) SetAddress(value string) {
	p.Address = (*Max256Text)(&value)
}
