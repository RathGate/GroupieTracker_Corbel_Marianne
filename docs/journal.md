# Journal - Groupie Tracker

Journal d’organisation et de suivi du projet Groupie Tracker, mené par Marianne Corbel (B1 Info).

## Lundi 06/03

<aside>
📒 **Projet lancé à 13h45 !**

</aside>

- **Après-midi (13h45 - 17h45) :**
    - Choix de l’API ([Hyrule Compendium API](https://github.com/Azure/azure-content/blob/master/contributor-guide/contributor-guide-index.md))
    - Familiarisation avec l’API, premières requêtes de test
    - Création du repo GitHub, arborescence de base d’un projet Go.

## Mardi 07/03

- **Matin (8h30 - 12h30) :**
    
    ![all.png](Journal%20-%20Groupie%20Tracker%20b01bf5e0725f468ca25ccfb8108ccca3/all.png)
    
    - Création des cartes pour les items renvoyés par l’API.
    - Traitement d’une requête basique en go (structs, fonctions d’appels à l’API)
    - Affichage des résultats de la requête (endpoint `/all`) avec les templates go.

- **Après-midi (13h45 - 17h45) :**
    
    ![Untitled](Journal%20-%20Groupie%20Tracker%20b01bf5e0725f468ca25ccfb8108ccca3/Untitled.png)
    
    - Puisque toutes les requêtes ne renvoient pas la même structure `json`, création des fonctions permettant “d’aplatir” les différents niveaux.
    - Création d’un système de recherche par nom
    - Création d’un fichier `.json` au lancement du serveur (fallback de l’endpoint `/all` pour réduire le temps de traitement des recherches)
    
- **Soir (20h30 - 21h30) :**
    
    ![Untitled](Journal%20-%20Groupie%20Tracker%20b01bf5e0725f468ca25ccfb8108ccca3/Untitled%201.png)
    
    - Différentiation: “perfect match” si un item porte exactement le nom entré par l’utilisateur, “all results” pour les autres.
    

## Jeudi 09/03

- **Après-midi (16h30 - 19h30 // 22h30 - 00h30) :**
    
    ![Untitled](Journal%20-%20Groupie%20Tracker%20b01bf5e0725f468ca25ccfb8108ccca3/Untitled%202.png)
    
    - Création de la navbar, du logo et du favicon
    - Tentative de design proche du compendium in-game de Zelda Breath of the Wild pour la page des catégories
    - Pas encore de responsive ni de gestion de l’affichage par catégorie.
    
    ## Samedi 11/03
    
    - **Après-midi (16h30 - 19h30 // 22h30 - 00h30) :**
        
        
        ![Untitled](Journal%20-%20Groupie%20Tracker%20b01bf5e0725f468ca25ccfb8108ccca3/Untitled%203.png)
        
        ![Screenshot_20230311-215545.jpg](Journal%20-%20Groupie%20Tracker%20b01bf5e0725f468ca25ccfb8108ccca3/Screenshot_20230311-215545.jpg)
        
        - Design de la navbar pour petits écrans (petit logo, burger menu), entièrement responsive.
        - Footer entièrement fait (responsive inclus).
        - Refonte de l’arborescence des fichiers, premier rangement du CSS