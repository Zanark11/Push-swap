# Push-swap

## Description

Push-swap est un projet développé en Go.

L'objectif est de trier une liste d'entiers dans l'ordre croissant en utilisant deux piles (A et B) et un ensemble limité d'instructions.

Le projet contient deux programmes :
- **push-swap** : génère une suite d'instructions pour trier les nombres.
- **checker** : exécute les instructions et vérifie si le tri est correct.

## Instructions autorisées

| Instruction | Description |
|---|---|
| sa | Échange les deux premiers éléments de A |
| sb | Échange les deux premiers éléments de B |
| ss | Effectue sa et sb |
| pa | Déplace le sommet de B vers A |
| pb | Déplace le sommet de A vers B |
| ra | Effectue une rotation de A |
| rb | Effectue une rotation de B |
| rr | Effectue une rotation de A et B |
| rra | Effectue une rotation inverse de A |
| rrb | Effectue une rotation inverse de B |
| rrr | Effectue une rotation inverse de A et B |

## Installation

Cloner le dépôt :

```bash
git clone https://github.com/Zanark11/Push-swap.git
cd Push-swap
```

## Compilation

Créer le dossier contenant les exécutables :

```bash
mkdir -p bin
```

Compiler push-swap :

```bash
go build -o bin/push-swap ./push-swap
```

Compiler checker :

```bash
go build -o bin/checker ./checker
```

## Utilisation

### Trier une liste de nombres

```bash
./bin/push-swap "4 2 7 1"
```

Le programme affiche les instructions nécessaires pour trier les nombres.

### Vérifier le résultat du tri

```bash
./bin/push-swap "4 2 7 1" | ./bin/checker "4 2 7 1"
```

Résultat attendu :

```text
OK
```

Le checker affiche :
- `OK` si la pile A est triée et la pile B est vide.
- `KO` si la pile A n'est pas triée ou si la pile B n'est pas vide.
- `Error` si les arguments ou les instructions sont invalides.

## Tests

Pour lancer tous les tests du projet :

```bash
go test -v ./...
```

Pour compiler tous les packages :

```bash
go build ./...
```

## Technologies utilisées

- Go
- Git
- GitHub
