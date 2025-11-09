package services

import (
	//"errors"
	"fmt"
	"os"
	"encoding/json"
	"example.com/library_management/models"
	"path/filepath"
)

var p string  = filepath.Join("services/", "library.json")

// This is where the file work will happen
func ReadLibraryFile ()(*models.Library, error){
	_, err := os.Stat(p)
	if os.IsNotExist(err){
		file, createErr := os.Create(p)

		if createErr != nil {
			fmt.Printf("Error creating File: %v\n",  createErr)
			return nil, createErr
		}
		defer file.Close()
		// can return new struct with nothing inside
		return models.NewLibrary(), nil
	}else if err != nil{
		fmt.Printf("Error checking if file exists: %v\n",  err)
		return nil, err
	}
	// 	I can finally read the file content,
	return ByteToJson(p)
	

}

func ByteToJson(p string )(*models.Library, error){
	readData, err:= os.ReadFile(p)

	if err != nil {
		fmt.Println("Error reading from file", err)
		return nil, err
	}

	if len(readData) == 0{
		return models.NewLibrary() , nil
	}

	var library models.Library

	err = json.Unmarshal(readData, &library)
	if err != nil {
		fmt.Println("Error ocurred unmarshaling JSON:", err)
		return nil, err
	}

	return &library, nil
}

func WriteToLibraryFile(lib *models.Library)error{
	jsonData, err := json.MarshalIndent(lib, "", " ")
	if err != nil{
		fmt.Println("Error Marshaling JSON: ", err)
		return err
	}

	err = os.WriteFile(p, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to File", err)
		return err
	}
	
	return nil

}





