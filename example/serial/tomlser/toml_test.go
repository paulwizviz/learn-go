package tomlser

import (
	"log"

	"github.com/spf13/viper"
)

// This example uses viper to read toml file
func Example_viper() {
	viper.SetConfigName("test")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	viper.ReadInConfig()

	keys := viper.AllKeys()
	for _, key := range keys {
		log.Println(key)
	}

	// Output:
}
