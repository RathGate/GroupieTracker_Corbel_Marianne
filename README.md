# COMPENDIUM DATA VIEWER

The goal of the exercice is to realise a website displaying data received from an external API, and to be able to make different requests to that API from the website itself.

## About
The Compendium Data Viewer uses [Aarav Borthakur](https://github.com/gadhagod)'s [Hyrule Compendium API](https://gadhagod.github.io/Hyrule-Compendium-API/#/), which provides the data of all the entries from the bestiary of the game The Legend of Zelda: Breath of the Wild.

The website has two major features: displaying all the entries just like the game goes (by category), and search for specific entries with more precise filters like name, item type, location and gamemode.

![Page /search from the Compendium Data Viewer](https://media.discordapp.net/attachments/1001959681004163103/1086648904340222003/image.png)

## Technical Specifications

-   Back-end: Golang
-   Front-end: HTML, CSS, JS + JQuery
-   API used:  [Hyrule Compendium API](https://gadhagod.github.io/Hyrule-Compendium-API/#/)

By default and to improve reactivity, the program stores a copy of the processed API data on server launch, notably used for search queries. It is possible to disable that feature by changing the boolean value of the variable `USEFALLBACK` at the top of the `main.go` file.

**COMPATIBILITY:** The website has been entirely tested on the latest versions of Chrome, Firefox and Edge. It should also appear responsive on Chrome, Safari and Firefox mobile browsers !

## How to use the program

As this project includes a server, it is not hosted (at least yet). In order to use it, you must clone the repo with

    git clone https://github.com/RathGate/Hyrule-Compendium-Data-Viewer

Then, you must open a terminal in the **src/** folder, at the root of the project, and use the following command to launch the server :

    go run .

The website should now be available on `localhost:8080` . If not, or in case of a port collision, feel free to change the `preferredPort` variable, at the end of the `main.go` file, with a port that is more convenient to you.

Enjoy ! ♫
