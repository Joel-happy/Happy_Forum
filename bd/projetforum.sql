-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : mer. 07 juin 2023 à 11:04
-- Version du serveur : 8.0.31
-- Version de PHP : 8.0.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `projetforum`
--

-- --------------------------------------------------------

--
-- Structure de la table `catégorie`
--

DROP TABLE IF EXISTS `catégorie`;
CREATE TABLE IF NOT EXISTS `catégorie` (
  `Id_Catégorie` int NOT NULL,
  `Nom_de_la_catégorie` varchar(50) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  PRIMARY KEY (`Id_Catégorie`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

-- --------------------------------------------------------

--
-- Structure de la table `discussion`
--

DROP TABLE IF EXISTS `discussion`;
CREATE TABLE IF NOT EXISTS `discussion` (
  `Id_Discussion` int NOT NULL,
  `Nom_de_la_discussion` varchar(50) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `Sujet` varchar(50) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `Nombre_d_utilisateur` decimal(15,2) DEFAULT NULL,
  `Date_du_débat` datetime DEFAULT NULL,
  `Date_de_fin` datetime DEFAULT NULL,
  PRIMARY KEY (`Id_Discussion`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

-- --------------------------------------------------------

--
-- Structure de la table `grade`
--

DROP TABLE IF EXISTS `grade`;
CREATE TABLE IF NOT EXISTS `grade` (
  `Id_Grade` int NOT NULL,
  `Nom_du_Grade` varchar(50) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  PRIMARY KEY (`Id_Grade`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

-- --------------------------------------------------------

--
-- Structure de la table `message`
--

DROP TABLE IF EXISTS `message`;
CREATE TABLE IF NOT EXISTS `message` (
  `Id_Message` int NOT NULL,
  `Texte` varchar(50) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `Date_du_message` date DEFAULT NULL,
  `Id_Reponse` int DEFAULT NULL,
  PRIMARY KEY (`Id_Message`),
  KEY `Id_Reponse` (`Id_Reponse`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

-- --------------------------------------------------------

--
-- Structure de la table `réponse`
--

DROP TABLE IF EXISTS `réponse`;
CREATE TABLE IF NOT EXISTS `réponse` (
  `Id_Reponse` int NOT NULL,
  `Texte` varchar(50) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  PRIMARY KEY (`Id_Reponse`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

-- --------------------------------------------------------

--
-- Structure de la table `utilisateur`
--

DROP TABLE IF EXISTS `utilisateur`;
CREATE TABLE IF NOT EXISTS `utilisateur` (
  `Id_Utilisateur` int NOT NULL,
  `Pseudo` varchar(50) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `Grade` int DEFAULT NULL,
  `Adresse_e_mail` varchar(50) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `Mot_de_passe` varchar(50) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  PRIMARY KEY (`Id_Utilisateur`),
  KEY `Grade` (`Grade`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `message`
--
ALTER TABLE `message`
  ADD CONSTRAINT `message_ibfk_1` FOREIGN KEY (`Id_Reponse`) REFERENCES `message` (`Id_Message`);

--
-- Contraintes pour la table `utilisateur`
--
ALTER TABLE `utilisateur`
  ADD CONSTRAINT `utilisateur_ibfk_1` FOREIGN KEY (`Grade`) REFERENCES `grade` (`Id_Grade`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
