# Mini CRM en Ligne de Commande 📇

Un gestionnaire de contacts simple et efficace développé en Go, permettant de gérer vos contacts via une interface en ligne de commande interactive ou par flags.

## 📋 Fonctionnalités

- ✅ **Menu interactif** en boucle pour une navigation intuitive
- ➕ **Ajouter un contact** (génération automatique d'ID unique)
- 📋 **Lister tous les contacts** avec affichage formaté
- ✏️ **Mettre à jour un contact** (nom et/ou email)
- 🗑️ **Supprimer un contact** par son ID
- 🚀 **Ajout via flags** pour une utilisation en ligne de commande
- 🔄 **Contacts par défaut** chargés au démarrage

## 🛠️ Concepts Go utilisés

Ce projet met en pratique les concepts fondamentaux de Go :

- `comma ok idiom` - Vérification d'existence dans les maps
- `for { }` - Boucle infinie pour le menu principal
- `switch` - Gestion des choix utilisateur
- `map` - Stockage des contacts avec accès rapide par ID
- `if err != nil` - Gestion des erreurs
- `strconv` - Conversion de types (string ↔ int)
- `os.Stdin` - Lecture des entrées utilisateur
- `bufio` - Lecture optimisée avec buffer
- `flag` - Parsing des arguments en ligne de commande

## 📁 Structure du projet

```
TP-fil-rouge-GO-efrei/
├── main.go              # Point d'entrée de l'application
├── contact/
│   └── contact.go       # Logique de gestion des contacts
├── main_menu/
│   └── main_menu.go     # Affichage et gestion du menu
├── go.mod               # Gestion des dépendances
└── README.md            # Documentation
```

## 🚀 Installation

### Prérequis

- Go 1.21 ou supérieur installé sur votre machine

### Cloner le projet

```bash
git clone https://github.com/Mathias002/TP-fil-rouge-GO-efrei.git
cd TP-fil-rouge-GO-efrei/cmd/api
```

### Compiler le projet

```bash
go build -o crm main.go
```

## 💻 Utilisation

### Mode interactif

Lancez l'application sans arguments pour accéder au menu interactif :

```bash
go run main.go
```

ou si vous avez compilé :

```bash
./crm
```

**Menu principal :**

```
--- Mini CRM ---
1. Ajouter un contact
2. Lister les contacts
3. Modifier un contact
4. Supprimer un contact
5. Quitter
Votre choix :
```

### Mode ligne de commande (flags)

Ajoutez un contact directement via des flags :

```bash
go run main.go -name "Jean Dupont" -email "jean.dupont@example.com"
```

**Flags disponibles :**

- `-name` : Nom du contact (obligatoire)
- `-email` : Email du contact (obligatoire)

## 📖 Exemples d'utilisation

### Ajouter un contact (mode interactif)

```
Votre choix : 1

--- Ajout d'un contact ---
Entrez le nom du contact :
Alice Martin
Entrez l'adresse email du contact :
alice.martin@example.com

Contact ajouté :
ID: 7234 | Nom: Alice Martin | Email: alice.martin@example.com
```

### Lister les contacts

```
Votre choix : 2

--- Liste des contacts ---
ID: 7234 | Name: Alice Martin | Email: alice.martin@example.com
ID: 9175 | Name: Poipoi | Email: poipoi@gmail.com
ID: 5241 | Name: Lala | Email: lala@gmail.com
```

### Modifier un contact

```
Votre choix : 3

--- Liste des contacts ---
ID: 7234 | Name: Alice Martin | Email: alice.martin@example.com

--- Mettre à jour un contact ---
Entrez l'ID du contact à mettre à jour:
7234
--- Début de la modification des informations de l'utilisateur.rice avec l'ID 7234 ---

Entrez le nouveau nom de l'utilisateur.rice (vide si pas de changement)
Alice Durand
Entrez le nouveau email de l'utilisateur.rice (vide si pas de changement)

 ✅✅✅ Utilisateur.rice avec l'ID 7234 modifié.e avec succés ✅✅✅
```

### Supprimer un contact

```
Votre choix : 4

--- Supprimer un contact ---
Entrez l'ID du contact à supprimer:
7234

 ✅✅✅ Utilisateur.rice avec l'ID 7234 supprimé.e avec succés ✅✅✅
```

## 🏗️ Architecture

### Types personnalisés

```go
type IDContact int
type NameContact string
type EmailContact string

type Contact struct {
    ID    IDContact
    Name  NameContact
    Email EmailContact
}
```

### Stockage

Les contacts sont stockés dans un `map` global :

```go
var Contacts map[IDContact]Contact
```

**Avantages :**

- Accès en O(1) par ID
- Unicité garantie des IDs
- Gestion simple des opérations CRUD

### Génération des IDs

Les IDs sont générés aléatoirement entre 0 et 9999 :

```go
func randomInteger() int {
    return rand.Intn(10000)
}
```

## ⚠️ Limitations actuelles

- Les données sont **stockées en mémoire uniquement** (non persistantes)
- Pas de vérification du format d'email
- Possible collision d'IDs (probabilité faible avec 10000 valeurs possibles)
- Pas de recherche par nom ou email

## 🔮 Améliorations futures

- [ ] Persistance des données (JSON, SQLite)
- [ ] Validation des emails
- [ ] IDs garantis uniques (UUID ou auto-incrémentation)
- [ ] Tests unitaires

## 👥 Auteur

**Mathias** - [GitHub](https://github.com/Mathias002)

## 📄 Licence

Ce projet est un exercice pédagogique réalisé dans le cadre du cours de Go à l'EFREI.
