package model

// Result of performed verifications for the transaction.
type TransactionVerificationResult4 struct {

	// Method of verification that has been performed.
	Method *AuthenticationMethod6Code `xml:"Mtd"`

	// Entity or device that has performed the verification.
	VerificationEntity *AuthenticationEntity2Code `xml:"VrfctnNtty,omitempty"`

	// Result of the verification.
	Result *Verification1Code `xml:"Rslt,omitempty"`

	// Additional result of the verification.
	AdditionalResult *Max500Text `xml:"AddtlRslt,omitempty"`
}

func (t *TransactionVerificationResult4) SetMethod(value string) {
	t.Method = (*AuthenticationMethod6Code)(&value)
}

func (t *TransactionVerificationResult4) SetVerificationEntity(value string) {
	t.VerificationEntity = (*AuthenticationEntity2Code)(&value)
}

func (t *TransactionVerificationResult4) SetResult(value string) {
	t.Result = (*Verification1Code)(&value)
}

func (t *TransactionVerificationResult4) SetAdditionalResult(value string) {
	t.AdditionalResult = (*Max500Text)(&value)
}
