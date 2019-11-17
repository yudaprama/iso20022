package model

// Result of performed verifications for the transaction.
type TransactionVerificationResult5 struct {

	// Method of verification that has been performed.
	Method *AuthenticationMethod7Code `xml:"Mtd"`

	// Entity or device that has performed the verification.
	VerificationEntity *AuthenticationEntity2Code `xml:"VrfctnNtty,omitempty"`

	// Result of the verification.
	Result *Verification1Code `xml:"Rslt,omitempty"`

	// Additional result of the verification.
	AdditionalResult *Max500Text `xml:"AddtlRslt,omitempty"`

	// Token provided to the ATM for further proof of authentication.
	AuthenticationToken *Max140Binary `xml:"AuthntcnTkn,omitempty"`
}

func (t *TransactionVerificationResult5) SetMethod(value string) {
	t.Method = (*AuthenticationMethod7Code)(&value)
}

func (t *TransactionVerificationResult5) SetVerificationEntity(value string) {
	t.VerificationEntity = (*AuthenticationEntity2Code)(&value)
}

func (t *TransactionVerificationResult5) SetResult(value string) {
	t.Result = (*Verification1Code)(&value)
}

func (t *TransactionVerificationResult5) SetAdditionalResult(value string) {
	t.AdditionalResult = (*Max500Text)(&value)
}

func (t *TransactionVerificationResult5) SetAuthenticationToken(value string) {
	t.AuthenticationToken = (*Max140Binary)(&value)
}
