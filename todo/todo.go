package todo

import (
	"fmt"
	"time"
)

// Todo représente une tâche à faire
type Todo struct {
	ID         int
	Task       string
	Done       bool
	Priority   int
	Echéance   time.Time
	SousTâches []Todo
}

// Todos gère une liste de tâches
type Todos []Todo

// Ajouter ajoute une nouvelle tâche à la liste
func (t *Todos) Ajouter(task string) {
	*t = append(*t, Todo{ID: len(*t) + 1, Task: task})
}

// Lister affiche toutes les tâches
func (t *Todos) Lister() {
	for _, todo := range *t {
		doneStatus := "Non terminé"
		if todo.Done {
			doneStatus = "Terminé"
		}
		fmt.Printf("%d: %s [%s]\n", todo.ID, todo.Task, doneStatus)
	}
}

// ListerTerminées liste uniquement les tâches terminées
func (t *Todos) ListerTerminées() {
	for _, todo := range *t {
		if todo.Done {
			fmt.Printf("%d: %s [Terminé]\n", todo.ID, todo.Task)
		}
	}
}

// ListerNonTerminées liste uniquement les tâches non terminées
func (t *Todos) ListerNonTerminées() {
	for _, todo := range *t {
		if !todo.Done {
			fmt.Printf("%d: %s [Non Terminé]\n", todo.ID, todo.Task)
		}
	}
}

// DéfinirPriorité définit la priorité d'une tâche spécifique
func (t *Todos) DéfinirPriorité(id, priorité int) {
	for i := range *t {
		if (*t)[i].ID == id {
			(*t)[i].Priority = priorité
			break
		}
	}
}

// DéfinirÉchéance définit l'échéance d'une tâche spécifique
func (t *Todos) DéfinirÉchéance(id int, échéance string) error {
	date, err := time.Parse("2006-01-02", échéance)
	if err != nil {
		return err
	}
	for i := range *t {
		if (*t)[i].ID == id {
			(*t)[i].Echéance = date
			return nil
		}
	}
	return fmt.Errorf("tâche non trouvée")
}

// AjouterSousTâche ajoute une sous-tâche à une tâche spécifique
func (t *Todos) AjouterSousTâche(idTâche int, sousTâche string) error {
	for i := range *t {
		if (*t)[i].ID == idTâche {
			(*t)[i].SousTâches = append((*t)[i].SousTâches, Todo{
				ID:   len((*t)[i].SousTâches) + 1,
				Task: sousTâche,
			})
			return nil
		}
	}
	return fmt.Errorf("tâche principale non trouvée")
}

// ListerSousTâches affiche les sous-tâches d'une tâche spécifique
func (t *Todos) ListerSousTâches(idTâche int) error {
	for _, todo := range *t {
		if todo.ID == idTâche {
			fmt.Printf("Sous-tâches pour '%s':\n", todo.Task)
			for _, st := range todo.SousTâches {
				fmt.Printf("  %d: %s\n", st.ID, st.Task)
			}
			return nil
		}
	}
	return fmt.Errorf("tâche principale non trouvée")
}

// TerminerSousTache affiche les sous-taches terminées
func (t *Todos) TerminerSousTâche(idTâche, idSousTâche int) error {
	for i := range *t {
		if (*t)[i].ID == idTâche {
			for j := range (*t)[i].SousTâches {
				if (*t)[i].SousTâches[j].ID == idSousTâche {
					(*t)[i].SousTâches[j].Done = true
					return nil
				}
			}
			return fmt.Errorf("sous-tâche non trouvée")
		}
	}
	return fmt.Errorf("tâche principale non trouvée")
}

// Terminer marque une tâche comme terminée
func (t *Todos) Terminer(id int) {
	for i := range *t {
		if (*t)[i].ID == id {
			(*t)[i].Done = true
			break
		}
	}
}

// Modifier modifie la tâche spécifiée par son ID
func (t *Todos) Modifier(id int, nouvelleTâche string) error {
	for i := range *t {
		if (*t)[i].ID == id {
			(*t)[i].Task = nouvelleTâche
			return nil
		}
	}
	return fmt.Errorf("tâche non trouvée")
}

// Supprimer supprime la tâche spécifiée par son ID
func (t *Todos) Supprimer(id int) error {
	for i, todo := range *t {
		if todo.ID == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("tâche non trouvée")
}
