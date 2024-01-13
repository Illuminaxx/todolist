# Todo List Application

## Description
Cette application Todo List est un outil de ligne de commande pour gérer vos tâches quotidiennes. Elle permet aux utilisateurs d'ajouter, lister, marquer comme terminées, et supprimer des tâches, ainsi que de gérer des sous-tâches, définir des priorités et des échéances.

## Fonctionnalités
- Ajouter des tâches et des sous-tâches.
- Lister toutes les tâches, les tâches terminées ou non terminées.
- Marquer des tâches comme terminées.
- Définir et modifier la priorité des tâches.
- Définir et modifier les échéances des tâches.
- Supprimer des tâches.

## Installation
Pour installer cette application, vous aurez besoin de Go installé sur votre système. Suivez ces étapes :

1. Clonez le dépôt
2. Naviguez dans le dossier cloné
3. Compilez le projet : ```go build```

## Utilisation
Après avoir compilé l'application, vous pouvez l'utiliser en ligne de commande. Voici quelques exemples de commandes :

- Ajouter une nouvelle tâche :
```./todolist add "Votre nouvelle tâche```

- Lister toutes les tâches :
```./todolist list```

- Marquer une tâche comme terminée :
```./todolist terminer <id de la tâche>```

- Définir la priorité d'une tâche :
```./todolist priorité <id de la tâche> <niveau de priorité>```

- Définir l'échéance d'une tâche :
```./todolist échéance <id de la tâche> <YYYY-MM-DD>```

- Supprimer une tâche :
```./todolist supprimer <id de la tâche>```

## Contribution
Les contributions à ce projet sont les bienvenues. N'hésitez pas à ouvrir des issues ou à soumettre des pull requests.
