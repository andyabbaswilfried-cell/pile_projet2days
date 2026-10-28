# pile_projet2days

Ce projet trie une pile d'entiers en générant des instructions de manipulation,
puis permet de vérifier ces instructions avec un programme séparé.

## Prérequis

- Go installé (la version du module est indiquée dans `go.mod`).

## Utilisation

### Générer les instructions

Depuis la racine du projet, passez les nombres dans une seule chaîne entre
guillemets :

```sh
go run ./push-swap "8 3 6 1 7 2"
```

Le programme affiche les instructions de tri, une par ligne. Par exemple :

```text
pb
ra
```

Les valeurs doivent être des entiers distincts et non négatifs. L'algorithme
utilise un tri radix binaire et ne prend pas actuellement en charge les valeurs
négatives.

### Vérifier une séquence d'instructions

Enregistrez les instructions générées dans un fichier, par exemple
`instructions.txt`, puis fournissez la pile initiale au vérificateur.

Dans PowerShell :

```powershell
Get-Content instructions.txt | go run ./checker "8 3 6 1 7 2"
```

Le vérificateur lit une instruction par ligne. Il affiche `OK` si la pile A est
triée à la fin et que la pile B est vide ; sinon, il affiche `KO`. Les
instructions inconnues entraînent une erreur.

### Lancer les tests

Depuis la racine du projet :

```sh
go test ./...
```

## Instructions disponibles

Les instructions suivantes manipulent les piles A et B. Le sommet d'une pile
est son premier élément.

| Instruction | Effet |
|---|---|
| `sa` / `sb` | Échange les deux premiers éléments de A / B. |
| `ss` | Effectue `sa` et `sb`. |
| `pa` / `pb` | Déplace le sommet de B vers A / de A vers B. |
| `ra` / `rb` | Place le sommet de A / B au fond de la pile. |
| `rr` | Effectue `ra` et `rb`. |
| `rra` / `rrb` | Place le dernier élément de A / B au sommet. |
| `rrr` | Effectue `rra` et `rrb`. |

Une opération qui ne peut pas être effectuée sur une pile trop courte ne la
modifie pas.

## Organisation du projet

- `push-swap/` : programme qui génère les instructions de tri.
- `checker/` : programme qui applique des instructions et vérifie le résultat.
- `fonctions_pile/` : opérations partagées entre les deux programmes.
