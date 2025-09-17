package main

import (
	"math/rand"
)

// Système d'ennemis par paliers (tutoriel / early / mid / late)

type EnemyTier int

const (
	TierTutorial EnemyTier = iota
	TierEarly
	TierMid
	TierLate
)

// Template d'ennemi, basé sur la struct Personnage et une arme équipée
type EnemyTemplate struct {
	Name   string
	Base   Personnage
	Weapon Arme
}

// Crée un ennemi à partir d'un template, avec option « super » qui scale un peu les stats
func NewEnemyFromTemplate(t EnemyTemplate, isSuper bool) Personnage {
	// copie de base
	p := t.Base
	// équipe l'arme (ajoute la précision/critique de l'arme et définit les compétences)
	_ = EquiperArme(&p, t.Weapon)

	if isSuper {
		// Léger scaling pour les ennemis surpuissants
		p.PV = int(float64(p.PV) * 1.8)
		p.PVMax = int(float64(p.PVMax) * 1.8)
		p.Armure = int(float64(p.Armure) * 1.25)
		p.ResistMag = int(float64(p.ResistMag) * 1.25)
		p.Precision += 0.05
		p.TauxCritique += 0.05
		if p.MultiplicateurCrit < 1.75 {
			p.MultiplicateurCrit += 0.15
		}
	}
	return p
}

// =====================
// Définition des templates
// =====================

// — Tutoriel — très simple
var TutoRat = EnemyTemplate{
	Name: "Rat",
	Base: Personnage{
		Nom:                "Rat",
		PV:                 20,
		PVMax:              20,
		Armure:             2,
		ResistMag:          0,
		Precision:          0.80,
		TauxCritique:       0.05,
		MultiplicateurCrit: 1.4,
	},
	Weapon: epeeBois, // une morsure déguisée en petite attaque physique
}

var TutoSlime = EnemyTemplate{
	Name: "Gelée",
	Base: Personnage{
		Nom:                "Gelée",
		PV:                 26,
		PVMax:              26,
		Armure:             4,
		ResistMag:          2,
		Precision:          0.75,
		TauxCritique:       0.02,
		MultiplicateurCrit: 1.3,
	},
	Weapon: lancePierre,
}

// — Early game — facile
var EarlyBrigand = EnemyTemplate{
	Name: "Brigand",
	Base: Personnage{
		Nom:                "Brigand",
		PV:                 45,
		PVMax:              45,
		Armure:             6,
		ResistMag:          4,
		Precision:          0.82,
		TauxCritique:       0.10,
		MultiplicateurCrit: 1.5,
	},
	Weapon: couteauCuisine,
}

var EarlyArcher = EnemyTemplate{
	Name: "Archer",
	Base: Personnage{
		Nom:                "Archer",
		PV:                 40,
		PVMax:              40,
		Armure:             5,
		ResistMag:          3,
		Precision:          0.88,
		TauxCritique:       0.18,
		MultiplicateurCrit: 1.6,
	},
	Weapon: arcBois,
}

var EarlyPyro = EnemyTemplate{
	Name: "Apprenti Pyro",
	Base: Personnage{
		Nom:                "Apprenti Pyro",
		PV:                 38,
		PVMax:              38,
		Armure:             3,
		ResistMag:          8,
		Precision:          0.80,
		TauxCritique:       0.12,
		MultiplicateurCrit: 1.5,
	},
	Weapon: briquet,
}

// — Mid game — moyen
var MidChevalier = EnemyTemplate{
	Name: "Chevalier",
	Base: Personnage{
		Nom:                "Chevalier",
		PV:                 90,
		PVMax:              90,
		Armure:             20,
		ResistMag:          10,
		Precision:          0.86,
		TauxCritique:       0.12,
		MultiplicateurCrit: 1.6,
	},
	Weapon: epeeFer,
}

var MidBerserker = EnemyTemplate{
	Name: "Berserker",
	Base: Personnage{
		Nom:                "Berserker",
		PV:                 110,
		PVMax:              110,
		Armure:             14,
		ResistMag:          8,
		Precision:          0.80,
		TauxCritique:       0.20,
		MultiplicateurCrit: 1.75,
	},
	Weapon: hacheoir,
}

var MidMage = EnemyTemplate{
	Name: "Mage Sombre",
	Base: Personnage{
		Nom:                "Mage Sombre",
		PV:                 75,
		PVMax:              75,
		Armure:             8,
		ResistMag:          22,
		Precision:          0.88,
		TauxCritique:       0.22,
		MultiplicateurCrit: 1.7,
	},
	Weapon: foudreSombre,
}

// — Late game — dur
var LateSeigneurDemon = EnemyTemplate{
	Name: "Seigneur Démon",
	Base: Personnage{
		Nom:                "Seigneur Démon",
		PV:                 180,
		PVMax:              180,
		Armure:             28,
		ResistMag:          24,
		Precision:          0.92,
		TauxCritique:       0.28,
		MultiplicateurCrit: 1.9,
	},
	Weapon: epeeNetherite,
}

var LateArchimage = EnemyTemplate{
	Name: "Archimage",
	Base: Personnage{
		Nom:                "Archimage",
		PV:                 150,
		PVMax:              150,
		Armure:             14,
		ResistMag:          36,
		Precision:          0.94,
		TauxCritique:       0.35,
		MultiplicateurCrit: 2.0,
	},
	Weapon: foudreDivine,
}

var LateChampion = EnemyTemplate{
	Name: "Champion déchu",
	Base: Personnage{
		Nom:                "Champion déchu",
		PV:                 200,
		PVMax:              200,
		Armure:             32,
		ResistMag:          20,
		Precision:          0.90,
		TauxCritique:       0.30,
		MultiplicateurCrit: 1.85,
	},
	Weapon: katanaLameCeleste,
}

// Pools par tier
var tutorialPool = []EnemyTemplate{TutoRat, TutoSlime}
var earlyPool = []EnemyTemplate{EarlyBrigand, EarlyArcher, EarlyPyro}
var midPool = []EnemyTemplate{MidChevalier, MidBerserker, MidMage}
var latePool = []EnemyTemplate{LateSeigneurDemon, LateArchimage, LateChampion}

// Détermine le tier d'une salle
func tierForMap(mapName string) EnemyTier {
	switch mapName {
	case "salle1":
		return TierTutorial
	case "salle2", "salle3":
		return TierEarly
	case "salle4", "salle5", "salle6", "salle7":
		return TierMid
	case "salle8", "salle9", "salle10":
		return TierLate
	default:
		return TierEarly
	}
}

// Renvoie un ennemi aléatoire pour une salle donnée
func CreateRandomEnemyForMap(mapName string, isSuper bool) Personnage {
	tier := tierForMap(mapName)
	switch tier {
	case TierTutorial:
		t := tutorialPool[rand.Intn(len(tutorialPool))]
		return NewEnemyFromTemplate(t, isSuper)
	case TierEarly:
		t := earlyPool[rand.Intn(len(earlyPool))]
		return NewEnemyFromTemplate(t, isSuper)
	case TierMid:
		t := midPool[rand.Intn(len(midPool))]
		return NewEnemyFromTemplate(t, isSuper)
	case TierLate:
		t := latePool[rand.Intn(len(latePool))]
		return NewEnemyFromTemplate(t, isSuper)
	default:
		t := earlyPool[rand.Intn(len(earlyPool))]
		return NewEnemyFromTemplate(t, isSuper)
	}
}

// Ennemi de tutoriel garanti (utile pour le tout premier combat)
func CreateTutorialEnemy() Personnage {
	return NewEnemyFromTemplate(TutoRat, false)
}
