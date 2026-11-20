# Push Swap

## Présentation

Push Swap est un projet de programmation en Go qui consiste à trier une liste d'entiers en utilisant deux piles (`A` et `B`) et un ensemble limité d'opérations.

L'objectif est de trier les nombres de la pile A dans l'ordre croissant avec le moins d'instructions possible.

## Technologies utilisées

- Go
- Git
- GitHub

## Structure du projet

```text
push-swap/
├── cmd/
│   ├── push-swap/
│   │   └── main.go
│   └── checker/
│       └── main.go
├── internal/
│   ├── operations/
│   │   ├── operations.go
│   │   └── operations_test.go
│   ├── parser/
│   │   ├── parser.go
│   │   └── parser_test.go
│   ├── sorting/
│   │   ├── sorting.go
│   │   └── sorting_test.go
│   └── stack/
│       ├── stack.go
│       └── stack_test.go
├── go.mod
└── README.md
```

## Compilation

Pour compiler le programme Push Swap :

```bash
go build -o push-swap.exe ./cmd/push-swap
```

Pour compiler le Checker :

```bash
go build -o checker.exe ./cmd/checker
```

## Utilisation

### Trier une liste

```bash
./push-swap.exe 3 2 1
```

Le programme affiche les instructions nécessaires pour trier les nombres.

### Vérifier le résultat

```bash
./push-swap.exe 3 2 1 | ./checker.exe 3 2 1
```

Résultat attendu :

```text
OK
```

Le Checker vérifie que la pile A est triée dans l'ordre croissant et que la pile B est vide.

## Instructions disponibles

| Instruction | Description |
|---|---|
| sa | Échanger les deux premiers éléments de A |
| sb | Échanger les deux premiers éléments de B |
| ss | Effectuer sa et sb |
| pa | Transférer le premier élément de B vers A |
| pb | Transférer le premier élément de A vers B |
| ra | Rotation de A |
| rb | Rotation de B |
| rr | Effectuer ra et rb |
| rra | Rotation inverse de A |
| rrb | Rotation inverse de B |
| rrr | Effectuer rra et rrb |

## Algorithme de tri

Le programme utilise les méthodes étudiées pendant les séances d'algorithmique.

### 1. Tri de petites listes

- Pour 2 nombres : utilisation de `sa` si nécessaire.
- Pour 3 nombres : utilisation de `trierTrois`, avec deux opérations maximum.
- Pour 4 ou 5 nombres : utilisation de `trierCinq`.

La méthode `trierCinq` place les plus petits nombres dans B, trie les trois nombres restants, puis remet les éléments dans A.

### 2. Tri par blocs

Pour les listes de plus de 5 nombres, le programme utilise `trierParBlocs`.

Cette méthode se déroule en deux phases :

- Conversion des nombres en rangs.
- Division des rangs en blocs.
- Transfert des éléments de A vers B selon leur rang.
- Utilisation de `rb` pour organiser les éléments dans B.
- Recherche du plus grand rang dans B.
- Utilisation de `rb` ou `rrb` pour placer ce rang au sommet.
- Transfert des éléments vers A avec `pa`.

La taille des blocs est de 15 pour les listes de 100 éléments maximum et de 30 au-delà.

## Tests

Pour exécuter tous les tests :

```bash
go test ./...
```

Pour afficher chaque test :

```bash
go test -v ./...
```

Pour vérifier le formatage :

```bash
gofmt -w .
```

Pour vérifier les erreurs potentielles :

```bash
go vet ./...
```

Pour compiler tous les packages :

```bash
go build ./...
```

## Gestion des erreurs

Le programme gère notamment :

- Les arguments invalides.
- Les nombres en double.
- Les entrées non numériques.
- Les dépassements de capacité des entiers.
- Les opérations invalides dans le Checker.

