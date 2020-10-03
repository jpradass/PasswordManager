package remotedbadapter

import (
	"fmt"

	"github.com/PasswordManager/configuration"
)

func connectPG(conf *configuration.Configuration) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf, conf.DB)
}
