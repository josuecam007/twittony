package bd

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName = os.Getenv("dbNametw")
var userDb = os.Getenv("userDbtw")
var passDb = os.Getenv("passDbtw")

var url = "mongodb+srv://" + userDb + ":" + passDb + "@cluster0-utqfa.mongodb.net/" + dbName + "?retryWrites=true&w=majority"

/*MongoCN es el objeto de conexion ala base de datos*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI(url)

/*ConectarBD es la funcion que conecta la base de datos*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Conexion exitosa con la BD")
	return client
}

/*ChequeoConexion es el ping a la base de datos*/
func ChequeoConexion() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}

	return true
}
