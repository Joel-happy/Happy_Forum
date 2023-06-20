package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // Déclarer une variable globale pour la connexion à la base de données

func main() {
	// Servir les fichiers statiques du répertoire "frontend"
	fs := http.FileServer(http.Dir("frontend"))
http.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	// Chaîne de connexion à la base de données MySQL
	connectionString := "root:@tcp(localhost:3306)/projetforum"

	// Ouvrir une connexion à la base de données
	var err error
	db, err = sql.Open("mysql", connectionString) // Assigner la connexion ouverte à la variable globale db
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Vérifier la connexion à la base de données
	// Test de la connexion à la base de données
err = db.Ping()
if err != nil {
    log.Fatal("Erreur lors de la connexion à la base de données:", err)
} else {
    log.Println("Connexion à la base de données réussie")
}

	// Créer un routeur Gin
	router := gin.Default()

	// Charger les fichiers HTML du répertoire "frontend"
	router.LoadHTMLGlob(`..\frontend\*.html`)

	router.POST("/inscription", handleInscription)
	router.POST("/connexion", handleConnexion)
	router.GET("/accueil", handleAccueil)

	// Démarrer le serveur HTTP
	router.Run(":8080")
}

func handleInscription(c *gin.Context) {
	// Récupérer les données du formulaire d'inscription
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Vérifier si les données sont dupliquées
	isDuplicate, err := verification(username, email)
	if err != nil {
		// Gérer l'erreur
		c.String(http.StatusInternalServerError, "Une erreur s'est produite lors de la vérification des données")
		return
	}

	if isDuplicate {
		// Afficher un message d'erreur indiquant que les données sont déjà utilisées
		c.String(http.StatusBadRequest, "Le pseudo ou l'adresse e-mail est déjà utilisé(e)")
		return
	}

	// Insérer les données dans la base de données
	err = insertUser(username, email, password)
	if err != nil {
		// Gérer l'erreur
		c.String(http.StatusInternalServerError, "Une erreur s'est produite lors de l'inscription")
		return
	}

	// Afficher un message de confirmation
	c.String(http.StatusOK, "Vous êtes maintenant inscrit sur le forum")
}

func verification(username string, email string) (bool, error) {
	query := "SELECT COUNT(*) FROM utilisateur WHERE Pseudo = ? OR Adresse_e_mail = ?"
	var count int
	err := db.QueryRow(query, username, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func insertUser(username string, email string, password string) error {
	query := "INSERT INTO utilisateur (Pseudo, Adresse_e_mail, Mot_de_passe) VALUES (?, ?, ?)"
	_, err := db.Exec(query, username, email, password)
	if err != nil {
		return err
	}

	return nil
}

func checkCredentials(username string, password string) (bool, error) {
	query := "SELECT COUNT(*) FROM utilisateur WHERE Pseudo = ? AND Mot_de_passe = ?"
	var count int
	err := db.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func handleConnexion(c *gin.Context) {
	// Récupérer les données du formulaire de connexion
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Vérifier les informations de connexion
	isValid, err := checkCredentials(username, password)
	if err != nil {
		// Gérer l'erreur
		c.String(http.StatusInternalServerError, "Une erreur s'est produite lors de la vérification des informations de connexion")
		return
	}

	if isValid {
		// Connexion réussie, rediriger vers la page d'accueil par exemple
		c.Redirect(http.StatusFound, "/frontend/accueil.html")
	} else {
		// Connexion échouée, afficher un message d'erreur
		c.String(http.StatusUnauthorized, "Identifiant ou mot de passe incorrect")
	}
}

func handleAccueil(c *gin.Context) {
	// Charger et afficher la page d'accueil
	c.HTML(http.StatusOK, "accueil.html", nil)
}
