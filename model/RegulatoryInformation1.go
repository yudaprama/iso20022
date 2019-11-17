package model

// Regulatory information.
type RegulatoryInformation1 struct {

	// Sector of economic activity, for example, SAE in the Italian market.
	Sector *Max35Text `xml:"Sctr,omitempty"`

	// Branch of economic activity, for example, RAE in the Italian market.
	Branch *Max35Text `xml:"Brnch,omitempty"`

	// Group of economic activity, for example, a code issued by a regulator.
	Group *Max35Text `xml:"Grp,omitempty"`

	// Other regulatory information.
	Other *Max35Text `xml:"Othr,omitempty"`
}

func (r *RegulatoryInformation1) SetSector(value string) {
	r.Sector = (*Max35Text)(&value)
}

func (r *RegulatoryInformation1) SetBranch(value string) {
	r.Branch = (*Max35Text)(&value)
}

func (r *RegulatoryInformation1) SetGroup(value string) {
	r.Group = (*Max35Text)(&value)
}

func (r *RegulatoryInformation1) SetOther(value string) {
	r.Other = (*Max35Text)(&value)
}
