package db_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spf13/viper"
	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var repository *db.Repository

func setupMockDb() (*sql.DB, sqlmock.Sqlmock, error) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	// defer db.Close()

	gdb, err := gorm.Open(sqlite.Dialector{Conn: mockDb}, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, nil, err
	}

	gdb.AutoMigrate(&db.Key{}, &db.IP{})

	repository = &db.Repository{Db: gdb}

	return mockDb, mock, nil
}

func setUpConfigurationFile() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/opt/wgManagerAPI/")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic("Unable to read configuration file")
	}
}

func Test_Returns_Empty_List_Of_Keys_When_Query_For_All_Keys(t *testing.T) {
	mockDb, mock, err := setupMockDb()
	if err != nil {
		t.Fatalf("An error has occurred, database mock set-up %v", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `keys`")).WillReturnRows(sqlmock.NewRows(nil))

	l, err := repository.FindKeys()
	if err != nil {
		t.Fatalf("An error has occurred, retrieving keys from database %v", err)
	}
	if len(l) != 0 {
		t.Errorf("expected keys to be empty, actual size of returned keys %d", len(l))
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("An error has occurred, expectations were not met %v", err)
	}

	mockDb.Close()
}

func Test_Generate_IPv4(t *testing.T) {
	mockDb, mock, err := setupMockDb()
	if err != nil {
		t.Fatalf("An error has occurred, database mock set-up %v", err)
	}

	setUpConfigurationFile()
	viper.Set("SERVER.MAX_IP", 1)

	ip := db.IP{
		IPv4Address: "10.6.0.3",
		IPv6Address: "",
		InUse:       "false",
		WGInterface: "wg0",
	}

	sql := "INSERT INTO `ips`"

	mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(ip.IPv4Address, ip.IPv6Address, ip.InUse, ip.WGInterface).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repository.PregenIPv4("wg0")
	if err != nil {
		t.Fatalf("An error has occurred, pre-generating IPs %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("An error has occurred, expectations were not met, %v", err)
	}

	mockDb.Close()
}
