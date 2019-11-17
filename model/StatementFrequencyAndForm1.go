package model

// Specifies a frequency, format and delivery address for statements.
type StatementFrequencyAndForm1 struct {

	// Specifies the frequency for sending statements.
	Frequency *Frequency7Code `xml:"Frqcy"`

	// Specifies the communication method for statements.
	CommunicationMethod *CommunicationMethod2Choice `xml:"ComMtd"`

	// Specifies the delivery address for statements.
	DeliveryAddress *Max350Text `xml:"DlvryAdr"`

	// Specifies the format for statements.
	Format *CommunicationFormat1Choice `xml:"Frmt"`
}

func (s *StatementFrequencyAndForm1) SetFrequency(value string) {
	s.Frequency = (*Frequency7Code)(&value)
}

func (s *StatementFrequencyAndForm1) AddCommunicationMethod() *CommunicationMethod2Choice {
	s.CommunicationMethod = new(CommunicationMethod2Choice)
	return s.CommunicationMethod
}

func (s *StatementFrequencyAndForm1) SetDeliveryAddress(value string) {
	s.DeliveryAddress = (*Max350Text)(&value)
}

func (s *StatementFrequencyAndForm1) AddFormat() *CommunicationFormat1Choice {
	s.Format = new(CommunicationFormat1Choice)
	return s.Format
}
