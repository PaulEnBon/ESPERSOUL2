package main

type Competence struct {
	Nom         string
	Description string
	Degats      int
	CoutMana    int
	Type        string // "physique" ou "magique"
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

var scarEnOr = Arme{
	Nom:             "SCAR en Or",
	DegatsPhysiques: 40,
	DegatsMagiques:  0,
	Precision:       0.95,
	TauxCritique:    0.35,
	Durabilite:      120,
	Competences: []Competence{
		{
			Nom:         "Rafale précise",
			Description: "Tire plusieurs balles avec précision.",
			Degats:      15,
			CoutMana:    0,
			Type:        "physique",
		},
		{
			Nom:         "Tir critique",
			Description: "Tir chargé infligeant un critique assuré.",
			Degats:      35,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}

var sabre = Arme{
	Nom:             "Sabre",
	DegatsPhysiques: 16,
	DegatsMagiques:  0,
	Precision:       0.90,
	TauxCritique:    0.25,
	Durabilite:      60,
	Competences: []Competence{
		{
			Nom:         "Frappe éclair",
			Description: "Attaque rapide à haute précision.",
			Degats:      12,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}

var trident = Arme{
	Nom:             "Trident",
	DegatsPhysiques: 20,
	DegatsMagiques:  5,
	Precision:       0.85,
	TauxCritique:    0.2,
	Durabilite:      80,
	Competences: []Competence{
		{
			Nom:         "Lancer tranchant",
			Description: "Projette le trident sur l'ennemi.",
			Degats:      25,
			CoutMana:    5,
			Type:        "physique",
		},
		{
			Nom:         "Frappe aquatique",
			Description: "Bonus si l'ennemi est mouillé.",
			Degats:      18,
			CoutMana:    4,
			Type:        "magique",
		},
	},
}

var batonDeMage = Arme{
	Nom:             "Bâton de Mage",
	DegatsPhysiques: 2,
	DegatsMagiques:  22,
	Precision:       0.88,
	TauxCritique:    0.15,
	Durabilite:      50,
	Competences: []Competence{
		{
			Nom:         "Rayon arcanique",
			Description: "Inflige des dégâts purs magiques.",
			Degats:      20,
			CoutMana:    8,
			Type:        "magique",
		},
	},
}

var foudreDeZeus = Arme{
	Nom:             "Foudre de Zeus",
	DegatsPhysiques: 5,
	DegatsMagiques:  40,
	Precision:       0.95,
	TauxCritique:    0.4,
	Durabilite:      30,
	Competences: []Competence{
		{
			Nom:         "Éclair divin",
			Description: "Inflige de lourds dégâts à un ennemi.",
			Degats:      50,
			CoutMana:    15,
			Type:        "magique",
		},
	},
}

var lanceFlamme = Arme{
	Nom:             "Lance-flamme",
	DegatsPhysiques: 0,
	DegatsMagiques:  28,
	Precision:       0.80,
	TauxCritique:    0.10,
	Durabilite:      40,
	Competences: []Competence{
		{
			Nom:         "Jet de feu",
			Description: "Brûle tous les ennemis en ligne.",
			Degats:      25,
			CoutMana:    10,
			Type:        "magique",
		},
	},
}

var potionDegats = Arme{
	Nom:             "Potion de Dégâts",
	DegatsPhysiques: 0,
	DegatsMagiques:  30,
	Precision:       1.0,
	TauxCritique:    0.0,
	Durabilite:      1, // usage unique
	Competences: []Competence{
		{
			Nom:         "Lancer explosif",
			Description: "Explose au contact, infligeant des dégâts magiques.",
			Degats:      30,
			CoutMana:    0,
			Type:        "magique",
		},
	},
}

var epeeBois = Arme{
	Nom:             "Épée en Bois",
	DegatsPhysiques: 12,
	DegatsMagiques:  0,
	Precision:       0.82,
	TauxCritique:    0.05,
	Durabilite:      35,
	Competences: []Competence{
		{
			Nom:         "Coup d’Apprenti",
			Description: "Un coup maladroit mais sans danger.",
			Degats:      10,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}

var epeePierre = Arme{
	Nom:             "Épée en Pierre",
	DegatsPhysiques: 18,
	DegatsMagiques:  0,
	Precision:       0.83,
	TauxCritique:    0.10,
	Durabilite:      55,
	Competences: []Competence{
		{
			Nom:         "Coup Solide",
			Description: "Un vrai outil de baston, pas juste un bâton.",
			Degats:      15,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}

var epeeFer = Arme{
	Nom:             "Épée en Fer",
	DegatsPhysiques: 24,
	DegatsMagiques:  0,
	Precision:       0.84,
	TauxCritique:    0.15,
	Durabilite:      75,
	Competences: []Competence{
		{
			Nom:         "Tranchant Métallique",
			Description: "Un coup net et précis.",
			Degats:      21,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}

var epeeOr = Arme{
	Nom:             "Épée en Or",
	DegatsPhysiques: 28,
	DegatsMagiques:  0,
	Precision:       0.85,
	TauxCritique:    0.30,
	Durabilite:      45,
	Competences: []Competence{
		{
			Nom:         "Frappe Brillante",
			Description: "Coup rapide avec plus de chance de critique.",
			Degats:      24,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}

var epeeDiamant = Arme{
	Nom:             "Épée en Diamant",
	DegatsPhysiques: 35,
	DegatsMagiques:  0,
	Precision:       0.86,
	TauxCritique:    0.20,
	Durabilite:      100,
	Competences: []Competence{
		{
			Nom:         "Frappe de Maître",
			Description: "Une attaque puissante et bien affûtée.",
			Degats:      30,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}

var epeeNetherite = Arme{
	Nom:             "Épée en Netherite",
	DegatsPhysiques: 42,
	DegatsMagiques:  5,
	Precision:       0.87,
	TauxCritique:    0.25,
	Durabilite:      150,
	Competences: []Competence{
		{
			Nom:         "Coup Infernal",
			Description: "Une frappe écrasante imprégnée d'énergie obscure.",
			Degats:      38,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}

var katana = Arme{
	Nom:             "Katana",
	DegatsPhysiques: 10,
	DegatsMagiques:  0,
	Precision:       0.90,
	TauxCritique:    0.40,
	Durabilite:      60,
	Competences: []Competence{
		{
			Nom:         "Taillade",
			Description: "Une coupe rapide et précise.",
			Degats:      10,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}

var hacheoir = Arme{
	Nom:             "Hacheoir",
	DegatsPhysiques: 18,
	DegatsMagiques:  0,
	Precision:       0.75,
	TauxCritique:    0.15,
	Durabilite:      70,
	Competences: []Competence{
		{
			Nom:         "Découpe Brutale",
			Description: "Un coup puissant, lent mais destructeur.",
			Degats:      25,
			CoutMana:    0,
			Type:        "physique",
		},
		{
			Nom:         "Frappe en Croix",
			Description: "Deux coups rapides qui percent l'armure.",
			Degats:      15,
			CoutMana:    5,
			Type:        "physique",
		},
	},
}

var matraque = Arme{
	Nom:             "Matraque",
	DegatsPhysiques: 12,
	DegatsMagiques:  0,
	Precision:       0.95,
	TauxCritique:    0.10,
	Durabilite:      80,
	Competences: []Competence{
		{
			Nom:         "Coup Assommant",
			Description: "Chance d'étourdir l'adversaire.",
			Degats:      10,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}

var briquet = Arme{
	Nom:             "Briquet",
	DegatsPhysiques: 2,
	DegatsMagiques:  15,
	Precision:       0.80,
	TauxCritique:    0.25,
	Durabilite:      40,
	Competences: []Competence{
		{
			Nom:         "Petite Flamme",
			Description: "Enflamme légèrement l’ennemi.",
			Degats:      12,
			CoutMana:    3,
			Type:        "magique",
		},
		{
			Nom:         "Explosion de Poche",
			Description: "Crée une étincelle instable.",
			Degats:      20,
			CoutMana:    7,
			Type:        "magique",
		},
	},
}

var lancePierre = Arme{
	Nom:             "Lance-pierre",
	DegatsPhysiques: 6,
	DegatsMagiques:  0,
	Precision:       0.92,
	TauxCritique:    0.30,
	Durabilite:      50,
	Competences: []Competence{
		{
			Nom:         "Jet Rapide",
			Description: "Un petit caillou bien placé.",
			Degats:      8,
			CoutMana:    0,
			Type:        "physique",
		},
		{
			Nom:         "Caillou Vicieux",
			Description: "Touche des points sensibles.",
			Degats:      12,
			CoutMana:    2,
			Type:        "physique",
		},
	},
}

var banane = Arme{
	Nom:             "Banane",
	DegatsPhysiques: 4,
	DegatsMagiques:  2,
	Precision:       0.99,
	TauxCritique:    0.05,
	Durabilite:      20,
	Competences: []Competence{
		{
			Nom:         "Glissade Fatale",
			Description: "Fait tomber l'ennemi (chance d'étourdissement).",
			Degats:      3,
			CoutMana:    0,
			Type:        "physique",
		},
		{
			Nom:         "Écrasement Mou",
			Description: "Tactique non conventionnelle mais efficace.",
			Degats:      6,
			CoutMana:    0,
			Type:        "physique",
		},
	},
}
