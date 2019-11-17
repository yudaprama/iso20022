package model

// Provides the placement agent identification for a hedge fund if the investor was referred by one.
type ReferredAgent1 struct {

	// Indicates if the investor was referred by a placement agent.
	Referred *Referred1Code `xml:"Rfrd"`

	// Placement agent that referred the investor.
	ReferredPlacementAgent *PartyIdentification2Choice `xml:"RfrdPlcmntAgt,omitempty"`
}

func (r *ReferredAgent1) SetReferred(value string) {
	r.Referred = (*Referred1Code)(&value)
}

func (r *ReferredAgent1) AddReferredPlacementAgent() *PartyIdentification2Choice {
	r.ReferredPlacementAgent = new(PartyIdentification2Choice)
	return r.ReferredPlacementAgent
}
