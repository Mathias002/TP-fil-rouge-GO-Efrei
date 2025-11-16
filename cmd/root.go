package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/config"
	"github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/database"
	storage "github.com/Mathias002/TP-fil-rouge-GO-efrei/internal/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	store storage.Storer
)

var rootCmd = &cobra.Command{
	Use:   "crm-fil-rouge",
	Short: "crm-fil-rouge est un gestionaire de contact",
	Long:  `crm-fil-rouge est un gestionaire de contact qui permet de consulter, d'ajouter, de modifier et de supprimer des contact avec une vérifications des données et une gestion des erreurs`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		TypeStorage := viper.GetString("storage.type")

		// fmt.Printf("type de storage: %v \n", typeStorage)

		switch TypeStorage {
		case "json":
			fileName := viper.GetString("storage.file")
			store = storage.NewJSONStore(fileName)
		case "memory":
			store = storage.NewMemoryStore()
		case "gorm":
			// 1. Initialiser la configuration
			config.InitConfig()

			// 2. Connecter à la base de données et exécuter les migrations
			database.ConnectDB()

			// 3. initialisation du store gorm
			var err error
			store, err = storage.NewGormStore()
			if err != nil {
				log.Fatalf("❌ Erreur lors de la création du store: %v", err)
			}

		default:
			log.Fatal("Aucun type de storage défini")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initViperConfig)
}

func initViperConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier config %s", err)
	}
}
