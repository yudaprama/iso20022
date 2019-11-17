package model

// Describes the activities that took place during a certain period for one trade transaction.
type ActivityDetails1 struct {

	// Date and time when the activity occurred.
	DateTime *ISODateTime `xml:"DtTm"`

	// Description of the reported activities.
	Activity *Activity1 `xml:"Actvty"`

	// Financial institution which initiated the activity.
	Initiator *BICIdentification1 `xml:"Initr"`
}

func (a *ActivityDetails1) SetDateTime(value string) {
	a.DateTime = (*ISODateTime)(&value)
}

func (a *ActivityDetails1) AddActivity() *Activity1 {
	a.Activity = new(Activity1)
	return a.Activity
}

func (a *ActivityDetails1) AddInitiator() *BICIdentification1 {
	a.Initiator = new(BICIdentification1)
	return a.Initiator
}
