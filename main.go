package main

import (
	"HT1_MIAB_2S2024/structures"
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strings"
)

const fullPath = "/home/jonathan/MIAB_2S/HT1/HT1_MIAB_2S2024/dataSheet.mia"

type Add struct {
	tipo   string
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
		case "2":
			addTeacher(scanner)
		case "3":
			readDataSheet1()
			//readDataSheet()
		case "4":
			return
		}

	}
}

func addStudent(scanner *bufio.Scanner) {
	toAdd := &Add{}
	toAdd.tipo = "1"
	fmt.Println("Ingrese el id del estudiante:")
	for {
		if scanner.Scan() {
			//Validar ID unico
			if !validateId(scanner.Text()) {
				fmt.Println("Ingrese nuevamente el id del estudiante:")
			} else {
				toAdd.Id = scanner.Text()
				break
			}
		}
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
	copy(stuToAdd.Tipo[:], toAdd.tipo)
	copy(stuToAdd.Id_Estu[:], toAdd.Id)
	copy(stuToAdd.CUI[:], toAdd.CUI)
	copy(stuToAdd.Nombre[:], toAdd.Nombre)
	copy(stuToAdd.Carnet[:], toAdd.Carnet)
	stuToAdd.WriteToFile(fullPath)
	//fmt.Println("Todo bien hasta aca")

}

func addTeacher(scanner *bufio.Scanner) {
	toAdd := &Add{}
	toAdd.tipo = "2"
	fmt.Println("Ingrese el id del profesor:")
	for {
		if scanner.Scan() {
			//Validar ID unico
			if !validateId(scanner.Text()) {
				fmt.Println("Ingrese nuevamente el id del profesor:")
			} else {
				toAdd.Id = scanner.Text()
				break
			}
		}
	}

	//if scanner.Scan() {
	//	toAdd.Id = scanner.Text()
	//}

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

	fmt.Printf("Profesor ingresado: %+v\n", toAdd)
	teachToAdd := &structures.Teacher{}
	copy(teachToAdd.Tipo[:], toAdd.tipo)
	copy(teachToAdd.Id_profesor[:], toAdd.Id)
	copy(teachToAdd.CUI[:], toAdd.CUI)
	copy(teachToAdd.Nombre[:], toAdd.Nombre)
	copy(teachToAdd.Curso[:], toAdd.Carnet)
	teachToAdd.WriteToFile(fullPath)
	//fmt.Println("Todo bien hasta aca")
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

func readDataSheet1() {
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("error al abrir la datasheet")
		return
	}
	defer file.Close()

	offset := int64(0)
	stuSize := binary.Size(structures.Student{})
	teachSize := binary.Size(structures.Teacher{})
	var students []structures.Student
	var teachers []structures.Teacher

	for {
		_, err = file.Seek(offset, 0)
		if err != nil {
			fmt.Println("error al buscar el offset")
			return
		}

		// Leer el primer byte para determinar el tipo
		var tipo byte
		err = binary.Read(file, binary.LittleEndian, &tipo)
		if err != nil {
			break // End of file
		}

		if tipo == '1' {
			// Leer estudiante
			var student structures.Student
			err = student.ReadFromFile(fullPath, offset)
			if err != nil {
				fmt.Println("Error leyendo estudiante:", err)
				return
			}
			students = append(students, student)
			offset += int64(stuSize)
			//fmt.Println("Estudiante leido")
			//fmt.Println(offset)
		} else if tipo == '2' {
			//fmt.Println("leyendo profesor")
			// Leer profesor
			var teacher structures.Teacher
			err = teacher.ReadFromFile(fullPath, offset)
			if err != nil {
				fmt.Println("Error leyendo profesor:", err)
				return
			}
			teachers = append(teachers, teacher)
			offset += int64(teachSize)
		} else {
			fmt.Println("Tipo desconocido:", tipo)
			break
		}
	}

	// Mostrar los resultados
	fmt.Println("Estudiantes:")
	for _, student := range students {
		fmt.Println(student.ToShow())
	}

	fmt.Println("Profesores:")
	for _, teacher := range teachers {
		fmt.Println(teacher.ToShow())
	}
}

func validateId(id string) bool {

	//len debe ser 5
	if len(id) > 5 {
		fmt.Println("El id debe tener como maximo 5 caracteres")
		return false
	}
	isUnic := true
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("error al abrir la datasheet")
		return false
	}
	defer file.Close()
	offset := int64(0)
	stuSize := binary.Size(structures.Student{})
	teachSize := binary.Size(structures.Teacher{})
	for {
		_, err = file.Seek(offset, 0)
		//fmt.Println("offset: ", offset)
		if err != nil {
			fmt.Println("error al buscar el offset")
			return false
		}
		var tipo byte
		err = binary.Read(file, binary.LittleEndian, &tipo)
		if err != nil {
			break
		}
		//fmt.Println("tipo: ", string(tipo))
		if tipo == '1' {
			var student structures.Student
			err = student.ReadFromFile(fullPath, offset)
			if err != nil {
				fmt.Println("Error leyendo estudiante:", err)
				return false
			}
			//fmt.Println("id estudiante: ", string(student.Id_Estu[:]))
			//fmt.Println("id a verficar: ", id)
			studentId := strings.TrimSpace(string(student.Id_Estu[:]))
			studentId = strings.Trim(studentId, "\x00")
			//size of cadenas
			//fmt.Println("size of cadenas: ", len(studentId), len(id))
			if studentId == id {
				fmt.Println("El id ya existe")
				isUnic = false
				break
			}
			offset += int64(stuSize)
		} else if tipo == '2' {
			var teacher structures.Teacher
			err = teacher.ReadFromFile(fullPath, offset)
			if err != nil {
				fmt.Println("Error leyendo profesor:", err)
				return false
			}
			//fmt.Println("id profesor: ", string(teacher.Id_profesor[:]))
			//fmt.Println("id a verficar: ", id)
			teacherId := strings.TrimSpace(string(teacher.Id_profesor[:]))
			teacherId = strings.Trim(teacherId, "\x00")
			if teacherId == id {
				isUnic = false
				fmt.Println("El id ya existe")
				break
			}
			offset += int64(teachSize)
		} else {
			fmt.Println("Tipo desconocido:", tipo)
			break
		}
	}
	return isUnic
}
