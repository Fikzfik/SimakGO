package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mahasiswa struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NIM       string             `bson:"nim" json:"nim"`
	Name      string             `bson:"name" json:"name"`
	Prodi     string             `bson:"prodi" json:"prodi"`
	Angkatan  int                `bson:"angkatan" json:"angkatan"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"` // Link to User
	CreatedAt int64              `bson:"created_at" json:"created_at"`
}
