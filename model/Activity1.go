package model

// Information about the message that is associated with the activity.
type Activity1 struct {

	// Name of the message associated with the activity.
	MessageName *Max70Text `xml:"MsgNm"`

	// Further information on a message associated with the activity.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (a *Activity1) SetMessageName(value string) {
	a.MessageName = (*Max70Text)(&value)
}

func (a *Activity1) SetDescription(value string) {
	a.Description = (*Max140Text)(&value)
}
