start:
	set -o allexport; source .env; set +o allexport && go run ./cmd/srv/main.go