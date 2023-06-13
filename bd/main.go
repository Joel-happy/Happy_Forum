package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Établir une connexion à la base de données
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/projetforum")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Générer des fausses données de catégories de jeux vidéo
	categories := []string{"Action", "Aventure", "Stratégie", "RPG", "Sport", "Simulation", "FPS", "Course"}

	// Insérer les fausses données de catégories de jeux vidéo dans la base de données
	for _, category := range categories {
		query := "INSERT INTO catégorie (Nom_de_la_catégorie) VALUES (?)"
		_, err := db.Exec(query, category)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Générer des fausses données de grades
	grades := []string{"Débutant", "Confirmé", "Expert"}

	// Insérer les fausses données de grades dans la base de données
	for _, grade := range grades {
		query := "INSERT INTO grade (Nom_du_Grade) VALUES (?)"
		_, err := db.Exec(query, grade)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Générer des fausses données d'utilisateurs
	users := []struct {
		Pseudo       string
		Grade        int
		AdresseEmail string
		MotDePasse   string
	}{
		{Pseudo: "Utilisateur1", Grade: 1, AdresseEmail: "utilisateur1@example.com", MotDePasse: "password1"},
		{Pseudo: "Utilisateur2", Grade: 2, AdresseEmail: "utilisateur2@example.com", MotDePasse: "password2"},
		{Pseudo: "Utilisateur3", Grade: 3, AdresseEmail: "utilisateur3@example.com", MotDePasse: "password3"},
		{Pseudo: "Utilisateur4", Grade: 1, AdresseEmail: "utilisateur4@example.com", MotDePasse: "password4"},
		{Pseudo: "Utilisateur5", Grade: 2, AdresseEmail: "utilisateur5@example.com", MotDePasse: "password5"},
	}

	// Insérer les fausses données d'utilisateurs dans la base de données
	for _, user := range users {
		query := "INSERT INTO utilisateur (Pseudo, Grade, Adresse_e_mail, Mot_de_passe) VALUES (?, ?, ?, ?)"
		_, err := db.Exec(query, user.Pseudo, user.Grade, user.AdresseEmail, user.MotDePasse)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Générer des fausses données de discussions
	discussions := []struct {
		NomDiscussion string
		Sujet         string
	}{
		{NomDiscussion: "Discussion1", Sujet: "Nouveau jeu à venir"},
		{NomDiscussion: "Discussion2", Sujet: "Stratégies pour un jeu en ligne"},
		{NomDiscussion: "Discussion3", Sujet: "Meilleurs jeux de l'année"},
		{NomDiscussion: "Discussion4", Sujet: "Problèmes techniques"},
		{NomDiscussion: "Discussion5", Sujet: "Conseils pour les débutants"},
	}

	// Insérer les fausses données de discussions dans la base de données
	for _, discussion := range discussions {
		query := "INSERT INTO discussion (Nom_de_la_discussion, Sujet) VALUES (?, ?)"
		_, err := db.Exec(query, discussion.NomDiscussion, discussion.Sujet)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Générer des fausses données de messages et réponses
	messageCount := 0
	rand.Seed(time.Now().Unix())

	for i := 1; i <= 50; i++ {
		// Générer un message pour une discussion aléatoire
		message := fmt.Sprintf("Message %d pour la discussion %d", i, rand.Intn(len(discussions))+1)

		query := "INSERT INTO message (Texte, Date_du_message) VALUES (?, NOW())"
		_, err := db.Exec(query, message)
		if err != nil {
			log.Fatal(err)
		}

		messageCount++

		// Générer une réponse pour le message
		response := fmt.Sprintf("Réponse %d pour le message %d", i, messageCount)

		query = "INSERT INTO réponse (Texte) VALUES (?)"
		_, err = db.Exec(query, response)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Données cohérentes pour le forum de jeu vidéo insérées avec succès dans la base de données !")
}
