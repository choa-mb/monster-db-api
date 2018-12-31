export PG_HOST=localhost
export PG_PORT=5432
export PG_USER=postgres
export PG_PASSWORD=postgres
export PG_DB=monsters
export POSTGRES_PASSWORD=postgres
export POSTGRES_USER=postgres
export POSTGRES_DB=monsters
export MONSTERLIB_PATH=./monsterlib/monsters.txt

dev-db:
	docker-compose up -d postgres  

dev:
	watcher

