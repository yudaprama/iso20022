package model

// Provides account identification information.
type AccountIdentification10 struct {

	// Standard code to specify that announcement applies to all safekeeping accounts that own underlying financial instrument.
	IdentificationCode *SafekeepingAccountIdentification1Code `xml:"IdCd"`
}

func (a *AccountIdentification10) SetIdentificationCode(value string) {
	a.IdentificationCode = (*SafekeepingAccountIdentification1Code)(&value)
}
