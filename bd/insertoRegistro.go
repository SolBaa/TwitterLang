package bd

import (
	"context"
	"time"

	"github.com/SolBaa/TwitterLang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro Registrar e  un usuario
func InsertoRegistro(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCT.Database("test")
	col := db.Collection("users")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
