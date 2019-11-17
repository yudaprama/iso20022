package model

// Configuration of the PIN online verification.
type ATMSecurityConfiguration5 struct {

	// PIN block format the security module is able to manage for online verification of the PIN.
	PINFormat []*PINFormat4Code `xml:"PINFrmt,omitempty"`

	// Maximum number of digits the security module is able to accept when the cardholder enters its PIN.
	PINLengthCapabilities *Number `xml:"PINLngthCpblties,omitempty"`
}

func (a *ATMSecurityConfiguration5) AddPINFormat(value string) {
	a.PINFormat = append(a.PINFormat, (*PINFormat4Code)(&value))
}

func (a *ATMSecurityConfiguration5) SetPINLengthCapabilities(value string) {
	a.PINLengthCapabilities = (*Number)(&value)
}
