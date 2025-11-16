package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Server struct {
		Port int
	} `mapstructure:"server"`
	Database struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"database"`
}

// Config est l'instance globale de notre configuration
var Config AppConfig

func InitConfig() {
	// Nom du fichier de configuration (sans extension) et chemins de recherche
	viper.SetConfigName("config")
	viper.AddConfigPath(".")        // Cherche dans le répertoire courant
	viper.AddConfigPath("./config") // Cherche dans un sous-répertoire 'config'

	// Lire le fichier de configuration
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Println("Fichier de configuration 'config.yaml' non trouvé, utilisation des valeurs par défaut ou des variables d'environnement.")
		}
	} else {
		// log.Println("Fichier de configuration 'config.yaml' chargé avec succès.")
	}

	// Désérialiser la configuration lue dans la struct AppConfig
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Impossible de désérialiser la configuration: %v", err)
	}

	// log.Printf("Configuration chargée : Port=%d, DBName=%s", Config.Server.Port, Config.Database.Name)
}
