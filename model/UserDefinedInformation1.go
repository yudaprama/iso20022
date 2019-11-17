package model

// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
type UserDefinedInformation1 struct {

	// Identifies the nature of the user information.
	Label *Max35Text `xml:"Labl"`

	// Specifies the content of the user information.
	Information *Max140Text `xml:"Inf"`
}

func (u *UserDefinedInformation1) SetLabel(value string) {
	u.Label = (*Max35Text)(&value)
}

func (u *UserDefinedInformation1) SetInformation(value string) {
	u.Information = (*Max140Text)(&value)
}
