package main

// Armes dédiées aux ennemis (pour éviter de réutiliser celles des joueurs)
var (
	// Tutoriel
	armeMorsureRat = Arme{
		Nom: "Morsure de Rat", DegatsPhysiques: 8, DegatsMagiques: 0, Precision: 0.80, TauxCritique: 0.05, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Griffure", Description: "Petite éraflure.", Degats: 6, CoutMana: 0, Type: "physique"},
			{Nom: "Morsure Infectée", Description: "Applique poison léger.", Degats: 4, CoutMana: 0, Type: "physique", TypeEffet: "Poison", Puissance: 1},
		},
	}
	armeProjectionSlime = Arme{
		Nom: "Projection Visqueuse", DegatsPhysiques: 4, DegatsMagiques: 10, Precision: 0.75, TauxCritique: 0.02, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Éclaboussure", Description: "Dégâts collants.", Degats: 9, CoutMana: 2, Type: "magique"},
			{Nom: "Gelée Visqueuse", Description: "Rend moins précis.", Degats: 0, CoutMana: 3, Type: "magique", TypeEffet: "Nébulation", Puissance: 1},
		},
	}

	// Early
	armeDagueBrigand = Arme{
		Nom: "Dague de Brigand", DegatsPhysiques: 14, DegatsMagiques: 0, Precision: 0.86, TauxCritique: 0.15, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Entaille Rapide", Description: "Saignement léger.", Degats: 10, CoutMana: 0, Type: "physique", TypeEffet: "Saignement", Puissance: 1},
			{Nom: "Poussière aux Yeux", Description: "Aveugle légèrement.", Degats: 0, CoutMana: 0, Type: "physique", TypeEffet: "Nébulation", Puissance: 1},
		},
	}
	armeArcRecu = Arme{
		Nom: "Arc Recu d'Apprenti", DegatsPhysiques: 16, DegatsMagiques: 0, Precision: 0.90, TauxCritique: 0.18, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Tir Précis", Description: "Un tir basique.", Degats: 14, CoutMana: 0, Type: "physique"},
			{Nom: "Tir Étourdissant", Description: "Faible étourdissement.", Degats: 10, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 1},
		},
	}
	armeBatonPyro = Arme{
		Nom: "Bâton Chaufferette", DegatsPhysiques: 0, DegatsMagiques: 22, Precision: 0.82, TauxCritique: 0.12, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Flammèche", Description: "Brûlure légère.", Degats: 14, CoutMana: 3, Type: "magique", TypeEffet: "Brûlure", Puissance: 1},
			{Nom: "Brasero", Description: "Brûlure continue faible.", Degats: 8, CoutMana: 2, Type: "magique", TypeEffet: "Brûlure", Puissance: 1},
		},
	}

	// Mid
	armeLameChevalier = Arme{
		Nom: "Lame de Chevalier", DegatsPhysiques: 26, DegatsMagiques: 0, Precision: 0.88, TauxCritique: 0.12, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Estoc", Description: "Coup maîtrisé.", Degats: 22, CoutMana: 0, Type: "physique"},
			{Nom: "Brise-Garde", Description: "Réduit l'armure.", Degats: 12, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 1},
		},
	}
	armeHacheBerserker = Arme{
		Nom: "Hache du Berserker", DegatsPhysiques: 34, DegatsMagiques: 0, Precision: 0.80, TauxCritique: 0.22, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Furie", Description: "Augmente les dégâts.", Degats: 0, CoutMana: 3, Type: "physique", TypeEffet: "Augmentation de Dégâts", Puissance: 2},
			{Nom: "Coup Sauvage", Description: "Gros dégâts bruts.", Degats: 28, CoutMana: 0, Type: "physique"},
		},
	}
	armeSceptreSombre = Arme{
		Nom: "Sceptre Sombre", DegatsPhysiques: 0, DegatsMagiques: 48, Precision: 0.90, TauxCritique: 0.22, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Ombre Fulgurante", Description: "Brise la résistance magique.", Degats: 42, CoutMana: 8, Type: "magique", TypeEffet: "Brise-Armure Magique", Puissance: 2},
			{Nom: "Nuit Affaiblissante", Description: "Affaiblit l'ennemi.", Degats: 0, CoutMana: 6, Type: "magique", TypeEffet: "Affaiblissement", Puissance: 2},
		},
	}

	// Late
	armeLameDemone = Arme{
		Nom: "Lame Démoniaque", DegatsPhysiques: 46, DegatsMagiques: 6, Precision: 0.90, TauxCritique: 0.30, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Tranche Infernale", Description: "Frappe puissante.", Degats: 40, CoutMana: 0, Type: "physique"},
			{Nom: "Entaille Ardente", Description: "Brûlure modérée.", Degats: 22, CoutMana: 4, Type: "magique", TypeEffet: "Brûlure", Puissance: 2},
		},
	}
	armeBatonArchimage = Arme{
		Nom: "Bâton de l'Archimage", DegatsPhysiques: 0, DegatsMagiques: 95, Precision: 0.95, TauxCritique: 0.35, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Choc Arcanique", Description: "Dégâts magiques élevés.", Degats: 90, CoutMana: 12, Type: "magique"},
			{Nom: "Bulle Mystique", Description: "Aveugle fortement.", Degats: 0, CoutMana: 10, Type: "magique", TypeEffet: "Nébulation", Puissance: 3},
		},
	}
	armeLameDuChampion = Arme{
		Nom: "Lame du Champion", DegatsPhysiques: 56, DegatsMagiques: 8, Precision: 0.92, TauxCritique: 0.40, Durabilite: 999,
		Competences: []Competence{
			{Nom: "Juge Final", Description: "Coup critique fréquent.", Degats: 48, CoutMana: 0, Type: "physique"},
			{Nom: "Lame Brise-Armure", Description: "Réduction d'armure importante.", Degats: 20, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 3},
		},
	}
)

// Armures dédiées aux ennemis
var (
	armureRat   = Armure{Nom: "Peau de Rat", Defense: 2, Resistance: 0, HP: 5}
	armureSlime = Armure{Nom: "Gel Visqueux", Defense: 3, Resistance: 2, HP: 6}

	armureBrigand      = Armure{Nom: "Cuir de Brigand", Defense: 6, Resistance: 2, HP: 10}
	armureArcher       = Armure{Nom: "Tunique d'Archer", Defense: 5, Resistance: 3, HP: 8}
	armureApprentiPyro = Armure{Nom: "Robe de Novice", Defense: 3, Resistance: 8, HP: 8}

	armureChevalier  = Armure{Nom: "Plastron du Chevalier", Defense: 18, Resistance: 8, HP: 18}
	armureBerserker  = Armure{Nom: "Peaux du Berserker", Defense: 12, Resistance: 6, HP: 22}
	armureMageSombre = Armure{Nom: "Robe Sombre", Defense: 6, Resistance: 18, HP: 14}

	armureSeigneurDemon = Armure{Nom: "Carapace Démoniaque", Defense: 28, Resistance: 22, HP: 30}
	armureArchimage     = Armure{Nom: "Parure de l'Archimage", Defense: 10, Resistance: 32, HP: 24}
	armureChampion      = Armure{Nom: "Cuirasse du Champion", Defense: 30, Resistance: 18, HP: 36}
)
