package database

import (
	"log"

	"github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// DB est l'instance globale de notre connexion à la base de données
var DB *gorm.DB

// ConnectDB initialise la connexion à la base de données et exécute les migrations
func ConnectDB() *gorm.DB {
	if DB != nil {
		return DB // Retourner la connexion existante
	}

	var err error
	dbName := config.Config.Database.Name
	// log.Printf("Tentative de connexion à la base de données : %s", dbName)

	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Échec de la connexion à la base de données '%s': %v", dbName, err)
	}

	// log.Println("Connexion à la base de données SQLite réussie !")

	return DB
}

func MigrateModels(models ...interface{}) error {
	if DB == nil {
		log.Fatal("La base de données n'est pas initialisée")
	}

	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Printf("Échec de la migration: %v", err)
		return err
	}

	log.Println("Migration de la base de données réussie !")
	return nil
}
