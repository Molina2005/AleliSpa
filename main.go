package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func FechaDDMMYYYY(fecha string)(time.Time, error){
	formatoFecha := "01-02-06"
	return time.Parse(formatoFecha, fecha)
}

func ConexionBaseDatos() *sql.DB{
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

func RegistrarCliente(db *sql.DB){
	fmt.Println("Nuevo cliente")
	var nombre,apellido,telefono,email,direccion,fechaRegistro string
	
	fmt.Println("Nombre")
	fmt.Scanln(&nombre)
	fmt.Println("Apellido")
	fmt.Scanln(&apellido)
	fmt.Println("Telefono")
	fmt.Scanln(&telefono)
	fmt.Println("Email")
	fmt.Scanln(&email)
	fmt.Println("Direccion")
	fmt.Scanln(&direccion)
	fmt.Println("Fecha registro")
	fmt.Scanln(&fechaRegistro)

	fecha, err := FechaDDMMYYYY(fechaRegistro)
	if err != nil {
		log.Fatal("error al formatear la fecha", err)
	}

	query := "INSERT INTO CLIENTES (NOMBRE, APELLIDO, TELEFONO, EMAIL, DIRECCION, FECHA_REGISTRO) VALUES(?,?,?,?,?,?)"
	_, err = db.Exec(query, nombre,apellido, telefono, email, direccion, fecha)
	if err != nil{
		log.Fatal("Error al insertar el cliente", err)
	}else{
		fmt.Println("Cliente insertado correctamente")
	}
}

func AgendarCita(db *sql.DB){
	fmt.Println("AGENDACION DE CITA")
	var fechaCita, horaCita, procedimiento, observaciones string

	fmt.Println("Inserte la fecha en que desea la cita, Ejemplo: 10-08-25 (DD-MM-YYY)")
	fmt.Scanln(&fechaCita)
	fmt.Println("Inserte la hora en que desea la cita, Ejmeplo: 14:00 (Hora militar)")
	fmt.Scanln(&horaCita)
	fmt.Println("Inserte el procedimiento que desea")
	fmt.Scanln(&procedimiento)
	fmt.Println("Inserte alguna observacion a tener en cuenta, en caso de no tener digite N/A")
	fmt.Scanln(&observaciones)

	fecha, err := FechaDDMMYYYY(fechaCita)
	if err != nil {
		log.Fatal("Error al formatear la fecha", err)
	}

	query := "INSERT INTO CITAS (FECHA_CITA, HORA_CITA, PROCEDIMIENTO, OBSERVACIONES) VALUES (?,?,?,?)"
	_, err = db.Exec(query, fecha, horaCita, procedimiento, observaciones)
	if err != nil {
		log.Fatal("Error al insertar la cita", err)
	}else {
		fmt.Println("cita insertada correctamente")
	}
}

func main() {
	db := ConexionBaseDatos()
	RegistrarCliente(db)
	AgendarCita(db)
}