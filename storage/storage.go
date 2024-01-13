package storage

import (
	"encoding/json"
	"os"
	"todoList/todo"
)

const fileName = "todolist.json"

// Sauvegarder sauvegarde les tâches Todo dans un fichier
func Sauvegarder(todos todo.Todos) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

// Charger charge les tâches Todo depuis un fichier
func Charger(todos *todo.Todos) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			// Créer un nouveau fichier si celui-ci n'existe pas
			*todos = make(todo.Todos, 0)
			return Sauvegarder(*todos) // Sauvegardez la liste vide pour créer le fichier
		}
		return err
	}
	return json.Unmarshal(data, todos)
}
