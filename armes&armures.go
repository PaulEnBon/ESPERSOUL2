package main

import "fmt"

type Armure struct {
	Nom        string
	Defense    int // Défense physique
	Resistance int // Résistance magique
	HP         int // Bonus de points de vie
}

type Competence struct {
	Nom         string
	Description string
	Degats      int
	CoutMana    int
	Type        string // "physique" ou "magique"
	TypeEffet   string // Type d'effet à créer
	Puissance   int    // Puissance de l'effet (0-5)
}

type Arme struct {
	Nom             string
	DegatsPhysiques int
	DegatsMagiques  int
	Precision       float64
	TauxCritique    float64
	Durabilite      int
	Competences     []Competence
}

func EquiperArme(p *Personnage, a Arme) error {
	p.ArmeEquipee = a
	p.Precision += a.Precision
	p.TauxCritique += a.TauxCritique
	// Si tu veux aussi gérer des bonus d'effet, c’est ici

	return nil
}

func EquiperArmure(p *Personnage, armures []Armure) error {
	if p.NiveauArmure >= len(armures) {
		return fmt.Errorf("NiveauArmure hors limites : %d", p.NiveauArmure)
	}
	a := armures[p.NiveauArmure]
	p.ArmureEquipee = a
	p.Armure += a.Defense
	p.ResistMag += a.Resistance
	p.PVMax += a.HP
	p.PV += a.HP
	return nil
}

func CoutAmelioration(niveau int) int {
	// Exemple simple : 5 roches au niveau 0, puis +5 à chaque niveau
	return 5 + niveau*5
}
func AmeliorerArme(p *Personnage, maxNiveau int) error {
	if p.NiveauArme >= maxNiveau-1 {
		return fmt.Errorf("arme déjà au niveau max")
	}
	cout := CoutAmelioration(p.NiveauArme)
	if p.Roches < cout {
		return fmt.Errorf("pas assez de roches pour améliorer l'arme (il faut %d, tu as %d)", cout, p.Roches)
	}
	p.Roches -= cout
	p.NiveauArme++
	return nil
}

func AmeliorerArmure(p *Personnage, maxNiveau int) error {
	if p.NiveauArmure >= maxNiveau-1 {
		return fmt.Errorf("armure déjà au niveau max")
	}
	cout := CoutAmelioration(p.NiveauArmure)
	if p.Roches < cout {
		return fmt.Errorf("pas assez de roches pour améliorer l'armure (il faut %d, tu as %d)", cout, p.Roches)
	}
	p.Roches -= cout
	p.NiveauArmure++
	return nil
}

var ArmesSoldat = []Arme{
	epeeBois,
	epeePierre,
	epeeFer,
	epeeOr,
	epeeDiamant,
	epeeNetherite,
}

var ArmesCRS = []Arme{
	matraqueStandard,
	matraqueFumigene,
	matraqueAntiEmeute,
	matraqueTelescopique,
}

var ArmesPyromane = []Arme{
	briquet,
	lanceFlamme,
	canonAFeu,
	volcanDeMagma,
}

var ArmesRobin = []Arme{
	arcBois,
	arbaleteLegere,
	arbaleteStandard,
	arbaleteVenimeuse,
}

var ArmesBoucher = []Arme{
	couteauCuisine,
	couteauBoucher,
	hacheoir,
	hacheDeGuerre,
}

var ArmesCroMagnon = []Arme{
	lancePierre,
	frondeRenforcee,
	lanceTribale,
	lanceMammouth,
}

var ArmesZeus = []Arme{
	etincelle,
	foudreMineure,
	foudreSombre,
	foudreDivine,
	foutreDeZeus,
}

var ArmesSamourai = []Arme{
	sabreBasique,
	katana,
	katanaShuriken,
	katanaLameCeleste,
}

var ArmesGandalf = []Arme{
	batonDeMage,
	batonArcanique,
	batonElementaire,
	batonGrandMage,
}

var ArmesSinge = []Arme{
	banane,
	bananierCombat,
	lanceBanane,
	bananeRoyale,
	bananeDivine,
}

var ArmesBambi = []Arme{
	arGrise,
	arVerte,
	arBleue,
	scarViolette,
	scarEnOr,
}

var ArmesPoseidon = []Arme{
	fourchetteDesMers,
	tridentDuMarais,
	tridentDesProfondeurs,
	tridentPoseidon,
}

var ArmuresSoldat = []Armure{
	{"Armure de Recrue", 10, 2, 10},
	{"Cuirasse du Sergent", 17, 5, 20},
	{"Armure de Général", 25, 8, 35},
	{"Armure du Maréchal", 34, 12, 50},
}
var ArmuresCRS = []Armure{
	{"Gilet Pare-Balles Standard", 12, 4, 15},
	{"Tenue Anti-Émeute", 20, 7, 25},
	{"Armure Blindée CRS", 30, 10, 40},
	{"Exo-Riot Intégrale", 40, 14, 60},
}
var ArmuresPyromane = []Armure{
	{"Veste Ignifugée", 6, 12, 5},
	{"Combinaison Thermique", 9, 18, 10},
	{"Tenue de Pyromancien", 12, 25, 20},
	{"Parure du Pyromancien Royal", 15, 33, 30},
}
var ArmuresRobin = []Armure{
	{"Tunique de Forêt", 8, 3, 10},
	{"Armure de Chasseur", 13, 6, 15},
	{"Cape de la Sylve", 18, 10, 25},
	{"Cape de l'Archéon", 23, 14, 35},
}
var ArmuresBoucher = []Armure{
	{"Tablier de Boucher", 10, 2, 10},
	{"Plastron Sanguinolent", 17, 5, 20},
	{"Armure du Massacreur", 24, 8, 35},
	{"Cuirasse du Carnassier", 32, 11, 50},
}
var ArmuresCroMagnon = []Armure{
	{"Peaux de Bête", 12, 2, 12},
	{"Tenue Tribale Renforcée", 18, 4, 22},
	{"Armure de Chasseur Primitif", 25, 7, 35},
	{"Armure du Chasseur Alpha", 33, 10, 50},
}
var ArmuresZeus = []Armure{
	{"Robe Électrique", 6, 15, 5},
	{"Tunique du Tonnerre", 9, 23, 10},
	{"Armure Divine de Zeus", 12, 32, 20},
	{"Panoplie Olympienne", 15, 42, 30},
}
var ArmuresSamourai = []Armure{
	{"Kimono de Combat", 10, 5, 8},
	{"Armure Légère de Samouraï", 16, 8, 15},
	{"Armure d’Élite Shogun", 22, 12, 25},
	{"Armure du Daimyo", 28, 16, 35},
}
var ArmuresGandalf = []Armure{
	{"Robe d’Apprenti", 6, 12, 5},
	{"Robe Arcanique", 10, 20, 10},
	{"Robe du Grand Mage", 14, 30, 20},
	{"Robe du Sage Éternel", 18, 40, 30},
}
var ArmuresSinge = []Armure{
	{"Peau de Banane", 6, 4, 5},
	{"Costume de Singe Ninja", 10, 7, 10},
	{"Armure Royale Simiesque", 15, 12, 18},
	{"Armure Mythique Simiesque", 20, 16, 28},
}
var ArmuresBambi = []Armure{
	{"Gilet de Chasseur", 10, 5, 10},
	{"Tactical Gear", 16, 9, 18},
	{"Exo-Armure de Combat", 24, 14, 30},
	{"Exo-Armure Tactique Mk II", 32, 19, 45},
}
var ArmuresPoseidon = []Armure{
	{"Cuirasse des Vagues", 15, 10, 20},
	{"Armure de l’Abysse", 22, 16, 30},
	{"Armure Royale de Poséidon", 30, 25, 45},
	{"Armure du Souverain des Mers", 40, 34, 65},
}
