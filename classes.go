package main

type ArmeEvolution struct {
	Versions []Arme
}

type ArmureEvolution struct {
	Versions []Armure
}

// Structure minimale du personnage (vous pouvez l'adapter)
type Personnage struct {
	Nom                string
	PV                 int
	PVMax              int
	Armure             int
	ResistMag          int
	Precision          float64
	TauxCritique       float64
	MultiplicateurCrit float64 // Nouveau champ : multiplicateur critique (ex: 2.0 = x2)
	EffetsActifs       []EffetActif
	ArmesDisponibles   []Arme
	ArmuresDisponibles []Armure
	NiveauArme         int // Commence à 0
	NiveauArmure       int // Commence à 0
	Roches             int // Ressource utilisée pour améliorer
	ArmeEquipee        Arme
	ArmureEquipee      Armure
	// ArtefactsEquipes devient un slice (taille max dynamique gérée par constante) pour simplifier l'évolution
	ArtefactsEquipes  []*Artefact // slots équipés (max 2 actuellement)
	ArtefactsPossedes []Artefact  // inventaire des artefacts possédés (illimité)
}

// Nombre maximum d'artefacts pouvant être équipés simultanément
const MaxArtefactsEquipes = 2

// ==========================
// LIGNÉE DES ÉPÉES
// ==========================

var Steeve = Personnage{
	Nom:                "Steeve",
	PV:                 95, // réduit (130 -> 95)
	PVMax:              95,
	Armure:             16, // réduit (24 -> 16)
	ResistMag:          7,  // réduit (10 -> 7)
	Precision:          0.85,
	TauxCritique:       0.10,
	MultiplicateurCrit: 1.6,
	ArmesDisponibles:   ArmesSoldat,
	ArmuresDisponibles: ArmuresSoldat,
	NiveauArme:         0,
	NiveauArmure:       0,
}

// ==========================
// LIGNÉE DES MATRAQUES
// ==========================
var CRS = Personnage{
	Nom:                "CRS",
	PV:                 120, // 170 -> 120
	PVMax:              120,
	Armure:             22, // 30 -> 22
	ResistMag:          11, // 15 -> 11
	Precision:          0.80,
	TauxCritique:       0.05,
	MultiplicateurCrit: 1.4,
	ArmesDisponibles: []Arme{
		matraqueStandard,
		matraqueFumigene,
		matraqueAntiEmeute,
		matraqueTelescopique,
	},
	ArmuresDisponibles: ArmuresCRS,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// LIGNÉE DES BRIQUETS ET ARMES FEU
// ==========================
var Pyromane = Personnage{
	Nom:                "Pyromane",
	PV:                 80, // 110 -> 80
	PVMax:              80,
	Armure:             5,  // 8 -> 5
	ResistMag:          12, // 16 -> 12
	Precision:          0.80,
	TauxCritique:       0.10,
	MultiplicateurCrit: 1.7,
	ArmesDisponibles: []Arme{
		briquet,
		lanceFlamme,
		canonAFeu,
		volcanDeMagma,
	},
	ArmuresDisponibles: ArmuresPyromane,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// LIGNÉE ARC & ARBALÈTE
// ==========================
var RobinDesBois = Personnage{
	Nom:                "Robin des Bois",
	PV:                 85, // 115 -> 85
	PVMax:              85,
	Armure:             9, // 12 -> 9
	ResistMag:          6, // 9 -> 6
	Precision:          0.90,
	TauxCritique:       0.20,
	MultiplicateurCrit: 1.9,
	ArmesDisponibles: []Arme{
		arcBois,
		arbaleteLegere,
		arbaleteStandard,
		arbaleteVenimeuse,
	},
	ArmuresDisponibles: ArmuresRobin,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// LIGNÉE COUTEAUX & HACHES
// ==========================
var Boucher = Personnage{
	Nom:                "Boucher",
	PV:                 105, // 145 -> 105
	PVMax:              105,
	Armure:             14, // 20 -> 14
	ResistMag:          6,  // 8 -> 6
	Precision:          0.85,
	TauxCritique:       0.20,
	MultiplicateurCrit: 1.85,
	ArmesDisponibles: []Arme{
		couteauCuisine,
		couteauBoucher,
		hacheoir,
		hacheDeGuerre,
	},
	ArmuresDisponibles: ArmuresBoucher,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// LIGNÉE LANCES & FRONDES
// ==========================
var CroMagnon = Personnage{
	Nom:                "Cro-Magnon",
	PV:                 115, // 160 -> 115
	PVMax:              115,
	Armure:             13, // 17 -> 13
	ResistMag:          7,  // 9 -> 7
	Precision:          0.85,
	TauxCritique:       0.25,
	MultiplicateurCrit: 1.7,
	ArmesDisponibles: []Arme{
		lancePierre,
		frondeRenforcee,
		lanceTribale,
		lanceMammouth,
	},
	ArmuresDisponibles: ArmuresCroMagnon,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// LIGNÉE MAGIQUE
// ==========================
var Zeus = Personnage{
	Nom:                "Zeus",
	PV:                 90, // 125 -> 90
	PVMax:              90,
	Armure:             7,  // 10 -> 7
	ResistMag:          26, // 36 -> 26
	Precision:          0.90,
	TauxCritique:       0.35,
	MultiplicateurCrit: 1.8,
	ArmesDisponibles: []Arme{
		etincelle,
		foudreMineure,
		foudreSombre,
		foudreDivine,
		foutreDeZeus,
	},
	ArmuresDisponibles: ArmuresZeus,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// LIGNÉE SABRES & KATANAS
// ==========================
var Samourai = Personnage{
	Nom:                "Samourai",
	PV:                 90, // 125 -> 90
	PVMax:              90,
	Armure:             12, // 17 -> 12
	ResistMag:          9,  // 13 -> 9
	Precision:          0.90,
	TauxCritique:       0.30,
	MultiplicateurCrit: 2.0,
	ArmesDisponibles: []Arme{
		sabreBasique,
		katana,
		katanaShuriken,
		katanaLameCeleste,
	},
	ArmuresDisponibles: ArmuresSamourai,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// LIGNÉE BÂTONS
// ==========================
var Gandalf = Personnage{
	Nom:                "Gandalf",
	PV:                 85, // 115 -> 85
	PVMax:              85,
	Armure:             7,  // 10 -> 7
	ResistMag:          23, // 32 -> 23
	Precision:          0.88,
	TauxCritique:       0.20,
	MultiplicateurCrit: 1.7,
	ArmesDisponibles: []Arme{
		batonDeMage,
		batonArcanique,
		batonElementaire,
		batonGrandMage,
	},
	ArmuresDisponibles: ArmuresGandalf,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// LIGNÉE BANANES
// ==========================
var Singe = Personnage{
	Nom:                "Singe",
	PV:                 80, // 110 -> 80
	PVMax:              80,
	Armure:             8, // 11 -> 8
	ResistMag:          9, // 13 -> 9
	Precision:          0.95,
	TauxCritique:       0.20,
	MultiplicateurCrit: 1.75,
	ArmesDisponibles: []Arme{
		banane,
		bananierCombat,
		lanceBanane,
		bananeRoyale,
		bananeDivine,
	},
	ArmuresDisponibles: ArmuresSinge,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// LIGNÉE AR / SCAR
// ==========================
var Bambi = Personnage{
	Nom:                "Bambi",
	PV:                 90, // 125 -> 90
	PVMax:              90,
	Armure:             12, // 17 -> 12
	ResistMag:          9,  // 13 -> 9
	Precision:          0.90,
	TauxCritique:       0.20,
	MultiplicateurCrit: 1.55,
	ArmesDisponibles: []Arme{
		arGrise,
		arVerte,
		arBleue,
		scarViolette,
		scarEnOr,
	},
	ArmuresDisponibles: ArmuresBambi,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// LIGNÉE TRIDENT / FOURCHETTE
// ==========================
var Poseidon = Personnage{
	Nom:                "Poséidon",
	PV:                 110, // 150 -> 110
	PVMax:              110,
	Armure:             20, // 28 -> 20
	ResistMag:          16, // 22 -> 16
	Precision:          0.90,
	TauxCritique:       0.25,
	MultiplicateurCrit: 1.75,
	ArmesDisponibles: []Arme{
		fourchetteDesMers,
		tridentDuMarais,
		tridentDesProfondeurs,
		tridentPoseidon,
	},
	ArmuresDisponibles: ArmuresPoseidon,
	ArtefactsEquipes:   make([]*Artefact, MaxArtefactsEquipes),
}

// ==========================
// CLASSE ERWANN (spécial)
// ==========================
var Erwann = Personnage{
	Nom:                "Erwann",
	PV:                 500,
	PVMax:              500,
	Armure:             30,
	ResistMag:          30,
	Precision:          0.10,
	TauxCritique:       0.10,
	MultiplicateurCrit: 1.5,
	ArmesDisponibles:   ArmesErwann,
	ArmuresDisponibles: ArmuresErwann,
}

// ==========================
// CLASSE GABRIEL (archange tank-magie)
// ==========================
var Gabriel = Personnage{
	Nom:                "Gabriel",
	PV:                 1000,
	PVMax:              1000,
	Armure:             100,
	ResistMag:          100,
	Precision:          0.80,
	TauxCritique:       0.30,
	MultiplicateurCrit: 2.0,
	ArmesDisponibles:   ArmesGabriel,
	ArmuresDisponibles: ArmuresGabriel,
}

// ==========================
// CLASSE VITALY (critique extrême + précision parfaite)
// ==========================
var Vitaly = Personnage{
	Nom:                "Vitaly",
	PV:                 200,
	PVMax:              200,
	Armure:             30,
	ResistMag:          50,
	Precision:          1.00,
	TauxCritique:       0.90,
	MultiplicateurCrit: 4.0,
	ArmesDisponibles:   ArmesVitaly,
	ArmuresDisponibles: ArmuresVitaly,
}

// AllClasses retourne la liste des classes/joueurs pré-définis.
func AllClasses() []Personnage {
	return []Personnage{
		Steeve,
		CRS,
		Pyromane,
		RobinDesBois,
		Boucher,
		CroMagnon,
		Zeus,
		Samourai,
		Gandalf,
		Singe,
		Bambi,
		Poseidon,
		Erwann,
		Gabriel,
		Vitaly,
	}
}

// AllClassNames retourne uniquement les noms des classes disponibles.
func AllClassNames() []string {
	classes := AllClasses()
	names := make([]string, 0, len(classes))
	for _, c := range classes {
		names = append(names, c.Nom)
	}
	return names
}
