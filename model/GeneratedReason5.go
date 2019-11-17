package model

// Specifies the reason why the transaction was generated.
type GeneratedReason5 struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReasons5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (g *GeneratedReason5) AddCode() *GeneratedReasons5Choice {
	g.Code = new(GeneratedReasons5Choice)
	return g.Code
}

func (g *GeneratedReason5) SetAdditionalReasonInformation(value string) {
	g.AdditionalReasonInformation = (*Max210Text)(&value)
}
