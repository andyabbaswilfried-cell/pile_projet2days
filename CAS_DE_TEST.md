# Catalogue exhaustif des cas de test

Ce catalogue couvre les classes de cas du projet. Les tests exécutables
correspondants se trouvent dans `fonctions_pile/fonctions_test.go`,
`push-swap/main_test.go` et `push-swap/exhaustive_test.go`.

Lancer toute la suite depuis la racine :

```sh
go test ./...
```

## Validation des valeurs

| Cas | Exemple | Résultat attendu |
|---|---|---|
| Aucune valeur | liste vide | valide |
| Un entier nul ou positif | `0`, `42` | valide |
| Entier signé | `-2`, `+3` | valide pour la validation des entiers |
| Texte vide, décimal ou non numérique | `""`, `1.5`, `abc` | invalide |
| Entier hors limites du type `int` | valeur trop grande | invalide |
| Valeurs toutes distinctes | `1 2 3` | valide |
| Valeurs répétées | `1 2 1` | invalide |

Le programme `push-swap` documente actuellement la prise en charge des valeurs
non négatives dans `README.md`; la validation entière seule n'est donc pas une
garantie de prise en charge par l'algorithme pour les nombres négatifs.

## Opérations sur les piles

Chaque instruction valide est vérifiée avec une pile assez grande pour observer
son effet :

| Instruction(s) | Cas vérifié |
|---|---|
| `pa`, `pb` | transfert du sommet d'une pile vers l'autre |
| `sa`, `sb`, `ss` | échange d'un ou des deux sommets |
| `ra`, `rb`, `rr` | rotation du sommet vers le fond |
| `rra`, `rrb`, `rrr` | rotation du fond vers le sommet |
| Nom inconnu | rejet de l'instruction |

Pour chacune des onze instructions, les tests couvrent les combinaisons de piles
vides, à un élément et à plusieurs éléments. Une opération nécessitant un sommet
ou une rotation sur une pile trop courte doit rester valide et ne rien modifier.

## Vérification de l'état final

| Pile A | Pile B | Résultat attendu |
|---|---|---|
| Vide | Vide | triée |
| Un élément | Vide | triée |
| Strictement croissante | Vide | triée |
| Décroissante ou contenant des doublons | Vide | non triée |
| Quelle que soit l'ordre | Non vide | non triée |

## Génération du tri

Les tests unitaires couvrent les cas vide, singleton, déjà trié, ordre inversé
et valeurs mélangées. En complément, `TestAlgoTrierAllPermutationsUpToSeven`
teste toutes les permutations de valeurs distinctes de `0` à `n-1`, pour chaque
taille `n` de 0 à 7 (5 914 permutations au total). Chaque séquence générée est
rejouée par le même interpréteur d'instructions que le vérificateur; la pile A
doit être strictement croissante et la pile B vide à la fin.
