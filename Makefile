# ==============================================================================
# Goflix - Makefile de développement
# ==============================================================================

.DEFAULT_GOAL := help

COMPOSE := docker compose -f docker-compose.yml

APP_CONTAINER := goflix_app

DOMAINS := api.goflix.local mail.goflix.local db.goflix.local

GREEN  := \033[0;32m
YELLOW := \033[1;33m
BLUE   := \033[0;34m
RED    := \033[0;31m
RESET  := \033[0m

CERT_FILE := certs/goflix.local+1.pem
CERT_KEY  := certs/goflix.local+1-key.pem

.PHONY: help run build \
        fmt vet test check \
        tidy update \
        clean doctor \
        hosts certs up down restart logs ps bash

help: ## Affiche les commandes disponibles
	@echo ""
	@echo "$(BLUE)Goflix Development Commands$(RESET)"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_-]+:.*##/ {printf "  \033[32m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""

# ==============================================================================
# Développement
# ==============================================================================

run: ## Lance le serveur
	cd app && go run ./cmd/goflix

build: ## Compile le binaire local
	@mkdir -p app/bin
	cd app && go build -o bin/goflix ./cmd/goflix
	@echo "$(GREEN)✓ Binaire généré dans app/bin/goflix$(RESET)"

# ==============================================================================
# Qualité
# ==============================================================================

fmt: ## Formate le code source
	cd app && go fmt ./...

vet: ## Lance go vet
	cd app && go vet ./...

test: ## Lance les tests unitaires
	cd app && go test ./...

check: fmt vet test ## Lance toutes les vérifications qualité

# ==============================================================================
# Dépendances
# ==============================================================================

tidy: ## Nettoie go.mod / go.sum
	cd app && go mod tidy

update: ## Met à jour les dépendances
	cd app && go get -u ./...
	cd app && go mod tidy

# ==============================================================================
# Docker
# ==============================================================================

hosts: ## Ajoute les domaines locaux à /etc/hosts (nécessite sudo)
	@echo "$(YELLOW)Mise à jour de /etc/hosts...$(RESET)"
	@for domain in $(DOMAINS); do \
		if grep -qE "^127\.0\.0\.1[[:space:]]+$$domain$$" /etc/hosts; then \
			echo "$(GREEN)$$domain déjà présent$(RESET)"; \
		else \
			echo "127.0.0.1 $$domain" | sudo tee -a /etc/hosts > /dev/null; \
			echo "$(GREEN)$$domain ajouté$(RESET)"; \
		fi; \
	done

certs: ## Génère les certificats TLS locaux si absents (nécessite mkcert)
	@if [ -f $(CERT_FILE) ] && [ -f $(CERT_KEY) ]; then \
		echo "$(GREEN)Certificats déjà présents$(RESET)"; \
	else \
		echo "$(YELLOW)Génération des certificats...$(RESET)"; \
		mkcert -install; \
		mkcert -cert-file $(CERT_FILE) -key-file $(CERT_KEY) goflix.local "*.goflix.local"; \
		echo "$(GREEN)Certificats générés dans certs/$(RESET)"; \
	fi

up: certs ## Build et démarre les conteneurs
	@echo "$(YELLOW)Démarrage des conteneurs...$(RESET)"
	$(COMPOSE) up -d --build
	@echo "$(GREEN)Conteneurs démarrés$(RESET)"
	@echo "$(BLUE)Dashboard Traefik : http://localhost:8080$(RESET)"
	@echo "$(BLUE)URL api : https://api.goflix.local$(RESET)"
	@echo "$(BLUE)URL Mailpit : https://mail.goflix.local$(RESET)"
	@echo "$(BLUE)URL PhpMyAdmin : https://db.goflix.local$(RESET)"

down: ## Arrête les conteneurs
	@echo "$(YELLOW)Arrêt des conteneurs...$(RESET)"
	$(COMPOSE) down
	@echo "$(GREEN)Conteneurs arrêtés$(RESET)"

restart: down up ## Redémarre les conteneurs

logs: ## Affiche les logs des conteneurs
	@echo "$(YELLOW)Affichage des logs...$(RESET)"
	$(COMPOSE) logs -f

ps: ## Liste les conteneurs
	@echo "$(YELLOW)Liste des conteneurs...$(RESET)"
	$(COMPOSE) ps

bash: ## Accède au conteneur app
	@echo "$(YELLOW)Accès au conteneur app...$(RESET)"
	docker exec -it $(APP_CONTAINER) sh

# ==============================================================================
# Utilitaires
# ==============================================================================

clean: ## Supprime les fichiers générés
	rm -rf app/bin

doctor: ## Affiche l'environnement de développement
	@echo ""
	@echo "$(BLUE)Environment$(RESET)"
	@echo ""
	@go version
	@git --version