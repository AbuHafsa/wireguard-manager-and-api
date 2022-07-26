module gitlab.com/raspberry.tech/wireguard-manager-and-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/go-co-op/gocron v1.6.2
	github.com/gorilla/mux v1.8.0
	github.com/spf13/viper v1.9.0
	github.com/vishvananda/netlink v1.1.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0
	golang.zx2c4.com/wireguard/wgctrl v0.0.0-20210506160403-92e472f520a5
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.12
	hyperlite v0.2.1
)

replace hyperlite => ../go/hyperlite
