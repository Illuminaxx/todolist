package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"todoList/storage"
	"todoList/todo"
)

func main() {
	var todos todo.Todos

	// Charger les tâches au démarrage
	err := storage.Charger(&todos)
	if err != nil {
		fmt.Println("Erreur lors du chargement des tâches:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: todolist <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todolist add <task>")
			return
		}
		task := strings.Join(os.Args[2:], " ")
		todos.Ajouter(task)
		fmt.Printf("Tâche ajoutée: %s\n", task)

		// Sauvegarder après l'ajout d'une tâche
		err := storage.Sauvegarder(todos)
		if err != nil {
			fmt.Println("Erreur lors de la sauvegarde des tâches:", err)
		}

	case "list":
		todos.Lister()

	case "terminer":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todolist terminer <id de la tâche>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Veuillez fournir un numéro d'ID valide.")
			return
		}
		todos.Terminer(id)
		fmt.Printf("Tâche %d marquée comme terminée.\n", id)

		// Sauvegarder après avoir marqué une tâche comme terminée
		err = storage.Sauvegarder(todos)
		if err != nil {
			fmt.Println("Erreur lors de la sauvegarde des tâches:", err)
		}
	case "list done":
		todos.ListerTerminées()
	case "list not done":
		todos.ListerNonTerminées()
	case "priorité":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todolist priorité <id de la tâche> <nouvelle priorité>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID invalide.")
			return
		}
		priorité, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Priorité invalide.")
			return
		}
		todos.DéfinirPriorité(id, priorité)
		fmt.Printf("Priorité définie pour la tâche %d.\n", id)

		// Sauvegarder après modification de la priorité
		err = storage.Sauvegarder(todos)
		if err != nil {
			fmt.Println("Erreur lors de la sauvegarde des tâches:", err)
		}
	case "échéance":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todolist échéance <id de la tâche> <date YYYY-MM-DD>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID invalide.")
			return
		}
		date := os.Args[3]
		err = todos.DéfinirÉchéance(id, date)
		if err != nil {
			fmt.Println("Erreur lors de la définition de l'échéance:", err)
			return
		}
		fmt.Printf("Échéance définie pour la tâche %d.\n", id)

		// Sauvegarder après modification de l'échéance
		err = storage.Sauvegarder(todos)
		if err != nil {
			fmt.Println("Erreur lors de la sauvegarde des tâches:", err)
		}
	case "ajouter-soustâche":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todolist ajouter-soustâche <id de la tâche principale> <sous-tâche>")
			return
		}
		idTâche, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID de tâche principale invalide.")
			return
		}
		sousTâche := strings.Join(os.Args[3:], " ")
		err = todos.AjouterSousTâche(idTâche, sousTâche)
		if err != nil {
			fmt.Println("Erreur lors de l'ajout de la sous-tâche:", err)
			return
		}
		fmt.Printf("Sous-tâche ajoutée à la tâche %d.\n", idTâche)

	case "lister-soustâches":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todolist lister-soustâches <id de la tâche principale>")
			return
		}
		idTâche, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID de tâche principale invalide.")
			return
		}
		err = todos.ListerSousTâches(idTâche)
		if err != nil {
			fmt.Println("Erreur lors du listage des sous-tâches:", err)
		}

	case "terminer-soustâche":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todolist terminer-soustâche <id de la tâche> <id de la sous-tâche>")
			return
		}
		idTâche, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID invalide")
			return
		}
		idSousTâche, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("ID invalide.")
			return
		}
		err = todos.TerminerSousTâche(idTâche, idSousTâche)
		if err != nil {
			fmt.Println("Erreur lors de la marque de la sous-tâche comme terminée:", err)
			return
		}
		fmt.Println("Sous-tâche marquée comme terminée avec succès.")

	case "modifier":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todolist modifier <id> <nouvelle tâche>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID invalide.")
			return
		}
		nouvelleTâche := strings.Join(os.Args[3:], " ")
		err = todos.Modifier(id, nouvelleTâche)
		if err != nil {
			fmt.Println("Erreur lors de la modification:", err)
			return
		}
		fmt.Println("Tâche modifiée avec succès.")

	case "supprimer":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todolist supprimer <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID invalide.")
			return
		}
		err = todos.Supprimer(id)
		if err != nil {
			fmt.Println("Erreur lors de la suppression:", err)
			return
		}
		fmt.Println("Tâche supprimée avec succès.")

	default:
		fmt.Printf("Commande inconnue: %s\n", command)
	}
}
