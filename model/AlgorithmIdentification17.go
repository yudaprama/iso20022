package model

// Identification of a cryptographic algorithm and parameters for digital signatures.
type AlgorithmIdentification17 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm14Code `xml:"Algo"`

	// Parameters of the RSASSA-PSS digital signature algorithm (RSA signature algorithm with appendix: Probabilistic Signature Scheme).
	Parameter *Parameter8 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification17) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm14Code)(&value)
}

func (a *AlgorithmIdentification17) AddParameter() *Parameter8 {
	a.Parameter = new(Parameter8)
	return a.Parameter
}
