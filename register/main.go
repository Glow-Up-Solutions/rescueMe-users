package main

import (
	"encoding/json"
	"fmt"
	db "register/connector"
	"register/model"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
)

type Response struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func verifyData(u *model.Usuarios) (string, error) {
	var msg string
	var err error
	if u.First_name == "" {
		msg = "| Ingresa un nombre"
	}
	if u.Last_name == "" {
		msg = msg + "| Ingresa un apellido"
	}
	if u.Birthday == "" {
		msg = msg + "| Ingresa una fecha de cumplenos"
	}
	if u.Username == "" {
		msg = msg + "| Ingresa un username"
	}
	if u.Password == "" {
		msg = msg + "| Ingresa un password"
	}
	if u.Email == "" {
		msg = msg + "| Ingresa un correo electronico"
	}
	if u.City == "" {
		msg = msg + "| Ingresa una ciudad"
	}
	if u.Code_zip == "" {
		msg = msg + "| Ingresa un codigo postal"
	}
	if u.State == "" {
		msg = msg + "| Ingresa un Estado"
	}
	if msg != "" {
		err = fmt.Errorf("BAD REQUEST BRUO")
	} else {
		err = nil
	}
	return msg, err
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var u model.Usuarios
	err := json.Unmarshal([]byte(request.Body), &u)
	if err != nil {
		brq, _ := json.Marshal("Send correct data")
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body:       string(brq),
			StatusCode: 400,
		}, err
	}
	//verify data not empty
	msg, status := verifyData(&u)
	if status != nil {
		brq, _ := json.Marshal(msg)
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body:       string(brq),
			StatusCode: 400,
		}, err
	}
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	defer db2.Close()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, err
	}
	db2.Exec("INSERT INTO users (first_name, last_name, birthday, username, password, email, city, code_zip, state) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);",
		u.First_name, u.Last_name, u.Birthday, u.Username, u.Password, u.Email, u.City, u.Code_zip, u.State)
	fmt.Println("User created successfully")

	newMes := Response{
		Message:  "User successfully created",
		Username: u.Username,
	}
	a, _ := json.Marshal(newMes)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:       string(a),
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(handler)
}
