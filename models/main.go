package models

type Monster struct {
  Id int `db:"id" json:"id"`
  Name string `db:"name" json:"name"`
  Size string `db:"size" json:"size"`
  Type string `db:"type" json:"type"`
  Alignment string `db:"alignment" json:"alignment"`
  ArmorClass int `db:"armor_class" json:"armor_class"`
  ChallengeRating float32 `db:"challenge_rating" json:"challenge_rating"`
  ExperiencePoints int `db:"experience_points" json:"experience_points"`
}

