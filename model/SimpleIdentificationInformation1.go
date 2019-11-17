package model

// Information related to an identification, eg, party identification or account identification.
type SimpleIdentificationInformation1 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *RestrictedFINXMax35Text `xml:"Id"`
}

func (s *SimpleIdentificationInformation1) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax35Text)(&value)
}
