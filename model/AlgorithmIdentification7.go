package model

// Cryptographic algorithms and parameters for the protection of transported keys by an asymmetric key.
type AlgorithmIdentification7 struct {

	// Asymmetric encryption algorithm of a transport key.
	Algorithm *Algorithm7Code `xml:"Algo"`

	// Parameters of the RSAES-OAEP encryption algorithm (RSA Encryption Scheme: Optimal Asymmetric Encryption Padding).
	Parameter *Parameter2 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification7) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm7Code)(&value)
}

func (a *AlgorithmIdentification7) AddParameter() *Parameter2 {
	a.Parameter = new(Parameter2)
	return a.Parameter
}
