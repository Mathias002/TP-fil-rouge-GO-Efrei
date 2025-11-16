# 📇 TP Fil Rouge GO - Mini CRM en Ligne de Commande

Un gestionnaire de contacts robuste et modulaire développé en Go, offrant plusieurs modes de stockage et une interface en ligne de commande.

![Go Version](https://img.shields.io/badge/Go-1.25.3-00ADD8?style=flat&logo=go)
![EFREI](https://img.shields.io/badge/EFREI-M2-blue)

## 🎯 Fonctionnalités

- ✅ **CRUD complet** : Créer, Lire, Mettre à jour et Supprimer des contacts
- 🔄 **Multi-stockage** : 3 modes de persistance au choix
  - **Memory** : Stockage en mémoire (éphémère) avec menu interactif
  - **JSON** : Persistance dans un fichier `contacts.json`
  - **GORM/SQLite** : Base de données SQL robuste dans `contacts.db`
- ⚙️ **Configuration externe** : Changez de mode de stockage sans recompiler grâce à Viper
- 🎨 **CLI professionnelle** : Interface en ligne de commande intuitive via Cobra
- ✏️ **Mode interactif** : Menu infini pour gérer vos contacts en mémoire
- ✉️ **Validation** : Vérification du format d'email

## 📁 Architecture du projet

```
TP-fil-rouge-GO-Efrei/
├── cmd/
│   ├── utils/
│   │   └── validation.go      # Validation des entrées
│   ├── add.go                  # Commande d'ajout
│   ├── delete.go               # Commande de suppression
│   ├── get.go                  # Affichage d'un contact
│   ├── interactive.go          # Mode interactif (memory)
│   ├── list.go                 # Liste des contacts
│   ├── root.go                 # Commande racine + init
│   └── update.go               # Mise à jour de contact
├── internal/
│   ├── config/
│   │   └── config.go           # Gestion config Viper
│   ├── database/
│   │   └── db.go               # Connexion GORM
│   ├── main_menu/
│   │   └── main_menu.go        # Menu interactif
│   └── store/
│       ├── storage.go          # Interface Storer
│       ├── gorm.go             # Store GORM/SQLite
│       ├── json.go             # Store JSON
│       └── memory.go           # Store Memory
├── config.yaml                 # Configuration
├── main.go                     # Point d'entrée
├── go.mod
└── README.md
```

## 🚀 Installation

### Prérequis

- Go 1.25.3 ou supérieur
- Git

### Cloner le projet

```bash
git clone https://github.com/Mathias002/TP-fil-rouge-GO-Efrei.git
cd TP-fil-rouge-GO-Efrei
```

### Installer les dépendances

```bash
go mod download
```

### Compiler l'application

```bash
go build -o crm-fil-rouge.exe .
```

## ⚙️ Configuration

Le fichier `config.yaml` permet de configurer le mode de stockage :

```yaml
storage:
  # Choisir le type de stockage (memory | json | gorm)
  type: gorm
  
  # Fichier pour le stockage JSON
  file: contacts.json

database:
  # Fichier de la base de données SQLite
  name: contacts.db
```

### Types de stockage disponibles

| Type | Description | Persistance | Fichier |
|------|-------------|-------------|---------|
| `memory` | Stockage en mémoire | ❌ Non | - |
| `json` | Fichier JSON | ✅ Oui | `contacts.json` |
| `gorm` | Base SQLite via GORM | ✅ Oui | `contacts.db` |

## 💻 Utilisation

### Mode interactif (Memory)

Lance un menu interactif pour gérer les contacts en mémoire :

```bash
# Dans config.yaml : storage.type = memory
./crm-fil-rouge.exe interact

--- Menu CRM ---
1. Ajouter un contact
2. Lister les contacts
3. Modifier un contact
4. Supprimer un contact
5. Quitter
Votre choix: 1

Entrez le nom du contact: David
Entrez l'email du contact: david@test.com
✅ Contact ajouté!

# Les données sont perdues à la fermeture du programme
```

### Commandes CLI (JSON / GORM)

#### Ajouter un contact

```bash
# Avec flags
./crm-fil-rouge.exe add --name "Alice Martin" --email "alice.martin@example.com"

# Mode interactif (demande nom et email)
./crm-fil-rouge.exe add
```

#### Lister tous les contacts

```bash
./crm-fil-rouge.exe list
```

**Sortie :**
```
--- Liste des contacts ---
ID: 1 | Nom: Alice Martin | Email: alice.martin@example.com
ID: 2 | Nom: Bob Dupont | Email: bob.dupont@example.com
```

#### Afficher un contact spécifique

```bash
./crm-fil-rouge.exe get --id 1
```

#### Mettre à jour un contact

```bash
# Avec flags (les champs non fournis sont demandés)
./crm-fil-rouge.exe update --id 1 --name "Alice Durand"

# Mode interactif complet
./crm-fil-rouge.exe update --id 1
```

#### Supprimer un contact

```bash
./crm-fil-rouge.exe delete --id 1
```

## 📖 Exemples d'utilisation

### Workflow complet avec JSON

```bash
# 1. Configurer le mode JSON dans config.yaml
# storage.type: json

# 2. Ajouter des contacts
./crm-fil-rouge.exe add --name "Alice" --email "alice@test.com"
./crm-fil-rouge.exe add --name "Bob" --email "bob@test.com"

# 3. Lister les contacts
./crm-fil-rouge.exe list
# ID: 1234 | Nom: Alice | Email: alice@test.com
# ID: 5678 | Nom: Bob | Email: bob@test.com

# 4. Modifier un contact
./crm-fil-rouge.exe update --id 1234 --email "alice.martin@test.com"

# 5. Supprimer un contact
./crm-fil-rouge.exe delete --id 5678

# Le fichier contacts.json est automatiquement mis à jour
```

## 🛠️ Technologies utilisées

- **[Go 1.25.3](https://golang.org/)** - Langage de programmation
- **[Cobra](https://github.com/spf13/cobra)** - Framework CLI
- **[Viper](https://github.com/spf13/viper)** - Gestion de configuration
- **[GORM](https://gorm.io/)** - ORM pour Go
- **[SQLite](https://www.sqlite.org/)** - Base de données embarquée

## 🏗️ Concepts Go utilisés

Ce projet met en pratique les concepts avancés de Go :

- ✅ **Interfaces** - Architecture modulaire avec `Storer`
- ✅ **Injection de dépendances** - Découplage via interfaces
- ✅ **Struct et méthodes** - POO en Go
- ✅ **Gestion d'erreurs** - `if err != nil`
- ✅ **Package organization** - Structure de projet professionnelle
- ✅ **JSON marshaling/unmarshaling** - Sérialisation
- ✅ **ORM patterns** - GORM avec SQLite
- ✅ **Configuration externe** - Viper YAML
- ✅ **CLI patterns** - Cobra commands et flags

## 📊 Schéma de l'architecture

```
┌─────────────┐
│   main.go   │
└──────┬──────┘
       │
       ▼
┌─────────────────┐
│   cmd/root.go   │ ◄─── Viper (config.yaml)
└────────┬────────┘
         │
         ├──► cmd/add.go
         ├──► cmd/list.go
         ├──► cmd/update.go
         ├──► cmd/delete.go
         └──► cmd/interactive.go
                │
                ▼
         ┌──────────────┐
         │ Storer (interface)
         └──────┬───────┘
                │
    ┌───────────┼───────────┐
    ▼           ▼           ▼
┌─────────┐ ┌─────────┐ ┌─────────┐
│ Memory  │ │  JSON   │ │  GORM   │
│ Store   │ │ Store   │ │ Store   │
└─────────┘ └─────────┘ └─────────┘
     │           │           │
     ▼           ▼           ▼
  (RAM)   contacts.json  contacts.db
```

## 🎓 Progression du projet

### ✅ Étape 1 : Base sans persistance
- Structure de données avec `map`
- CRUD basique en mémoire

### ✅ Étape 2 : Architecture modulaire
- 2.1 : Ajout de `struct` pour la modélisation
- 2.2 : Interface `Storer` pour la modularité

### ✅ Étape 3 : CLI et persistance
- 3.1 : Transformation en CLI avec Cobra
- 3.2 : Persistance JSON

### ✅ Étape 4 : Base de données
- Intégration de GORM avec SQLite

### ✅ Étape 5 : Configuration externe
- Gestion de config avec Viper

## 🔮 Améliorations futures possibles

- [ ] Export/Import CSV
- [ ] Recherche avancée (par nom, email, date)
- [ ] Pagination des résultats
- [ ] API REST avec Gin
- [ ] Interface web
- [ ] Tests unitaires et d'intégration
- [ ] Chiffrement des données sensibles

## 🐛 Résolution de problèmes

### Le fichier de configuration n'est pas trouvé

```bash
# Vérifier que config.yaml est à la racine du projet
ls config.yaml

# Ou spécifier le chemin dans config.go
```

### Erreur "database is locked" avec SQLite

```bash
# Fermer toutes les connexions à la base de données
# Redémarrer l'application
```

### Les contacts ne persistent pas en mode memory

C'est normal ! Le mode `memory` est **éphémère**. Utilisez `json` ou `gorm` pour la persistance.

### Lecture du fichier .db SQLite

Vous avez deux options pour consulter le contenu de votre fichier SQLite :

#### Option 1 : Visualisation en ligne
Rendez-vous sur [ce site](https://sqliteviewer.app) qui permet de lire le contenu d'un fichier `.db` directement dans votre navigateur.

#### Option 2 : Extension IDE
Si vous utilisez un IDE compatible avec les extensions (comme VS Code), installez une extension dédiée à la lecture de bases de données SQLite.

**Exemple avec VS Code :**

![Exemple de visualisation SQLite dans VS Code](https://github.com/user-attachments/assets/620c734c-b35b-4360-90cd-abe37c1112fd)

## 👥 Auteur

**Mathias002** - Étudiant M2 EFREI  
[GitHub](https://github.com/Mathias002) | [Projet](https://github.com/Mathias002/TP-fil-rouge-GO-Efrei)

## 📄 Licence

Ce projet est un exercice pédagogique réalisé dans le cadre du cours de Go à l'EFREI Paris.

---

**Développé avec ❤️ en Go**
