package db

import (
  "bufio"
  "encoding/csv"
  "fmt"
  "io"
  "log"
  "os"
  _ "github.com/lib/pq"
  "github.com/jmoiron/sqlx"
)

var createTablesSchema = `
CREATE TABLE monsters (
  id serial primary key,
  name text,
  size text,
  type text,
  alignment text,
  armor_class integer,
  challenge_rating decimal,
  experience_points integer
);
`

var db *sqlx.DB

func prepopulateDbFromFilePath(filePath string) {
  fmt.Println("Prepopulating database with monsters dictionary...")
  defer fmt.Println("Completed prepopulating database")

  csvFile, _ := os.Open(filePath)
  reader := csv.NewReader(bufio.NewReader(csvFile))

  // skip first line
  reader.Read()

  // read CSV files line by line and insert monster data into Postgres
  tx := db.MustBegin()
  for {
    line, err := reader.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }

    tx.MustExec("INSERT INTO monsters (name,size,type,alignment,armor_class,challenge_rating,experience_points) VALUES ($1,$2,$3,$4,$5,$6,$7)", line[0], line[1], line[2], line[3], line[4], line[5], line[6])
  }

  err := tx.Commit()
  if err != nil {
    log.Fatal(err)
  }
}

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

    // pre-populate monster database from monsters CSV file
    prepopulateDbFromFilePath(os.Getenv("MONSTERLIB_PATH"))
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

