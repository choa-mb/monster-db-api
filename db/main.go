package db

import (
  "fmt"
  "log"
  "os"
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
)

var createTablesSchema = `
CREATE TABLE monsters (
  name text
);

COPY monsters (name) FROM '/var/monsterlib/monsters.txt' DELIMITER ',' CSV HEADER;
`

var db *sqlx.DB

func Init() {
  fmt.Println("Initializing connection to database...")

  var err error
  db, err = sqlx.Open("postgres", GetDbEnvVariables())
  if err != nil {
    log.Fatalln(err)
  }

  // check if table already exists before creating table
  _, err = db.Query(fmt.Sprintf("SELECT 1 FROM %s LIMIT 1", os.Getenv("PG_DB")))
  if err != nil {
    db.MustExec(createTablesSchema)
  }
}

func GetDB() *sqlx.DB {
  return db
}

func GetDbEnvVariables() string {
  return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    os.Getenv("PG_HOST"),
    os.Getenv("PG_PORT"),
    os.Getenv("PG_USER"),
    os.Getenv("PG_PASSWORD"),
    os.Getenv("PG_DB"),
  )
}

