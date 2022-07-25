package db

import (
	"net"

	"github.com/spf13/viper"
)

type IP struct {
	IPv4Address string `gorm:"primaryKey"`
	IPv6Address string `gorm:"unique"`
	InUse       string
	WGInterface string
}

func (p *Repository) FindIPNotInUse() (IP, error) {
	var ip IP

	exec := p.Db.Where("in_use = ?", "false").First(&ip) //find IP not in use

	return ip, exec.Error
}

func (p *Repository) SetIPInUse(ip *IP) error {
	ip.InUse = "true"

	exec := p.Db.Save(ip)

	return exec.Error
}

func pregenIP(ipv4 []byte, maxIP int, wgInterface string) []IP {
	var ips []IP
	a := ipv4[0]
	b := ipv4[1]
	c := ipv4[2]
	d := ipv4[3] + 1
	for i := 0; i < maxIP; i++ {
		d++
		if d == 0xF3 {
			c++
			d = 0x00
		}
		if c == 0xFF {
			b++
		}
		l := net.IPv4(a, b, c, d)
		ips = append(ips, IP{
			IPv4Address: l.String(),
			IPv6Address: "",
			InUse:       "false",
			WGInterface: wgInterface,
		})
	}

	return ips
}

func (p *Repository) PregenIPv4(wgInterface string) error {
	maxIP := viper.GetInt("SERVER.MAX_IP")
	configIPv4 := viper.GetString("INSTANCE.IP.LOCAL.IPV4.ADDRESS")

	ip := net.ParseIP(configIPv4)
	ipv4 := ip[12:]

	if ipv4[0] != 0x0A {
		panic("You must use a class A network number i.e. 10.x.x.x")
	}

	ips := pregenIP(ipv4, maxIP, wgInterface)
	exec := p.Db.CreateInBatches(&ips, 100)

	return exec.Error
}
