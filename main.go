package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func conexionBaseDatos() *sql.DB{
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error al cargar el archivo .env",err)
	}

	usuario := os.Getenv("USUARIO")
	contrasena := os.Getenv("CONTRASENA")
	direccion := os.Getenv("DIRECCION")
	puerto := os.Getenv("PUERTO")
	nombredatabase := os.Getenv("NOMBREDATABASE")

	dns := fmt.Sprintf("%v:%v@tcp(%v:%s)/%v", usuario, contrasena, direccion, puerto, nombredatabase)

	db, err := sql.Open("mysql", dns)
	if err != nil{
		log.Fatal("Error al abrir la conexion", err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatal("No se pudo conectar a la base de datos", err)
	}

	fmt.Println("Conexion exitosa a la base de datos")
	return db
}

func main() {
	// CONEXION BASE DE DATOS
	db := conexionBaseDatos()

	// REGISTRAR UNA CLIENTE
	fmt.Println("Nuevo cliente")

	var nombre string
	var telefono string
	var email string

	fmt.Println("Nombre")
	fmt.Scanln(&nombre)
	fmt.Println("telefono")
	fmt.Scanln(&telefono)
	fmt.Println("email")
	fmt.Scanln(&email)
	
	query := "INSERT INTO clientes (nombre, telefono, email) VALUES(?,?,?)"
	_, err := db.Exec(query, nombre,telefono,email)
	if err != nil{
		log.Fatal("Error al insertar el cliente", err)
	}else{
		fmt.Println("Cliente insertado correctamente")
	}
}