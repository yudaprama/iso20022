package model

// Provides the details on the user identification or any user key that allows to check if the initiating party is allowed to issue the transaction.
type Authorisation1Choice struct {

	// Specifies the authorisation, in a coded form.
	Code *Authorisation1Code `xml:"Cd"`

	// Specifies the authorisation, in a free text form.
	Proprietary *Max128Text `xml:"Prtry"`
}

func (a *Authorisation1Choice) SetCode(value string) {
	a.Code = (*Authorisation1Code)(&value)
}

func (a *Authorisation1Choice) SetProprietary(value string) {
	a.Proprietary = (*Max128Text)(&value)
}
