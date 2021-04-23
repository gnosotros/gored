package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()

var clienteOptions = options.Client().ApplyURI("mongodb+srv://yoa:bAsededatosmong0@cluster0.gw0xs.mongodb.net/twittor?retryWrites=true&w=majority")

//var clientOptions = options.Client().ApplyURI("mongodb+srv://root:EntrenamientoTwitter@twitter-kb2wp.mongodb.net/test?retryWrites=true&w=majority")

//conectar Bd es la funcion que me permite conectar la base de datos
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil) //muestra si hay un error en la base de datos
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("conexion Exitosa con la BD")
	return client
}

//ChequeoConnection() es el ping a la base de datos
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil) //se pone mongoCN que es la variable que debo poner
	if err != nil {
		return 0
	}
	return 1

}
