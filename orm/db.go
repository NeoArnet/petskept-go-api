package orm

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() {
	//DNS=yp6aoj2jr8kafyqq3vet:pscale_pw_Hl35PiAWZbAgr6MLDUfjBMd8HtCG69lmqc7Dj0zvkqS@tcp(aws.connect.psdb.cloud)/petskept?tls=true&interpolateParams=true&charset=utf8mb4&parseTime=True&loc=Local
	dsn := os.Getenv("DNS")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	Db.AutoMigrate(&User{})
}
