package model

// Acceptor parameters dedicated to the merchant.
type MerchantConfigurationParameters1 struct {

	// Identification of the merchant for the MTM, if the POI manages several merchants.
	MerchantIdentification *Max35Text `xml:"MrchntId,omitempty"`

	// Version of the merchant parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Merchant parameters.
	Parameters *Max10000Binary `xml:"Params"`
}

func (m *MerchantConfigurationParameters1) SetMerchantIdentification(value string) {
	m.MerchantIdentification = (*Max35Text)(&value)
}

func (m *MerchantConfigurationParameters1) SetVersion(value string) {
	m.Version = (*Max256Text)(&value)
}

func (m *MerchantConfigurationParameters1) SetParameters(value string) {
	m.Parameters = (*Max10000Binary)(&value)
}
