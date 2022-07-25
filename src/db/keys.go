package db

type Key struct {
	KeyID        int    `gorm:"primaryKey;autoIncrement"`
	PublicKey    string `gorm:"unique"`
	PresharedKey string `gorm:"unique"`
	IPv4Address  string `gorm:"foreignKey:IPv4Address"`
	Enabled      string
}

func (p *Repository) FindKeys() ([]Key, error) {
	var keys []Key

	exec := p.Db.Find(&keys)

	return keys, exec.Error

}

func (p *Repository) CreateKey(publicKey string, presharedKey string) (Key, error) {
	var key Key

	ip, err := p.FindIPNotInUse()
	if err != nil {
		return key, err
	}

	key = Key{
		PublicKey:    publicKey,
		PresharedKey: presharedKey,
		IPv4Address:  ip.IPv4Address,
		Enabled:      "true",
	}
	exec := p.Db.Create(&key)
	err = exec.Error
	if err != nil {
		return key, err
	}

	err = p.SetIPInUse(&ip)

	return key, err
}
