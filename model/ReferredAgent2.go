package model

// Provides the placement agent identification for a hedge fund if the investor was referred by one.
type ReferredAgent2 struct {

	// Indicates if the investor was referred by a placement agent.
	Referred *Referred1Code `xml:"Rfrd"`

	// Placement agent that referred the investor.
	ReferredPlacementAgent *PartyIdentification70Choice `xml:"RfrdPlcmntAgt,omitempty"`
}

func (r *ReferredAgent2) SetReferred(value string) {
	r.Referred = (*Referred1Code)(&value)
}

func (r *ReferredAgent2) AddReferredPlacementAgent() *PartyIdentification70Choice {
	r.ReferredPlacementAgent = new(PartyIdentification70Choice)
	return r.ReferredPlacementAgent
}
