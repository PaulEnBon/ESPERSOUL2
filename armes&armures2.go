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

// Retourne le template de base de la classe d'un personnage selon son nom
func getBaseTemplateByName(name string) (Personnage, bool) {
	switch name {
	case Steeve.Nom:
		return Steeve, true
	case CRS.Nom:
		return CRS, true
	case Pyromane.Nom:
		return Pyromane, true
	case RobinDesBois.Nom:
		return RobinDesBois, true
	case Boucher.Nom:
		return Boucher, true
	case CroMagnon.Nom:
		return CroMagnon, true
	case Zeus.Nom:
		return Zeus, true
	case Samourai.Nom:
		return Samourai, true
	case Gandalf.Nom:
		return Gandalf, true
	case Singe.Nom:
		return Singe, true
	case Bambi.Nom:
		return Bambi, true
	case Poseidon.Nom:
		return Poseidon, true
	case Erwann.Nom:
		return Erwann, true
	case Gabriel.Nom:
		return Gabriel, true
	case Vitaly.Nom:
		return Vitaly, true
	default:
		return Personnage{}, false
	}
}

// Ré-initialise les stats du perso à celles de sa classe puis ré-applique l'arme/armure équipées
func RecomputeFromBaseAndEquip(p *Personnage) {
	base, ok := getBaseTemplateByName(p.Nom)
	if !ok {
		return // Classe inconnue: ne rien faire pour éviter des états incohérents
	}

	oldPV := p.PV
	oldPVMax := p.PVMax

	// Remettre les stats de base de la classe
	p.PVMax = base.PVMax
	p.PV = base.PV // valeur provisoire; ajustée après équipement
	p.Armure = base.Armure
	p.ResistMag = base.ResistMag
	p.Precision = base.Precision
	p.TauxCritique = base.TauxCritique
	p.MultiplicateurCrit = base.MultiplicateurCrit

	// Réinitialiser l'équipement courant (structs)
	p.ArmeEquipee = Arme{}
	p.ArmureEquipee = Armure{}

	// Ré-appliquer l'armure selon le niveau actuel
	if p.NiveauArmure >= 0 && p.NiveauArmure < len(p.ArmuresDisponibles) {
		_ = EquiperArmure(p, p.ArmuresDisponibles)
	}
	// Ré-appliquer l'arme selon le niveau actuel
	if p.NiveauArme >= 0 && p.NiveauArme < len(p.ArmesDisponibles) {
		_ = EquiperArme(p, p.ArmesDisponibles[p.NiveauArme])
	}

	// Préserver le ratio de PV après la montée de stats
	if oldPVMax > 0 {
		ratio := float64(oldPV) / float64(oldPVMax)
		newPV := int(ratio * float64(p.PVMax))
		if newPV < 1 {
			newPV = 1
		}
		if newPV > p.PVMax {
			newPV = p.PVMax
		}
		p.PV = newPV
	} else {
		if p.PV > p.PVMax {
			p.PV = p.PVMax
		}
	}
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
	// Ancien comportement: on ajoutait seulement le bonus HP.
	// Nouveau: on considère qu'un changement d'armure correspond à un "repos" => PV plein.
	// On fixe donc PV au cap après recalcul.
	p.PV += a.HP
	if p.PV > p.PVMax { // sécurité (devrait toujours être vrai ici)
		p.PV = p.PVMax
	}
	// Heal complet explicite
	p.PV = p.PVMax
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
	// Ré-initialiser les stats de base puis ré-appliquer l'équipement au nouveau niveau
	RecomputeFromBaseAndEquip(p)
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
	// Ré-initialiser les stats de base puis ré-appliquer l'équipement au nouveau niveau
	RecomputeFromBaseAndEquip(p)
	// Heal complet après amélioration (logique de renforcement + repos)
	p.PV = p.PVMax
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

// Armes de Vitaly
var ArmesVitaly = []Arme{
	flashDeVodka,
	bouteilleDeVodka,
	griffeDOurs,
	apocalypseVodka,
}

// Armes de Gabriel
var ArmesGabriel = []Arme{
	vergeCeleste,
	lanceArchange,
	trompeteJugement,
	glaiveApocalypse,
}

// Armes de la classe Erwann
var ArmesErwann = []Arme{
	macMini,
	macAir,
	macPro,
	pcDuCDI,
}

var ArmuresSoldat = []Armure{
	{"Armure de Recrue", 8, 2, 8},
	{"Cuirasse du Sergent", 15, 5, 18}, // ~x1.8 déf / +125% HP vs tier1
	{"Armure de Général", 27, 9, 35},   // saut plus marqué
	{"Armure du Maréchal", 45, 14, 60}, // exponentiel final
}
var ArmuresCRS = []Armure{
	{"Gilet Pare-Balles Standard", 10, 4, 15},
	{"Tenue Anti-Émeute", 19, 7, 30},
	{"Armure Blindée CRS", 33, 11, 55},
	{"Exo-Riot Intégrale", 55, 16, 85},
}
var ArmuresPyromane = []Armure{
	{"Veste Ignifugée", 5, 10, 5},
	{"Combinaison Thermique", 8, 18, 12},
	{"Tenue de Pyromancien", 12, 30, 22},
	{"Parure du Pyromancien Royal", 18, 45, 38},
}
var ArmuresRobin = []Armure{
	{"Tunique de Forêt", 7, 3, 9},
	{"Armure de Chasseur", 12, 6, 18},
	{"Cape de la Sylve", 20, 10, 32},
	{"Cape de l'Archéon", 32, 15, 55},
}
var ArmuresBoucher = []Armure{
	{"Tablier de Boucher", 9, 2, 10},
	{"Plastron Sanguinolent", 16, 5, 22},
	{"Armure du Massacreur", 27, 9, 40},
	{"Cuirasse du Carnassier", 44, 13, 65},
}
var ArmuresCroMagnon = []Armure{
	{"Peaux de Bête", 10, 2, 12},
	{"Tenue Tribale Renforcée", 17, 4, 24},
	{"Armure de Chasseur Primitif", 29, 7, 42},
	{"Armure du Chasseur Alpha", 48, 11, 68},
}
var ArmuresZeus = []Armure{
	{"Robe Électrique", 5, 14, 5},
	{"Tunique du Tonnerre", 8, 24, 12},
	{"Armure Divine de Zeus", 12, 38, 22},
	{"Panoplie Olympienne", 18, 55, 38},
}
var ArmuresSamourai = []Armure{
	{"Kimono de Combat", 9, 5, 8},
	{"Armure Légère de Samouraï", 15, 8, 18},
	{"Armure d’Élite Shogun", 24, 12, 32},
	{"Armure du Daimyo", 38, 17, 52},
}
var ArmuresGandalf = []Armure{
	{"Robe d’Apprenti", 5, 11, 5},
	{"Robe Arcanique", 8, 21, 12},
	{"Robe du Grand Mage", 12, 34, 22},
	{"Robe du Sage Éternel", 18, 50, 38},
}
var ArmuresSinge = []Armure{
	{"Peau de Banane", 5, 4, 5},
	{"Costume de Singe Ninja", 9, 7, 12},
	{"Armure Royale Simiesque", 16, 13, 22},
	{"Armure Mythique Simiesque", 26, 18, 38},
}
var ArmuresBambi = []Armure{
	{"Gilet de Chasseur", 9, 5, 10},
	{"Tactical Gear", 15, 9, 20},
	{"Exo-Armure de Combat", 26, 15, 36},
	{"Exo-Armure Tactique Mk II", 42, 21, 58},
}
var ArmuresPoseidon = []Armure{
	{"Cuirasse des Vagues", 14, 10, 18},
	{"Armure de l’Abysse", 24, 17, 34},
	{"Armure Royale de Poséidon", 38, 27, 58},
	{"Armure du Souverain des Mers", 60, 40, 90},
}

// Armure Vitaly (Ensemble Adidas) progression (focus défense magique + critique élevé)
var ArmuresVitaly = []Armure{
	{"Ensemble Adidas Classique", 14, 18, 25},
	{"Ensemble Adidas Renforcé", 24, 30, 45},
	{"Ensemble Adidas Légendaire", 39, 44, 70},
	{"Ensemble Adidas Ultime", 62, 62, 105},
}

// Armures de Gabriel (orienté ultra tank + PV massifs)
var ArmuresGabriel = []Armure{
	{"Toge Bénie", 24, 24, 110},
	{"Plastron Séraphique", 40, 40, 190},
	{"Armure des Archontes", 66, 66, 300},
	{"Rempart Céleste", 100, 100, 470},
}

// Armures de la classe Erwann (équilibrées autour de tank techno)
var ArmuresErwann = []Armure{
	{"Coque Aluminium", 14, 11, 28},
	{"Châssis Optimisé", 24, 19, 52},
	{"Station de Travail", 38, 27, 86},
	{"Serveur Blindé", 58, 38, 130},
}
