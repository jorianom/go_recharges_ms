package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorianom/go-recharges-ms/routes"
	"github.com/streadway/amqp"
)

func main() {
	// Define los parámetros de conexión a RabbitMQ
	rabbitMQURL := "amqp://guest:guest@localhost:5672/"
	queueName := "recharges2"

	// Establece una conexión a RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatal("Error al conectar a RabbitMQ:", err)
	}
	defer conn.Close()

	fmt.Printf("Mensaje recibido:")
	// Crea un canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error al abrir un canal:", err)
	}
	defer ch.Close()

	// Consume mensajes de la cola
	msgs, err := ch.Consume(
		queueName, // Nombre de la cola
		"",        // Etiqueta del consumidor (dejar en blanco para una etiqueta generada automáticamente)
		true,      // AutoAck (true para confirmar automáticamente los mensajes)
		false,     // Exclusive (true para que solo esta conexión pueda acceder a la cola)
		false,     // NoLocal (true para no recibir mensajes publicados por esta conexión)
		false,     // NoWait (true para no esperar una respuesta de confirmación)
		nil,       // Argumentos adicionales
	)
	if err != nil {
		log.Fatal("Error al consumir mensajes:", err)
	}

	// Escucha los mensajes
	for msg := range msgs {
		fmt.Printf("Mensaje recibido: %s\n", msg.Body)
	}
	// client := driver.Connection()
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	router := mux.NewRouter()
	//	s := r.PathPrefix("/api").Subrouter()
	router.HandleFunc("/api/recharge", routes.RechargeHandler).Methods("POST")
	router.HandleFunc("/api/recharges/{id}", routes.HistoryHandler).Methods("GET")
	//
	router.HandleFunc("/api/method", routes.PostMethodHandler).Methods("POST")
	router.HandleFunc("/api/methods/{id}", routes.GetMethodHandler).Methods("GET")
	router.HandleFunc("/api/method/{id}", routes.UpdateMethodHandler).Methods("PUT")
	router.HandleFunc("/api/method/{id}", routes.DeleteMethodHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
