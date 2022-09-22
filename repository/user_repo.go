package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/salimkun/sample-crud-go/model"
)

func CreateUser(user *model.User) error {
	getAllfile, err := ReadFile()
	if err != nil {
		return err
	}
	if len(getAllfile) > 0 {
		user.ID = getAllfile[len(getAllfile)].ID + 1
	} else {
		user.ID = 1
	}

	getAllfile = append(getAllfile, user)
	file, _ := json.MarshalIndent(getAllfile, "", " ")

	err = ioutil.WriteFile("repository/users.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(users []*model.User) error {
	file, _ := json.MarshalIndent(users, "", " ")

	err := ioutil.WriteFile("repository/users.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile() ([]*model.User, error) {
	users := []*model.User{}
	jsonFile, err := os.Open("repository/users.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &users)

	return users, nil
}
