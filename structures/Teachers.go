package structures

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Teacher struct {
	Tipo        [1]byte  //tipo
	Id_profesor [5]byte  //id_profesor
	CUI         [13]byte //cui
	Nombre      [25]byte //nombre
	Curso       [25]byte //cursos
	//total size =	69
}

//Write to file
func (t *Teacher) WriteToFile(fullPath string) error {
	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Error al abrir la dataSheet en Teacher WriteToFile")
	}

	defer file.Close()

	teachSize := binary.Size(*t)
	offset, err := FindFreeBlock(file, teachSize)
	if err != nil {
		return fmt.Errorf("Error findFreeBlock Teacher WriteToFile")
	}

	_, err = file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("Error en file Seek Teacher WriteToFile")
	}
	//fmt.Println("offset profesor: ", offset)
	err = binary.Write(file, binary.LittleEndian, t)
	if err != nil {
		return fmt.Errorf("Error en binary.Write Teacher WriteToFile")
	}
	return nil
}

func (t *Teacher) ReadFromFile(fullPath string, offset int64) error {
	file, err := os.Open(fullPath)
	if err != nil {
		return fmt.Errorf("Error al abrir la dataSheet en Teacher ReadFromFile")
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("Error en file Seek Teacher ReadFromFile")
	}

	err = binary.Read(file, binary.LittleEndian, t)
	if err != nil {
		return fmt.Errorf("Error en binary.Read Teacher ReadFromFile")
	}
	return nil
}

//toSHow
func (t Teacher) ToShow() string {
	return fmt.Sprintf("Id_profesor: %s, CUI: %s, Nombre: %s, Curso: %s",
		string(t.Id_profesor[:]),
		string(t.CUI[:]),
		string(t.Nombre[:]),
		string(t.Curso[:]))
}
