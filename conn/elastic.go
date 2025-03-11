package conn

import (
	"TobiasFP/BotNana/config"
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

// GetElasticDB returns a pointer to the standard NoSQL db.
func GetElasticDB() (*elasticsearch.TypedClient, error) {
	config := config.GetConfig()

	server := config.GetString("elastic.SERVER")
	port := config.GetString("elastic.PORT")
	userName := config.GetString("elastic.USERNAME")
	password := config.GetString("elastic.PASSWORD")
	useCert := config.GetBool("elastic.USECERT")
	scheme := "http://"
	if useCert {
		scheme = "https://"
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			scheme + server + ":" + port,
		},
		Username: userName,
		Password: password,
	}
	if useCert {
		certFilePath := config.GetString("elastic.CERTFILEPATH")

		cert, _ := os.ReadFile(certFilePath)
		cfg.CACert = cert
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	return es, err
}
