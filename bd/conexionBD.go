package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion ala base de datos*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://adminDbtw007:<password>@cluster0-utqfa.mongodb.net/<dbname>?retryWrites=true&w=majority")

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
