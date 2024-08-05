package main

import (
	"HT1_MIAB_2S2024/structures"
	"bufio"
	"errors"
	"fmt"
	"os"
)

const fullPath = "/home/jonathan/MIAB_2S/HT1/HT1_MIAB_2S2024/dataSheet.mia"

type Add struct {
	tipo   int64
	Carnet string
	Nombre string
	CUI    string
	Id     string
}

func main() {
	createDataSheet(1, "K")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1)Ingresar estudiante.")
		fmt.Println("2)Ingresar profesor.")
		fmt.Println("3)Ver Registros.")
		fmt.Println("4)Salir.")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		switch input {
		case "1":
			addStudent(scanner)

		case "4":
			return
		}

	}
}

func addStudent(scanner *bufio.Scanner) {
	toAdd := &Add{}
	toAdd.tipo = 1
	fmt.Println("Ingrese el id del estudiante:")
	if scanner.Scan() {
		toAdd.Id = scanner.Text()
	}

	fmt.Println("Ingrese el CUI del estudiante:")
	if scanner.Scan() {
		toAdd.CUI = scanner.Text()
	}

	fmt.Println("Ingrese el Nombre del estudiante:")
	if scanner.Scan() {
		toAdd.Nombre = scanner.Text()
	}

	fmt.Println("Ingrese el carne del estudiante:")
	if scanner.Scan() {
		toAdd.Carnet = scanner.Text()
	}

	fmt.Printf("Estudiante ingresado: %+v\n", toAdd)
	stuToAdd := &structures.Student{}
	copy(stuToAdd.Id_Estu[:], toAdd.Id)
	copy(stuToAdd.CUI[:], toAdd.CUI)
	copy(stuToAdd.Nombre[:], toAdd.Nombre)
	copy(stuToAdd.Carnet[:], toAdd.Carnet)
	stuToAdd.WriteToFile(fullPath)

}

func addTeacher(scanner *bufio.Scanner) {
	toAdd := &Add{}
	toAdd.tipo = 0
	fmt.Println("Ingrese el id del profesor:")
	if scanner.Scan() {
		toAdd.Id = scanner.Text()
	}

	fmt.Println("Ingrese el CUI del profesor:")
	if scanner.Scan() {
		toAdd.CUI = scanner.Text()
	}

	fmt.Println("Ingrese el Nombre del profesor:")
	if scanner.Scan() {
		toAdd.Nombre = scanner.Text()
	}

	fmt.Println("Ingrese el curso del profesor:")
	if scanner.Scan() {
		toAdd.Carnet = scanner.Text()
	}

	fmt.Printf("Estudiante ingresado: %+v\n", toAdd)
}

func writeFile(file *os.File, sizeInBytes int) error {
	buffer := make([]byte, 1024*1024)
	for sizeInBytes > 0 {
		writeSize := len(buffer)
		if sizeInBytes < writeSize {
			writeSize = sizeInBytes
		}
		if _, err := file.Write(buffer[:writeSize]); err != nil {
			return err
		}
		sizeInBytes -= writeSize
	}
	fmt.Println("Archivo creado con exito")
	return nil
}

func convertToBytes(size int, unit string) (int, error) {
	switch unit {
	case "K":
		return size * 1024, nil // Convierte kilobytes a bytes
	case "M":
		return size * 1024 * 1024, nil // Convierte megabytes a bytes
	default:
		return 0, errors.New("unidad inválida") // Devuelve un error si la unidad es inválida
	}
}

// /home/jonathan/MIAB_2S/HT1/HT1_MIAB_2S2024/structures
func createDataSheet(size int, unit string) error {
	sizeInbytes, err := convertToBytes(size, unit)
	if err != nil {
		return err
	}

	//crear dataSheet
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("error el crear la datasheet")
		return err
	}

	defer file.Close()

	return writeFile(file, sizeInbytes)

}
