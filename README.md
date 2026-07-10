# Goflix

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go&logoColor=white)](https://github.com/mzeahmed/goflix/blob/main/go.mod)
[![Gorilla Mux](https://img.shields.io/badge/Gorilla-Mux-6E4A7E?logo=go&logoColor=white)](https://github.com/gorilla/mux)
[![SQLite](https://img.shields.io/badge/SQLite-3-003B57?logo=sqlite&logoColor=white)](https://www.sqlite.org)
[![JWT](https://img.shields.io/badge/Auth-JWT-000000?logo=jsonwebtokens&logoColor=white)](https://jwt.io)

API REST en Go permettant de gérer un catalogue de films (CRUD basique) avec authentification par JWT.

## Prérequis

- Go 1.18+
- SQLite (via `github.com/mattn/go-sqlite3`)

## Lancer le serveur

```bash
go run .
```

Le serveur démarre sur le port **9000** :

```
Goflix
Serving HTTP on port 9000
```

L'API est alors accessible sur `http://localhost:9000`.

## Tester l'API avec curl

### 1. Obtenir un token JWT

`POST /api/token`

Identifiants valides : `golang` / `rocks`.

```bash
curl -X POST http://localhost:9000/api/token \
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
curl http://localhost:9000/api/movies/
```

### 3. Détail d'un film

`GET /api/movies/{id}`

```bash
curl http://localhost:9000/api/movies/1
```

### 4. Créer un film

`POST /api/movies/`

```bash
curl -X POST http://localhost:9000/api/movies/ \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Inception",
    "release_date": "2010-07-16",
    "duration": 148,
    "trailer_url": "https://youtube.com/watch?v=xxx"
  }'
```

## Routes disponibles

| Méthode | Route                    | Description                  |
|---------|---------------------------|-------------------------------|
| POST    | `/api/token`              | Génère un token JWT           |
| GET     | `/api/movies/`            | Liste les films               |
| GET     | `/api/movies/{id}`        | Détail d'un film               |
| POST    | `/api/movies/`            | Crée un film                  |
