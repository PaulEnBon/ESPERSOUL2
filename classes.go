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
	PV:                 130,
	PVMax:              130,
	Armure:             24,
	ResistMag:          10,
	Precision:          0.85,
	TauxCritique:       0.10,
	MultiplicateurCrit: 1.6,
	ArmesDisponibles:   ArmesSoldat, // 👈 ici tu dois avoir toutes les armes
	ArmuresDisponibles: ArmuresSoldat,
	NiveauArme:         0,
	NiveauArmure:       0,
}

// ==========================
// LIGNÉE DES MATRAQUES
// ==========================
var CRS = Personnage{
	Nom:                "CRS",
	PV:                 170,
	PVMax:              170,
	Armure:             30,
	ResistMag:          15,
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
	PV:                 110,
	PVMax:              110,
	Armure:             8,
	ResistMag:          16,
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
	PV:                 115,
	PVMax:              115,
	Armure:             12,
	ResistMag:          9,
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
	PV:                 145,
	PVMax:              145,
	Armure:             20,
	ResistMag:          8,
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
	PV:                 160,
	PVMax:              160,
	Armure:             17,
	ResistMag:          9,
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
	PV:                 125,
	PVMax:              125,
	Armure:             10,
	ResistMag:          36,
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
	PV:                 125,
	PVMax:              125,
	Armure:             17,
	ResistMag:          13,
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
	PV:                 115,
	PVMax:              115,
	Armure:             10,
	ResistMag:          32,
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
	PV:                 110,
	PVMax:              110,
	Armure:             11,
	ResistMag:          13,
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
	PV:                 125,
	PVMax:              125,
	Armure:             17,
	ResistMag:          13,
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
	PV:                 150,
	PVMax:              150,
	Armure:             28,
	ResistMag:          22,
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
