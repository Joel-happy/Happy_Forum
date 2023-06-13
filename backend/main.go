package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:8080)/projetforum")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := http.NewServeMux()

	router.HandleFunc("/inscription", handleInscription)
	router.HandleFunc("/connexion", handleConnexion)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleInscription(w http.ResponseWriter, r *http.Request) {
	// Récupérer les données du formulaire d'inscription
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Vérifier si les données sont dupliquées
	isDuplicate, err := verification(username, email)
	if err != nil {
		// Gérer l'erreur
		http.Error(w, "Une erreur s'est produite lors de la vérification des données", http.StatusInternalServerError)
		return
	}

	if isDuplicate {
		// Afficher un message d'erreur indiquant que les données sont déjà utilisées
		fmt.Fprintf(w, "Le pseudo ou l'adresse e-mail est déjà utilisé(e)")
		return
	}

	// Insérer les données dans la base de données
	err = insertUser(username, email, password)
	if err != nil {
		// Gérer l'erreur
		http.Error(w, "Une erreur s'est produite lors de l'inscription", http.StatusInternalServerError)
		return
	}

	// Afficher un message de confirmation
	fmt.Fprintf(w, "Vous êtes maintenant inscrit sur le forum")

	// Redirection ou réponse en fonction du résultat de l'inscription
	// ...
}

func verification(username string, email string) (bool, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:8080)/projetforum")
	if err != nil {
		return false, err
	}
	defer db.Close()

	query := "SELECT COUNT(*) FROM utilisateur WHERE Pseudo = ? OR Adresse_e_mail = ?"
	var count int
	err = db.QueryRow(query, username, email).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		// Les données sont déjà présentes dans la base de données
		return true, nil
	}

	// Les données sont uniques
	return false, nil
}

func insertUser(username string, email string, password string) error {
	db, err := sql.Open("mysql", "root:@tcp(localhost:8080)/projetforum")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO utilisateur (Pseudo, Adresse_e_mail, Mot_de_passe) VALUES (?, ?, ?)"
	_, err = db.Exec(query, username, email, password)
	if err != nil {
		return err
	}

	return nil
}

func checkCredentials(username string, password string) (bool, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:8080)/projetforum")
	if err != nil {
		return false, err
	}
	defer db.Close()

	query := "SELECT COUNT(*) FROM utilisateur WHERE Pseudo = ? AND Mot_de_passe = ?"
	var count int
	err = db.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func handleConnexion(w http.ResponseWriter, r *http.Request) {
	// Récupérer les données du formulaire de connexion
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Vérifier les informations de connexion
	isValid, err := checkCredentials(username, password)
	if err != nil {
		// Gérer l'erreur
		http.Error(w, "Une erreur s'est produite lors de la vérification des informations de connexion", http.StatusInternalServerError)
		return
	}

	if isValid {
		// Connexion réussie, rediriger vers la page d'accueil par exemple
		http.Redirect(w, r, "/accueil", http.StatusFound)
	} else {
		// Connexion échouée, afficher un message d'erreur
		fmt.Fprintf(w, "Identifiant ou mot de passe incorrect")
	}
}
