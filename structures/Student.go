package structures

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Tipo    [1]byte  //tipo
	Id_Estu [5]byte  //id_estdui
	CUI     [13]byte //cui
	Nombre  [25]byte //nombre
	Carnet  [25]byte //carnet 201903022
	//total size =	69
}

//Write to file

func (s *Student) WriteToFile(fullPath string) error {
	fmt.Println("writeToFile")
	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Error al abrir la dataSheet en Studen WriteToFile")
	}
	defer file.Close()

	stuSize := binary.Size(*s)
	fmt.Println("stuSize", stuSize)

	offset, err := FindFreeBlock(file, stuSize)
	if err != nil {
		return fmt.Errorf("Error findFreeBlock Studntent WriteToFile")
	}

	_, err = file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("Error en file Seek Student WrtiteToFile")
	}
	fmt.Println("offset", offset)
	err = binary.Write(file, binary.LittleEndian, s)
	if err != nil {
		return fmt.Errorf("Error en binary.Write Stundet WriteToFile")
	}
	return nil
}
func (s *Student) ReadFromFile(fullPath string, offset int64) error {
	file, err := os.Open(fullPath)
	if err != nil {
		return fmt.Errorf("Error al abrir la dataSheet en Student ReadFromFile")
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("Error en file Seek Student ReadFromFile")
	}

	err = binary.Read(file, binary.LittleEndian, s)
	if err != nil {
		return fmt.Errorf("Error en binary.Read Student ReadFromFile")
	}
	return nil
}
func (s Student) ToShow() string {
	return fmt.Sprintf("Carnet: %s, CUI: %s, Name: %s, ID: %s",
		strings.TrimSpace(string(s.Carnet[:])),
		strings.TrimSpace(string(s.CUI[:])),
		strings.TrimSpace(string(s.Nombre[:])),
		strings.TrimSpace(string(s.Id_Estu[:])),
	)
}

//find a freeBlock

func FindFreeBlock(file *os.File, blocksize int) (int64, error) {

	buffer := make([]byte, blocksize)
	var offset int64 //posicion actual
	fmt.Println("blocksize", blocksize)
	fmt.Println("offset: ", offset)
	for {
		_, err := file.ReadAt(buffer, offset)
		if err != nil {
			break
		}

		isFree := true

		for _, byte := range buffer {
			if byte != 0 { //no todo el bloque esta libre, algun byte esta ocupado buscar otro bloque
				isFree = false
				break
			}
		}
		if isFree {
			return offset, nil
		}

		fmt.Println("offset: ", offset)
		offset += int64(blocksize)

	}

	return offset, nil

}
