package model

// Specifies the reason for an action.
type Reason2 struct {

	// Detailed description of the rejection.
	Description *Max140Text `xml:"Desc"`
}

func (r *Reason2) SetDescription(value string) {
	r.Description = (*Max140Text)(&value)
}
