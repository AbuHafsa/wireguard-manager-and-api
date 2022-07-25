package db

type WireguardInterface struct {
	InterfaceName string `gorm:"primaryKey"`
	PrivateKey    string `gorm:"unique"`
	PublicKey     string `gorm:"unique"`
	ListenPort    int    `gorm:"unique"`
	IPv4Address   string
	IPv6Address   string
}
type Subscription struct {
	KeyID             int    `gorm:"foreignKey:KeyID"`
	PublicKey         string `gorm:"foreignKey:PublicKey"`
	BandwidthUsed     int64
	BandwidthAllotted int64
	SubscriptionEnd   string
}
