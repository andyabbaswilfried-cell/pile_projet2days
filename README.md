# pile_projet2days

Ce projet contient deux programmes Go :

- **`push-swap`** trie une pile d'entiers et affiche les instructions nécessaires.
- **`checker`** applique une séquence d'instructions et vérifie le résultat.

## Prérequis

- Go **1.27.1** ou une version compatible avec celle indiquée dans [`go.mod`](./go.mod).

Toutes les commandes ci-dessous sont à lancer depuis la racine du dépôt.

## Démarrage rapide

Générer des instructions pour trier les nombres `8 3 6 1 7 2` :

```powershell
go run ./push-swap "8 3 6 1 7 2"
```

Les instructions sont affichées une par ligne. Pour vérifier automatiquement
qu'elles trient bien la pile d'origine, envoyez la sortie de `push-swap` à
`checker` :

```powershell
go run ./push-swap "8 3 6 1 7 2" | go run ./checker "8 3 6 1 7 2"
```

Résultat attendu :

```text
OK
```

## Utilisation détaillée

### Trier une pile

Fournissez tous les nombres dans **un seul argument entre guillemets** :

```powershell
go run ./push-swap "5 2 8 1"
```

Le programme utilise un tri radix binaire et affiche les instructions à exécuter
dans l'ordre. Si la pile est déjà triée ou contient zéro ou un élément, aucune
instruction n'est nécessaire.

Les valeurs doivent être des entiers distincts qui tiennent dans le type `int`
de Go. Les nombres négatifs sont décalés à partir de la valeur minimale avant
le tri; évitez toutefois les entrées dont l'écart entre la valeur minimale et
maximale dépasse la capacité de `int`.

En cas d'argument manquant, de valeur non entière ou de doublon, le programme
affiche `Error` sur la sortie d'erreur.

### Vérifier des instructions

Le vérificateur lit une instruction par ligne sur l'entrée standard. Par
exemple, dans PowerShell, enregistrez d'abord les instructions, puis vérifiez-les :

```powershell
go run ./push-swap "8 3 6 1 7 2" | Set-Content instructions.txt
Get-Content instructions.txt | go run ./checker "8 3 6 1 7 2"
```

- `OK` signifie que la pile A est triée dans l'ordre croissant et que la pile B
  est vide.
- `KO` signifie que la séquence est valide, mais que l'état final n'est pas trié.
- `Error` est affiché sur la sortie d'erreur si les valeurs d'entrée sont
  invalides ou si une instruction est inconnue.

Une séquence d'instructions vide est valide : le vérificateur renvoie `OK` si la
pile fournie est déjà triée (y compris une pile vide ou à un seul élément).

## Instructions des piles

Le sommet d'une pile est son premier élément.

| Instruction | Effet |
|---|---|
| `sa` / `sb` | Échange les deux premiers éléments de A / B. |
| `ss` | Effectue `sa` et `sb`. |
| `pa` | Déplace le sommet de B au sommet de A. |
| `pb` | Déplace le sommet de A au sommet de B. |
| `ra` / `rb` | Place le sommet de A / B au fond de la pile. |
| `rr` | Effectue `ra` et `rb`. |
| `rra` / `rrb` | Place le dernier élément de A / B au sommet. |
| `rrr` | Effectue `rra` et `rrb`. |

Une opération qui ne peut pas être effectuée parce qu'une pile est trop courte
ne modifie pas cette pile.

## Tests

Exécuter tous les tests :

```powershell
go test ./...
```

Exécuter uniquement le test exhaustif du tri :

```powershell
go test ./push-swap -run '^TestAlgoTrierAllPermutationsUpToSeven$' -v
```

Ce test rejoue les instructions pour toutes les permutations de 0 à 7 valeurs
distinctes, pour chaque taille de pile de 0 à 7.

Le catalogue des scénarios est décrit dans [`CAS_DE_TEST.md`](./CAS_DE_TEST.md).

## Organisation du dépôt

| Chemin | Rôle |
|---|---|
| [`push-swap/`](./push-swap/) | Génération des instructions de tri et tests associés. |
| [`checker/`](./checker/) | Application des instructions et vérification de l'état final. |
| [`fonctions_pile/`](./fonctions_pile/) | Opérations sur les piles, interpréteur d'instructions et tests partagés. |
| [`samples/`](./samples/) | Fichier d'exemple. |
