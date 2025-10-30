# Power4 - Puissance 4 en Go

Un jeu de Puissance 4 (Connect Four) développé en Go avec une interface web.

## 🚀 Démarrage rapide

```bash
go run .
```

Puis ouvrez votre navigateur à l'adresse : `http://localhost:8080`

## 📁 Structure du projet

```
Power4/
├── main.go          # Serveur HTTP et routes
├── power4.go        # Logique du jeu
├── index.html       # Interface utilisateur
└── template/
    └── style.css    # Styles CSS
```

## 🎮 Fonctionnalités

- Interface web responsive
- Jeu Puissance 4 classique (6 lignes × 7 colonnes)
- Détection automatique des victoires (horizontal, vertical, diagonal)
- Détection des matchs nuls
- Bouton de réinitialisation
- Alternance automatique des joueurs (X et O)

---

## 📚 Documentation des fonctions Go

### `main.go`

#### `func main()`
Point d'entrée de l'application. Configure le serveur HTTP et les routes.

**Routes disponibles :**
- `GET /` - Affiche la page du jeu
- `POST /play` - Joue un coup dans une colonne
- `POST /reset` - Réinitialise le jeu
- `/static/` - Sert les fichiers statiques (CSS, images, etc.)

---

### `power4.go`

#### `type Power struct`
Structure principale du jeu contenant :
```go
type Power struct {
    Table         [6][7]string  // Grille de jeu 6×7
    Player        [2]string     // Symboles des joueurs ["X", "O"]
    CurrentPlayer string        // Joueur actuel
    Winner        string        // Gagnant (vide si pas de gagnant)
    IsDraw        bool          // True si match nul
}
```

#### `func NewPower() *Power`
Crée une nouvelle instance du jeu avec une grille vide.

**Retour :** Pointeur vers une nouvelle structure `Power` initialisée

**Exemple :**
```go
power := NewPower()
```

---

#### `func Starter(Player [2]string)`
Affiche un message de bienvenue avec les noms/symboles des joueurs.

**Paramètres :**
- `Player [2]string` - Tableau contenant les symboles des deux joueurs

**Exemple :**
```go
Starter([2]string{"X", "O"})
```

---

#### `func Play(p *Power, column int) bool`
Joue un pion dans la colonne spécifiée.

**Paramètres :**
- `p *Power` - Pointeur vers l'instance du jeu
- `column int` - Numéro de la colonne (0-6)

**Retour :** 
- `true` si le coup a été joué avec succès
- `false` si la colonne est pleine ou invalide

**Logique :**
1. Vérifie que la colonne est valide (0-6)
2. Trouve la case vide la plus basse dans la colonne
3. Place le pion du joueur actuel
4. Change le joueur actuel
5. Affiche la grille (en console)

**Exemple :**
```go
success := Play(power, 3) // Joue dans la colonne 3
```

---

#### `func (p *Power) switchPlayer()`
Change le joueur actuel (X → O ou O → X).

**Exemple :**
```go
power.switchPlayer()
```

---

#### `func (p *Power) PrintBoard()`
Affiche la grille de jeu dans la console.

**Format d'affichage :**
```
. . . . . . . 
. . . . . . . 
. . . . . . . 
. . . X . . . 
. . . O X . . 
. . O O X . . 
```

---

#### `func (p *Power) Win() bool`
Vérifie si le dernier coup a créé une victoire.

**Retour :** `true` si un joueur a gagné, `false` sinon

**Détections :**
- ✅ 4 pions horizontaux alignés
- ✅ 4 pions verticaux alignés
- ✅ 4 pions en diagonale (↗)
- ✅ 4 pions en diagonale (↘)

**Effet secondaire :** 
- Définit `p.Winner` avec le symbole du gagnant
- Affiche un message de victoire dans la console

**Exemple :**
```go
if power.Win() {
    fmt.Println("Partie terminée!")
}
```

---

#### `func (p *Power) IsDraw() bool`
Vérifie si la grille est pleine (match nul).

**Retour :** `true` si toutes les cases sont remplies, `false` sinon

**Exemple :**
```go
if power.IsDraw() {
    fmt.Println("Match nul!")
}
```

---

## 🌐 Documentation HTML/Template

### Structure du template `index.html`

#### Variables disponibles dans le template

```go
.Power          // Instance du jeu
.Power.Table    // Grille [6][7]string
.Power.CurrentPlayer  // Joueur actuel
.Power.Winner   // Gagnant (si existant)
.Power.IsDraw   // Match nul (bool)
.Cols           // []int{0,1,2,3,4,5,6} pour les boutons
```

#### Boucles Go Template

**Affichage de la grille :**
```html
{{range $i, $row := .Power.Table}}
    {{range $j, $cell := $row}}
        <div class="cell {{if eq $cell "X"}}x{{else if eq $cell "O"}}o{{end}}">
            {{$cell}}
        </div>
    {{end}}
{{end}}
```

**Génération des boutons de colonnes :**
```html
{{range $i := .Cols}}
    <button name="col" value="{{$i}}" type="submit">
        colonne {{$i}}
    </button>
{{end}}
```

#### Conditions dans le template

**Affichage du gagnant :**
```html
{{if .Power.Winner}}
    <h2>Le joueur {{.Power.Winner}} a gagné !</h2>
{{else if .Power.IsDraw}}
    <h2>Match nul !</h2>
{{end}}
```

---

## 🔧 Utilisation

### Jouer un coup

Envoyez une requête POST à `/play` avec le paramètre `col` :

```html
<form action="/play" method="post">
    <button name="col" value="0">Colonne 0</button>
    <button name="col" value="1">Colonne 1</button>
    <!-- ... -->
</form>
```

### Réinitialiser le jeu

Envoyez une requête POST à `/reset` :

```html
<form action="/reset" method="post">
    <button type="submit">Réinitialiser</button>
</form>
```

---

## 🎨 Classes CSS

- `.cell` - Case de la grille
- `.cell.x` - Case avec pion X (rouge)
- `.cell.o` - Case avec pion O (jaune)
- `.container` - Conteneur principal
- `.grid` - Grille de jeu
- `.info` - Informations du jeu
- `.cols` - Conteneur des boutons

---

## 📝 Notes techniques

- **Port par défaut :** 8080
- **Template engine :** `html/template` (standard Go)
- **Méthode HTTP :** GET pour affichage, POST pour actions
- **Redirection :** Après chaque action, redirection vers `/` (pattern PRG)

---

## 🐛 Résolution de problèmes

### Erreur : `bind: Only one usage of each socket address`

Le port 8080 est déjà utilisé. Tuez le processus :

```powershell
netstat -ano | findstr :8080
taskkill /PID <PID> /F
```

### Erreur de template

Vérifiez que `index.html` est dans le même dossier que `main.go`.

---

## 📄 Licence

Projet éducatif - Libre d'utilisation