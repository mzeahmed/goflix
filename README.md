# Goflix

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)](https://github.com/mzeahmed/goflix/blob/main/app/go.mod)
[![Gorilla Mux](https://img.shields.io/badge/Gorilla-Mux-6E4A7E?logo=go&logoColor=white)](https://github.com/gorilla/mux)
[![MariaDB](https://img.shields.io/badge/MariaDB-11-003545?logo=mariadb&logoColor=white)](https://mariadb.org)
[![JWT](https://img.shields.io/badge/Auth-JWT-000000?logo=jsonwebtokens&logoColor=white)](https://jwt.io)

API REST en Go permettant de gérer un catalogue de films (CRUD basique) avec authentification par JWT.

## Prérequis

- Go 1.21+
- Docker & Docker Compose (pour Traefik, MariaDB, phpMyAdmin, Mailpit)
- [mkcert](https://github.com/FiloSottile/mkcert) (pour générer les certificats TLS locaux)

## Démarrage rapide (Docker)

```bash
make hosts   # ajoute api.goflix.local, mail.goflix.local, db.goflix.local à /etc/hosts (sudo)
make up      # génère les certificats si absents, build et démarre les conteneurs
```

`make up` démarre Traefik, l'application Go, MariaDB, phpMyAdmin et Mailpit, puis affiche les URLs
disponibles :

- Application : `https://api.goflix.local`
- Dashboard Traefik : `http://localhost:8080`
- phpMyAdmin : `https://db.goflix.local`
- Mailpit : `https://mail.goflix.local`

## Lancer le serveur en local (hors Docker)

```bash
docker compose up -d database
make run
```

Le serveur démarre sur le port **9000** :

```
Goflix
Serving HTTP on port 9000
```

L'API est alors accessible sur `http://localhost:9000`.

## Commandes Makefile

```bash
make help
```

| Commande      | Description                                          |
|---------------|-------------------------------------------------------|
| `make run`    | Lance le serveur en local                             |
| `make build`  | Compile le binaire dans `app/bin/goflix`               |
| `make fmt`    | Formate le code source                                |
| `make vet`    | Lance `go vet`                                        |
| `make test`   | Lance les tests unitaires                             |
| `make check`  | Lance `fmt`, `vet` et `test`                          |
| `make tidy`   | Nettoie `go.mod` / `go.sum`                           |
| `make update` | Met à jour les dépendances                            |
| `make hosts`  | Ajoute les domaines locaux à `/etc/hosts` (sudo)      |
| `make certs`  | Génère les certificats TLS locaux si absents          |
| `make up`     | Build et démarre les conteneurs Docker                |
| `make down`   | Arrête les conteneurs Docker                          |
| `make restart`| Redémarre les conteneurs Docker                       |
| `make logs`   | Affiche les logs des conteneurs                       |
| `make ps`     | Liste les conteneurs                                  |
| `make bash`   | Ouvre un shell dans le conteneur de l'application      |
| `make clean`  | Supprime les fichiers générés                         |
| `make doctor` | Affiche l'environnement de dev                        |

## Structure du projet

```
app/cmd/goflix/       # Point d'entrée (main)
app/internal/server/  # Routage HTTP, handlers, middlewares
app/internal/store/   # Modèles et accès à la base de données
docker/app/           # Dockerfile de l'application
traefik/              # Configuration Traefik (statique + dynamique)
```

## Tester l'API avec curl

Les exemples ci-dessous utilisent `https://api.goflix.local` (stack Docker complète via Traefik).
En local sans Docker, remplacer par `http://localhost:9000`.

### 1. Obtenir un token JWT

`POST /api/token`

Identifiants valides : `golang` / `rocks`.

```bash
curl -X POST https://api.goflix.local/api/token \
  -H "Content-Type: application/json" \
  -d '{"username":"golang","password":"rocks"}'
```

Réponse :

```json
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."}
```

### 2. Lister les films

`GET /api/movies/`

```bash
curl https://api.goflix.local/api/movies/
```

Réponse :

```json
[{"id":1,"title":"Inception","release_date":"2010-07-16","duration":148,"trailer_url":"https://youtube.com/watch?v=xxx"}]
```

### 3. Détail d'un film

`GET /api/movies/{id}`

```bash
curl https://api.goflix.local/api/movies/1
```

Réponse :

```json
{"id":1,"title":"Inception","release_date":"2010-07-16","duration":148,"trailer_url":"https://youtube.com/watch?v=xxx"}
```

### 4. Créer un film

`POST /api/movies/`

```bash
curl -X POST https://api.goflix.local/api/movies/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Inception",
    "release_date": "2010-07-16",
    "duration": 148,
    "trailer_url": "https://youtube.com/watch?v=xxx"
  }'
```

Réponse :

```json
{"id":1,"title":"Inception","release_date":"2010-07-16","duration":148,"trailer_url":"https://youtube.com/watch?v=xxx"}
```

## Routes disponibles

| Méthode | Route                    | Description                    |
|---------|--------------------------|---------------------------------|
| POST    | `/api/token`             | Génère un token JWT             |
| GET     | `/api/movies/`           | Liste les films                 |
| GET     | `/api/movies/{id}`       | Détail d'un film                |
| POST    | `/api/movies/`           | Crée un film                    |