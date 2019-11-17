package model

// Acceptor parameters dedicated to the merchant.
type MerchantConfigurationParameters2 struct {

	// Type of action for the configuration parameters.
	ActionType *TerminalManagementAction3Code `xml:"ActnTp"`

	// Identification of the merchant for the MTM, if the POI manages several merchants.
	MerchantIdentification *Max35Text `xml:"MrchntId,omitempty"`

	// Version of the merchant parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Local proxy configuration.
	Proxy *NetworkParameters6 `xml:"Prxy,omitempty"`

	// Other merchant parameters.
	OtherParameters *Max10000Binary `xml:"OthrParams,omitempty"`
}

func (m *MerchantConfigurationParameters2) SetActionType(value string) {
	m.ActionType = (*TerminalManagementAction3Code)(&value)
}

func (m *MerchantConfigurationParameters2) SetMerchantIdentification(value string) {
	m.MerchantIdentification = (*Max35Text)(&value)
}

func (m *MerchantConfigurationParameters2) SetVersion(value string) {
	m.Version = (*Max256Text)(&value)
}

func (m *MerchantConfigurationParameters2) AddProxy() *NetworkParameters6 {
	m.Proxy = new(NetworkParameters6)
	return m.Proxy
}

func (m *MerchantConfigurationParameters2) SetOtherParameters(value string) {
	m.OtherParameters = (*Max10000Binary)(&value)
}
