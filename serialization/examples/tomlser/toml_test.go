package tomlser

import (
	"fmt"
	"sort"

	"github.com/spf13/viper"
)

// This example uses viper to read toml file
func Example_viper() {
	viper.SetConfigName("test")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	viper.ReadInConfig()

	keys := viper.AllKeys()
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("Key: %v Value: %v\n", key, viper.Get(key))
	}

	// Output:
	// Key: clients.data Value: [[gamma delta] [1 2]]
	// Key: clients.hosts Value: [alpha omega]
	// Key: database.connection_max Value: 5000
	// Key: database.enabled Value: true
	// Key: database.ports Value: [8000 8001 8002]
	// Key: database.server Value: 192.168.1.1
	// Key: owner.dob Value: 1979-05-27 07:32:00 -0800 -0800
	// Key: owner.name Value: Tom Preston-Werner
	// Key: servers.alpha.dc Value: eqdc10
	// Key: servers.alpha.ip Value: 10.0.0.1
	// Key: servers.beta.dc Value: eqdc10
	// Key: servers.beta.ip Value: 10.0.0.2
	// Key: title Value: TOML Example
}
