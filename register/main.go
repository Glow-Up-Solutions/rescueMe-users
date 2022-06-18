package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
)

type Usuarios struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Birthday   string `json:"birthday"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	City       string `json:"city"`
	Code_zip   string `json:"code_zip"`
	State      string `json:"state"`
}

type Response struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

const (
	// host     = "database-rescueme.caizuq7lxwzo.us-east-1.rds.amazonaws.com"
	// port     = 5431
	// password = "pws$69%2022LaMaou"
	// dbname = "rescuemedb"
	// user     = "masteruser"
	host     = "localhost"
	dbname   = "postgres"
	port     = 54321
	user     = "postgres"
	password = ""
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var u Usuarios
	_ = json.Unmarshal([]byte(request.Body), &u)

	// val := Usuarios{
	// 	First_name: "JOathan",
	// 	Last_name:  "espinsoa",
	// 	Birthday:   "17 marzo",
	// 	Username:   "asfafASDF",
	// 	Password:   "1235413FG",
	// 	Email:      "jonathan@hotmail.com",
	// 	City:       "zitacuaro",
	// 	Code_zip:   "61518",
	// 	State:      "Michoacn",
	// }

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}
	db.Exec(
		"INSERT INTO users (first_name, last_name, birthday, username, password, email, city, code_zip, state) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);",
		u.First_name, u.Last_name, u.Birthday, u.Username, u.Password, u.Email, u.City, u.Code_zip, u.State)
	fmt.Println("User created successfully")

	// rows, _ := db.Query("SELECT * FROM users;")
	// value := Usuarios{}
	// for rows.Next() {
	// 	err = rows.Scan(&value.First_name, &val.Last_name, &value.State)
	// 	fmt.Println(value.First_name, value.Last_name, value.State)
	// }
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
