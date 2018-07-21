package conf

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

const (
	ConfigFilePostfix = "_CONF_FILE"
)

func FromFile(envPrefix string) (*oauth2.Config, error) {
	filename := getFilename(envPrefix)
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	conf := &oauth2.Config{}
	jd := json.NewDecoder(configFile)
	if err = jd.Decode(conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func getFilename(prefix string) string {
	return os.Getenv(prefix + ConfigFilePostfix)
}
