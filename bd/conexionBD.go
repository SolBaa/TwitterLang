package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCT >mongo ct conecta la base de datps
var MongoCT = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://solbaa:ginofionaluna@cluster0.kxgcm.mongodb.net/test")

// ConectarDB es la funcion para conectar la base de datos
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con la BD")
	return client
}

// CheckConection es la funcion que permite que checquear la conexion
func CheckConection() int {
	err := MongoCT.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1

}
