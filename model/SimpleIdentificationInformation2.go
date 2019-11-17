package model

// Information related to a party identification or account identification.
type SimpleIdentificationInformation2 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max34Text `xml:"Id"`
}

func (s *SimpleIdentificationInformation2) SetIdentification(value string) {
	s.Identification = (*Max34Text)(&value)
}
