package main

// --- Lignée des Épées (exemples avec effets progressifs) ---
var epeeBois = Arme{
	Nom: "Épée en Bois", DegatsPhysiques: 12, DegatsMagiques: 0, Precision: 0.82, TauxCritique: 0.05, Durabilite: 35,
	Competences: []Competence{
		{Nom: "Coup d'Apprenti", Description: "Un coup maladroit mais sans danger.", Degats: 10, CoutMana: 0, Type: "physique"},
	},
}

var epeePierre = Arme{
	Nom: "Épée en Pierre", DegatsPhysiques: 18, DegatsMagiques: 0, Precision: 0.83, TauxCritique: 0.10, Durabilite: 55,
	Competences: []Competence{
		{Nom: "Coup Solide", Description: "Un vrai outil de baston.", Degats: 15, CoutMana: 0, Type: "physique"},
	},
}

var epeeFer = Arme{
	Nom: "Épée en Fer", DegatsPhysiques: 24, DegatsMagiques: 0, Precision: 0.84, TauxCritique: 0.15, Durabilite: 75,
	Competences: []Competence{
		{Nom: "Tranchant Métallique", Description: "Un coup net et précis.", Degats: 21, CoutMana: 0, Type: "physique"},
		{Nom: "Entaille", Description: "Provoque un léger saignement.", Degats: 15, CoutMana: 2, Type: "physique", TypeEffet: "Saignement", Puissance: 1},
	},
}

var epeeOr = Arme{
	Nom: "Épée en Or", DegatsPhysiques: 28, DegatsMagiques: 0, Precision: 0.85, TauxCritique: 0.25, Durabilite: 50,
	Competences: []Competence{
		{Nom: "Frappe Brillante", Description: "Coup rapide avec chance critique accrue.", Degats: 24, CoutMana: 0, Type: "physique"},
		{Nom: "Taillade Dorée", Description: "Provoque un saignement modéré.", Degats: 20, CoutMana: 3, Type: "physique", TypeEffet: "Saignement", Puissance: 2},
	},
}

var epeeDiamant = Arme{
	Nom: "Épée en Diamant", DegatsPhysiques: 35, DegatsMagiques: 0, Precision: 0.86, TauxCritique: 0.30, Durabilite: 100,
	Competences: []Competence{
		{Nom: "Frappe de Maître", Description: "Attaque puissante et bien affûtée.", Degats: 30, CoutMana: 0, Type: "physique"},
		{Nom: "Lame Diamantée", Description: "Brise l'armure ennemie.", Degats: 25, CoutMana: 4, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 3},
		{Nom: "Hémorragie", Description: "Saignement important.", Degats: 22, CoutMana: 5, Type: "physique", TypeEffet: "Saignement", Puissance: 3},
	},
}

var epeeNetherite = Arme{
	Nom: "Épée en Netherite", DegatsPhysiques: 42, DegatsMagiques: 5, Precision: 0.87, TauxCritique: 0.35, Durabilite: 150,
	Competences: []Competence{
		{Nom: "Coup Infernal", Description: "Frappe écrasante imprégnée d'énergie obscure.", Degats: 38, CoutMana: 0, Type: "physique"},
		{Nom: "Destruction d'Armure", Description: "Détruit l'armure ennemie.", Degats: 30, CoutMana: 6, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 5},
		{Nom: "Saignée Mortelle", Description: "Hémorragie massive.", Degats: 28, CoutMana: 7, Type: "physique", TypeEffet: "Saignement", Puissance: 5},
	},
}

// --- Lignée des Matraques avec progression ---
var matraqueStandard = Arme{
	Nom: "Matraque Standard", DegatsPhysiques: 12, DegatsMagiques: 0, Precision: 0.95, TauxCritique: 0.10, Durabilite: 80,
	Competences: []Competence{
		{Nom: "Coup Assommant", Description: "Faible chance d'étourdir.", Degats: 10, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 1},
	},
}

var matraqueFumigene = Arme{
	Nom: "Matraque Fumigène", DegatsPhysiques: 14, DegatsMagiques: 0, Precision: 0.93, TauxCritique: 0.12, Durabilite: 85,
	Competences: []Competence{
		{Nom: "Coup Assommant", Description: "Chance modérée d'étourdir.", Degats: 12, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 2},
		{Nom: "Fumigène", Description: "Nuage de fumée léger.", Degats: 0, CoutMana: 5, Type: "magique", TypeEffet: "Nébulation", Puissance: 1},
	},
}

var matraqueAntiEmeute = Arme{
	Nom: "Matraque Anti-Émeute", DegatsPhysiques: 18, DegatsMagiques: 0, Precision: 0.91, TauxCritique: 0.15, Durabilite: 100,
	Competences: []Competence{
		{Nom: "Coup Assommant", Description: "Bonne chance d'étourdir.", Degats: 14, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 3},
		{Nom: "Fumigène Puissant", Description: "Nuage de fumée dense.", Degats: 0, CoutMana: 8, Type: "magique", TypeEffet: "Nébulation", Puissance: 3},
	},
}

var matraqueTelescopique = Arme{
	Nom: "Matraque Télescopique", DegatsPhysiques: 22, DegatsMagiques: 0, Precision: 0.93, TauxCritique: 0.18, Durabilite: 110,
	Competences: []Competence{
		{Nom: "Coup Assommant", Description: "Forte chance d'étourdir.", Degats: 16, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 4},
		{Nom: "Fumigène Amélioré", Description: "Nuage de fumée très dense.", Degats: 0, CoutMana: 10, Type: "magique", TypeEffet: "Nébulation", Puissance: 4},
		{Nom: "Impact de Force", Description: "Réduit fortement le taux critique.", Degats: 0, CoutMana: 5, Type: "physique", TypeEffet: "Défavorisation", Puissance: 3},
	},
}

// --- Lignée des armes à feu avec progression ---
var briquet = Arme{
	Nom: "Briquet", DegatsPhysiques: 2, DegatsMagiques: 15, Precision: 0.80, TauxCritique: 0.25, Durabilite: 40,
	Competences: []Competence{
		{Nom: "Petite Flamme", Description: "Brûlure légère.", Degats: 12, CoutMana: 3, Type: "magique", TypeEffet: "Brûlure", Puissance: 1},
		{Nom: "Explosion de Poche", Description: "Étincelle instable.", Degats: 20, CoutMana: 7, Type: "magique"},
	},
}

var lanceFlamme = Arme{
	Nom: "Lance-flamme", DegatsPhysiques: 0, DegatsMagiques: 28, Precision: 0.80, TauxCritique: 0.10, Durabilite: 40,
	Competences: []Competence{
		{Nom: "Jet de feu", Description: "Brûlure modérée.", Degats: 25, CoutMana: 10, Type: "magique", TypeEffet: "Brûlure", Puissance: 2},
	},
}

var canonAFeu = Arme{
	Nom: "Canon à Feu", DegatsPhysiques: 0, DegatsMagiques: 35, Precision: 0.85, TauxCritique: 0.15, Durabilite: 60,
	Competences: []Competence{
		{Nom: "Explosion en chaîne", Description: "Brûlure importante.", Degats: 20, CoutMana: 12, Type: "magique", TypeEffet: "Brûlure", Puissance: 3},
	},
}

var volcanDeMagma = Arme{
	Nom: "Volcan de Magma", DegatsPhysiques: 0, DegatsMagiques: 50, Precision: 0.90, TauxCritique: 0.25, Durabilite: 50,
	Competences: []Competence{
		{Nom: "Éruption de Magma", Description: "Brûlure dévastatrice.", Degats: 35, CoutMana: 15, Type: "magique", TypeEffet: "Brûlure", Puissance: 5},
		{Nom: "Fumée Enflammée", Description: "Réduit fortement la précision.", Degats: 0, CoutMana: 10, Type: "magique", TypeEffet: "Nébulation", Puissance: 4},
	},
}

// --- Lignée Arc & Arbalète avec progression ---
var arcBois = Arme{
	Nom: "Arc en Bois", DegatsPhysiques: 12, DegatsMagiques: 0, Precision: 0.85, TauxCritique: 0.15, Durabilite: 50,
	Competences: []Competence{
		{Nom: "Tir Simple", Description: "Un tir de base.", Degats: 12, CoutMana: 0, Type: "physique"},
	},
}

var arbaleteLegere = Arme{
	Nom: "Arbalète Légère", DegatsPhysiques: 20, DegatsMagiques: 0, Precision: 0.88, TauxCritique: 0.20, Durabilite: 70,
	Competences: []Competence{
		{Nom: "Tir Précis", Description: "Faible chance d'étourdir.", Degats: 20, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 1},
	},
}

var arbaleteStandard = Arme{
	Nom: "Arbalète Standard", DegatsPhysiques: 28, DegatsMagiques: 0, Precision: 0.90, TauxCritique: 0.25, Durabilite: 80,
	Competences: []Competence{
		{Nom: "Tir Puissant", Description: "Un tir concentré.", Degats: 28, CoutMana: 0, Type: "physique"},
		{Nom: "Tir Assommant", Description: "Chance modérée d'étourdir.", Degats: 20, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 2},
	},
}

var arbaleteVenimeuse = Arme{
	Nom: "Arbalète de Maître", DegatsPhysiques: 32, DegatsMagiques: 0, Precision: 0.92, TauxCritique: 0.30, Durabilite: 90,
	Competences: []Competence{
		{Nom: "Tir de Ronces Venimeuses", Description: "Poison puissant.", Degats: 32, CoutMana: 0, Type: "physique", TypeEffet: "Poison", Puissance: 4},
		{Nom: "Tir Assommant", Description: "Forte chance d'étourdir.", Degats: 20, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 3},
	},
}

// --- Lignée Couteaux & Haches avec progression ---
var couteauCuisine = Arme{
	Nom: "Couteau de Cuisine", DegatsPhysiques: 8, DegatsMagiques: 0, Precision: 0.92, TauxCritique: 0.15, Durabilite: 50,
	Competences: []Competence{
		{Nom: "Tranche Rapide", Description: "Saignement léger.", Degats: 8, CoutMana: 0, Type: "physique", TypeEffet: "Saignement", Puissance: 1},
	},
}

var couteauBoucher = Arme{
	Nom: "Couteau de Boucher", DegatsPhysiques: 14, DegatsMagiques: 0, Precision: 0.88, TauxCritique: 0.20, Durabilite: 60,
	Competences: []Competence{
		{Nom: "Coup Tranchant", Description: "Affaiblit l'ennemi.", Degats: 14, CoutMana: 0, Type: "physique", TypeEffet: "Affaiblissement", Puissance: 2},
	},
}

var hacheoir = Arme{
	Nom: "Hacheoir", DegatsPhysiques: 25, DegatsMagiques: 0, Precision: 0.75, TauxCritique: 0.15, Durabilite: 80,
	Competences: []Competence{
		{Nom: "Découpe Brutale", Description: "Brise l'armure modérément.", Degats: 25, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 2},
		{Nom: "Frappe en Croix", Description: "Saignement modéré.", Degats: 15, CoutMana: 5, Type: "physique", TypeEffet: "Saignement", Puissance: 2},
	},
}

var hacheDeGuerre = Arme{
	Nom: "Hache de Guerre", DegatsPhysiques: 35, DegatsMagiques: 0, Precision: 0.78, TauxCritique: 0.20, Durabilite: 120,
	Competences: []Competence{
		{Nom: "Coup Dévastateur", Description: "Détruit l'armure.", Degats: 40, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 4},
		{Nom: "Tourbillon Sanglant", Description: "Hémorragie sévère.", Degats: 22, CoutMana: 8, Type: "physique", TypeEffet: "Saignement", Puissance: 4},
	},
}

// --- Lignée Lances & Frondes avec progression ---
var lancePierre = Arme{
	Nom: "Lance-pierre", DegatsPhysiques: 6, DegatsMagiques: 0, Precision: 0.92, TauxCritique: 0.30, Durabilite: 50,
	Competences: []Competence{
		{Nom: "Jet Rapide", Description: "Petit caillou bien placé.", Degats: 8, CoutMana: 0, Type: "physique"},
		{Nom: "Caillou Vicieux", Description: "Faible chance d'étourdir.", Degats: 12, CoutMana: 2, Type: "physique", TypeEffet: "Étourdissement", Puissance: 1},
	},
}

var frondeRenforcee = Arme{
	Nom: "Fronde Renforcée", DegatsPhysiques: 12, DegatsMagiques: 0, Precision: 0.90, TauxCritique: 0.25, Durabilite: 65,
	Competences: []Competence{
		{Nom: "Projectile Lourd", Description: "Brise légèrement l'armure.", Degats: 16, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 1},
		{Nom: "Pluie de Cailloux", Description: "Ciblage multiple.", Degats: 6, CoutMana: 4, Type: "physique"},
	},
}

var lanceTribale = Arme{
	Nom: "Lance Tribale", DegatsPhysiques: 22, DegatsMagiques: 0, Precision: 0.88, TauxCritique: 0.35, Durabilite: 85,
	Competences: []Competence{
		{Nom: "Lancer de Guerre", Description: "Saignement modéré.", Degats: 25, CoutMana: 0, Type: "physique", TypeEffet: "Saignement", Puissance: 2},
		{Nom: "Appel Tribal", Description: "Boost modéré des dégâts.", Degats: 0, CoutMana: 3, Type: "physique", TypeEffet: "Augmentation de Dégâts", Puissance: 2},
	},
}

var lanceMammouth = Arme{
	Nom: "Lance de Mammouth", DegatsPhysiques: 35, DegatsMagiques: 5, Precision: 0.85, TauxCritique: 0.25, Durabilite: 120,
	Competences: []Competence{
		{Nom: "Coup du Mammouth", Description: "Brise fortement l'armure.", Degats: 40, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 3},
		{Nom: "Rage du Chasseur", Description: "Boost important des dégâts.", Degats: 10, CoutMana: 5, Type: "physique", TypeEffet: "Augmentation de Dégâts", Puissance: 3},
	},
}

// --- Lignée Magique avec progression ---
var etincelle = Arme{
	Nom: "Étincelle", DegatsPhysiques: 0, DegatsMagiques: 10, Precision: 0.85, TauxCritique: 0.15, Durabilite: 30,
	Competences: []Competence{
		{Nom: "Jet de Foudre", Description: "Faibles dégâts magiques.", Degats: 10, CoutMana: 3, Type: "magique"},
	},
}

var foudreMineure = Arme{
	Nom: "Foudre Mineure", DegatsPhysiques: 0, DegatsMagiques: 25, Precision: 0.88, TauxCritique: 0.20, Durabilite: 40,
	Competences: []Competence{
		{Nom: "Décharge Magique", Description: "Dégâts moyens.", Degats: 25, CoutMana: 5, Type: "magique"},
		{Nom: "Brouillard Funeste", Description: "Réduit légèrement la précision.", Degats: 0, CoutMana: 5, Type: "magique", TypeEffet: "Nébulation", Puissance: 1},
	},
}

var foudreSombre = Arme{
	Nom: "Foudre Sombre", DegatsPhysiques: 0, DegatsMagiques: 50, Precision: 0.90, TauxCritique: 0.25, Durabilite: 60,
	Competences: []Competence{
		{Nom: "Frappe Obscure", Description: "Brise modérément la défense magique.", Degats: 50, CoutMana: 8, Type: "magique", TypeEffet: "Brise-Armure Magique", Puissance: 2},
		{Nom: "Brouillard Funeste", Description: "Réduit bien la précision.", Degats: 0, CoutMana: 10, Type: "magique", TypeEffet: "Nébulation", Puissance: 3},
	},
}

var foudreDivine = Arme{
	Nom: "Foudre Divine", DegatsPhysiques: 0, DegatsMagiques: 100, Precision: 0.95, TauxCritique: 0.40, Durabilite: 100,
	Competences: []Competence{
		{Nom: "Choc Divin", Description: "Détruit la défense magique.", Degats: 100, CoutMana: 15, Type: "magique", TypeEffet: "Brise-Armure Magique", Puissance: 4},
		{Nom: "Brouillard Funeste", Description: "Diminue fortement la précision.", Degats: 0, CoutMana: 12, Type: "magique", TypeEffet: "Nébulation", Puissance: 4},
	},
}

var foutreDeZeus = Arme{
	Nom: "Foutre de Zeus", DegatsPhysiques: 0, DegatsMagiques: 200, Precision: 0.99, TauxCritique: 0.60, Durabilite: 200,
	Competences: []Competence{
		{Nom: "Décharge Ultime", Description: "Annihile la défense magique.", Degats: 200, CoutMana: 20, Type: "magique", TypeEffet: "Brise-Armure Magique", Puissance: 5},
		{Nom: "Brouillard Funeste", Description: "Aveugle presque totalement.", Degats: 0, CoutMana: 15, Type: "magique", TypeEffet: "Nébulation", Puissance: 5},
	},
}

// --- Lignée Sabres & Katanas avec progression ---
var sabreBasique = Arme{
	Nom: "Sabre Basique", DegatsPhysiques: 16, DegatsMagiques: 0, Precision: 0.90, TauxCritique: 0.25, Durabilite: 60,
	Competences: []Competence{
		{Nom: "Frappe éclair", Description: "Attaque rapide.", Degats: 12, CoutMana: 0, Type: "physique"},
	},
}

var katana = Arme{
	Nom: "Katana", DegatsPhysiques: 22, DegatsMagiques: 0, Precision: 0.92, TauxCritique: 0.35, Durabilite: 80,
	Competences: []Competence{
		{Nom: "Taillade", Description: "Une coupe rapide et précise.", Degats: 18, CoutMana: 0, Type: "physique"},
		{Nom: "Coupe Sanglante", Description: "Saignement léger.", Degats: 15, CoutMana: 2, Type: "physique", TypeEffet: "Saignement", Puissance: 1},
	},
}

var katanaShuriken = Arme{
	Nom: "Katana Shuriken", DegatsPhysiques: 25, DegatsMagiques: 5, Precision: 0.91, TauxCritique: 0.40, Durabilite: 90,
	Competences: []Competence{
		{Nom: "Frappe Tranchante", Description: "Frappe physique du Katana.", Degats: 20, CoutMana: 0, Type: "physique"},
		{Nom: "Shuriken Volant", Description: "Chance modérée d'étourdissement.", Degats: 10, CoutMana: 3, Type: "physique", TypeEffet: "Étourdissement", Puissance: 2},
		{Nom: "Lame Sanglante", Description: "Saignement modéré.", Degats: 18, CoutMana: 4, Type: "physique", TypeEffet: "Saignement", Puissance: 2},
	},
}

var katanaLameCeleste = Arme{
	Nom: "Katana Lame Céleste", DegatsPhysiques: 35, DegatsMagiques: 10, Precision: 0.93, TauxCritique: 0.50, Durabilite: 120,
	Competences: []Competence{
		{Nom: "Frappe Céleste", Description: "Augmente les dégâts.", Degats: 30, CoutMana: 0, Type: "physique", TypeEffet: "Augmentation de Dégâts", Puissance: 3},
		{Nom: "Shuriken Fantôme", Description: "Forte chance d'étourdissement.", Degats: 15, CoutMana: 5, Type: "physique", TypeEffet: "Étourdissement", Puissance: 4},
		{Nom: "Hémorragie Céleste", Description: "Saignement sévère.", Degats: 25, CoutMana: 6, Type: "physique", TypeEffet: "Saignement", Puissance: 4},
	},
}

// --- Lignée Bâtons avec progression ---
var batonDeMage = Arme{
	Nom: "Bâton de Mage", DegatsPhysiques: 2, DegatsMagiques: 22, Precision: 0.88, TauxCritique: 0.15, Durabilite: 50,
	Competences: []Competence{
		{Nom: "Rayon Arcanique", Description: "Dégâts magiques directs.", Degats: 20, CoutMana: 8, Type: "magique"},
	},
}

var batonArcanique = Arme{
	Nom: "Bâton Arcanique", DegatsPhysiques: 3, DegatsMagiques: 30, Precision: 0.89, TauxCritique: 0.18, Durabilite: 70,
	Competences: []Competence{
		{Nom: "Rayon de Feu", Description: "Brûlure légère.", Degats: 25, CoutMana: 10, Type: "magique", TypeEffet: "Brûlure", Puissance: 1},
	},
}

var batonElementaire = Arme{
	Nom: "Bâton des Élémentaires", DegatsPhysiques: 5, DegatsMagiques: 40, Precision: 0.91, TauxCritique: 0.25, Durabilite: 90,
	Competences: []Competence{
		{Nom: "Explosion Élémentaire", Description: "Brûlure modérée.", Degats: 35, CoutMana: 12, Type: "magique", TypeEffet: "Brûlure", Puissance: 3},
		{Nom: "Vague de Glace", Description: "Ralentit l'ennemi.", Degats: 28, CoutMana: 10, Type: "magique", TypeEffet: "Nébulation", Puissance: 2},
	},
}

var batonGrandMage = Arme{
	Nom: "Bâton du Grand Mage", DegatsPhysiques: 5, DegatsMagiques: 20, Precision: 0.92, TauxCritique: 0.30, Durabilite: 120,
	Competences: []Competence{
		{Nom: "Sceau Arcanique", Description: "Brise fortement l'armure magique.", Degats: 20, CoutMana: 10, Type: "magique", TypeEffet: "Brise-Armure Magique", Puissance: 3},
		{Nom: "Explosion Magique", Description: "Réduit fortement le critique.", Degats: 50, CoutMana: 15, Type: "magique", TypeEffet: "Défavorisation", Puissance: 4},
	},
}

// --- Lignée Bananes avec progression exponentielle ---
var banane = Arme{
	Nom: "Banane", DegatsPhysiques: 4, DegatsMagiques: 3, Precision: 0.99, TauxCritique: 0.05, Durabilite: 20,
	Competences: []Competence{
		{Nom: "Glissade Fatale", Description: "Très faible chance d'étourdissement.", Degats: 3, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 0}, // Puissance 0 = 30% de chance
		{Nom: "Jus de Banane Magique", Description: "Soigne très légèrement.", Degats: 0, CoutMana: 2, Type: "magique", TypeEffet: "Régénération", Puissance: 0},
	},
}

var bananierCombat = Arme{
	Nom: "Bananier de Combat", DegatsPhysiques: 10, DegatsMagiques: 8, Precision: 0.98, TauxCritique: 0.12, Durabilite: 45,
	Competences: []Competence{
		{Nom: "Jet de Banane", Description: "Faible chance d'étourdir.", Degats: 8, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 1},
		{Nom: "Peau Glissante", Description: "Brise légèrement l'armure.", Degats: 6, CoutMana: 3, Type: "magique", TypeEffet: "Brise-Armure", Puissance: 1},
		{Nom: "Smoothie Régénérant", Description: "Soigne légèrement.", Degats: 0, CoutMana: 3, Type: "magique", TypeEffet: "Régénération", Puissance: 1},
	},
}

var lanceBanane = Arme{
	Nom: "Lance-Banane", DegatsPhysiques: 28, DegatsMagiques: 24, Precision: 0.96, TauxCritique: 0.20, Durabilite: 75,
	Competences: []Competence{
		{Nom: "Frappe Glissante", Description: "Chance modérée d'étourdissement.", Degats: 22, CoutMana: 0, Type: "physique", TypeEffet: "Étourdissement", Puissance: 2},
		{Nom: "Pulvérisation de Peau", Description: "Soins modérés.", Degats: 20, CoutMana: 3, Type: "magique", TypeEffet: "Régénération", Puissance: 2},
		{Nom: "Destruction d'Armure Fruitée", Description: "Brise bien l'armure.", Degats: 15, CoutMana: 4, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 2},
	},
}

var bananeRoyale = Arme{
	Nom: "Banane Royale", DegatsPhysiques: 45, DegatsMagiques: 60, Precision: 0.95, TauxCritique: 0.30, Durabilite: 120,
	Competences: []Competence{
		{Nom: "Furie du Singe", Description: "Forte probabilité d'étourdissement.", Degats: 35, CoutMana: 5, Type: "physique", TypeEffet: "Étourdissement", Puissance: 4},
		{Nom: "Peau de Banane Explosive", Description: "Soins importants.", Degats: 55, CoutMana: 5, Type: "magique", TypeEffet: "Régénération", Puissance: 3},
		{Nom: "Écrasement Fruité", Description: "Détruit l'armure.", Degats: 30, CoutMana: 6, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 4},
	},
}

var bananeDivine = Arme{
	Nom: "Banane Divine", DegatsPhysiques: 80, DegatsMagiques: 120, Precision: 0.98, TauxCritique: 0.50, Durabilite: 200,
	Competences: []Competence{
		{Nom: "Frappe Primordiale", Description: "Étourdissement quasi-garanti.", Degats: 70, CoutMana: 5, Type: "physique", TypeEffet: "Étourdissement", Puissance: 5},
		{Nom: "Explosion Bananière", Description: "Soins massifs.", Degats: 110, CoutMana: 5, Type: "magique", TypeEffet: "Régénération", Puissance: 5},
		{Nom: "Annihilation d'Armure", Description: "Annihile l'armure.", Degats: 50, CoutMana: 8, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 5},
	},
}

// --- Lignée AR/SCAR avec progression ---
var arGrise = Arme{
	Nom: "AR Grise", DegatsPhysiques: 20, DegatsMagiques: 0, Precision: 0.88, TauxCritique: 0.10, Durabilite: 50,
	Competences: []Competence{
		{Nom: "Tir Simple", Description: "Tir direct simple.", Degats: 15, CoutMana: 0, Type: "physique"},
		{Nom: "Tir Appuyé", Description: "Brise légèrement l'armure.", Degats: 18, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 1},
	},
}

var arVerte = Arme{
	Nom: "AR Verte", DegatsPhysiques: 30, DegatsMagiques: 0, Precision: 0.90, TauxCritique: 0.15, Durabilite: 70,
	Competences: []Competence{
		{Nom: "Rafale Rapide", Description: "Tire plusieurs balles.", Degats: 25, CoutMana: 0, Type: "physique"},
		{Nom: "Tir Perçant", Description: "Brise modérément l'armure.", Degats: 28, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 2},
	},
}

var arBleue = Arme{
	Nom: "AR Bleue", DegatsPhysiques: 45, DegatsMagiques: 0, Precision: 0.94, TauxCritique: 0.30, Durabilite: 95,
	Competences: []Competence{
		{Nom: "Tir Perçant", Description: "Brise bien l'armure.", Degats: 40, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 3},
		{Nom: "Bush", Description: "Diminue bien la précision.", Degats: 0, CoutMana: 3, Type: "magique", TypeEffet: "Nébulation", Puissance: 3},
	},
}

var scarViolette = Arme{
	Nom: "SCAR Violette", DegatsPhysiques: 60, DegatsMagiques: 10, Precision: 0.95, TauxCritique: 0.40, Durabilite: 125,
	Competences: []Competence{
		{Nom: "Tir Perçant", Description: "Détruit l'armure.", Degats: 50, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 4},
		{Nom: "Bush", Description: "Diminue fortement la précision.", Degats: 0, CoutMana: 4, Type: "magique", TypeEffet: "Nébulation", Puissance: 4},
	},
}

var scarEnOr = Arme{
	Nom: "SCAR en Or", DegatsPhysiques: 80, DegatsMagiques: 20, Precision: 0.97, TauxCritique: 0.55, Durabilite: 160,
	Competences: []Competence{
		{Nom: "Tir Perçant", Description: "Annihile l'armure.", Degats: 65, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 5},
		{Nom: "Bush", Description: "Aveugle presque totalement.", Degats: 0, CoutMana: 5, Type: "magique", TypeEffet: "Nébulation", Puissance: 5},
	},
}

// --- Lignée Trident/Fourchette avec progression ---
var fourchetteDesMers = Arme{
	Nom: "Fourchette des Mers", DegatsPhysiques: 8, DegatsMagiques: 5, Precision: 0.85, TauxCritique: 0.20, Durabilite: 50,
	Competences: []Competence{
		{Nom: "Pique Marine", Description: "Brise très légèrement l'armure.", Degats: 8, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 0},
		{Nom: "Vague Distrayante", Description: "Réduit légèrement le critique.", Degats: 5, CoutMana: 4, Type: "magique", TypeEffet: "Défavorisation", Puissance: 1},
	},
}

var tridentDuMarais = Arme{
	Nom: "Trident du Marais", DegatsPhysiques: 14, DegatsMagiques: 12, Precision: 0.88, TauxCritique: 0.25, Durabilite: 75,
	Competences: []Competence{
		{Nom: "Lancer Tranchant", Description: "Brise modérément l'armure.", Degats: 14, CoutMana: 2, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 2},
		{Nom: "Vague Aquatique", Description: "Réduit bien le critique.", Degats: 12, CoutMana: 5, Type: "magique", TypeEffet: "Défavorisation", Puissance: 2},
	},
}

var tridentDesProfondeurs = Arme{
	Nom: "Trident des Profondeurs", DegatsPhysiques: 32, DegatsMagiques: 30, Precision: 0.90, TauxCritique: 0.35, Durabilite: 110,
	Competences: []Competence{
		{Nom: "Lancer Perçant", Description: "Détruit l'armure.", Degats: 32, CoutMana: 4, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 4},
		{Nom: "Vague Déstabilisante", Description: "Réduit fortement le critique.", Degats: 30, CoutMana: 8, Type: "magique", TypeEffet: "Défavorisation", Puissance: 3},
	},
}

var tridentPoseidon = Arme{
	Nom: "Trident de Poséidon", DegatsPhysiques: 48, DegatsMagiques: 55, Precision: 0.95, TauxCritique: 0.50, Durabilite: 150,
	Competences: []Competence{
		{Nom: "Lancer Divin", Description: "Annihile l'armure.", Degats: 48, CoutMana: 6, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 5},
		{Nom: "Déluge de Poséidon", Description: "Supprime le critique adverse.", Degats: 55, CoutMana: 12, Type: "magique", TypeEffet: "Défavorisation", Puissance: 5},
	},
}

// --- Armes spéciales ---
var dragonLore = Arme{
	Nom: "AWP Dragon Lore", DegatsPhysiques: 120, DegatsMagiques: 30, Precision: 0.80, TauxCritique: 0.65, Durabilite: 80,
	Competences: []Competence{
		{Nom: "Tir Légendaire", Description: "Un tir dévastateur.", Degats: 150, CoutMana: 0, Type: "physique"},
		{Nom: "Souffle du Dragon", Description: "Brûlure dévastatrice.", Degats: 50, CoutMana: 10, Type: "magique", TypeEffet: "Brûlure", Puissance: 5},
	},
}

// --- Potions de soin --- //
var potionMineure = Arme{
	Nom: "Potion Mineure", DegatsPhysiques: 0, DegatsMagiques: 0,
	Precision: 1.0, TauxCritique: 0.0, Durabilite: 1,
	Competences: []Competence{
		{Nom: "Soin Léger", Description: "Restaure 30 PV.", Degats: -30, CoutMana: 0, Type: "magique", TypeEffet: "Régénération", Puissance: 1},
	},
}

var potionMajeure = Arme{
	Nom: "Potion Majeure", DegatsPhysiques: 0, DegatsMagiques: 0,
	Precision: 1.0, TauxCritique: 0.0, Durabilite: 1,
	Competences: []Competence{
		{Nom: "Soin Puissant", Description: "Restaure 80 PV.", Degats: -80, CoutMana: 0, Type: "magique", TypeEffet: "Régénération", Puissance: 3},
	},
}

var potionSupreme = Arme{
	Nom: "Potion Suprême", DegatsPhysiques: 0, DegatsMagiques: 0,
	Precision: 1.0, TauxCritique: 0.0, Durabilite: 1,
	Competences: []Competence{
		{Nom: "Soin Divin", Description: "Restaure 200 PV.", Degats: -200, CoutMana: 0, Type: "magique", TypeEffet: "Régénération", Puissance: 6},
	},
}

// --- Potions offensives --- //
var potionDegats = Arme{
	Nom: "Potion de Dégâts", DegatsPhysiques: 0, DegatsMagiques: 30,
	Precision: 1.0, TauxCritique: 0.0, Durabilite: 1,
	Competences: []Competence{
		{Nom: "Lancer Explosif", Description: "Explose au contact.", Degats: 30, CoutMana: 0, Type: "magique"},
	},
}

var bombeIncendiaire = Arme{
	Nom: "Bombe Incendiaire", DegatsPhysiques: 0, DegatsMagiques: 50,
	Precision: 0.95, TauxCritique: 0.25, Durabilite: 1,
	Competences: []Competence{
		{Nom: "Explosion de Feu", Description: "Brûlure intense.", Degats: 50, CoutMana: 0, Type: "magique", TypeEffet: "Brûlure", Puissance: 4},
	},
}

var bombeGivrante = Arme{
	Nom: "Bombe Givrante", DegatsPhysiques: 0, DegatsMagiques: 40,
	Precision: 0.9, TauxCritique: 0.15, Durabilite: 1,
	Competences: []Competence{
		// Remplacé "Gel" par "Étourdissement"
		{Nom: "Explosion de Glace", Description: "Inflige des dégâts et ralentit la cible.", Degats: 40, CoutMana: 0, Type: "magique", TypeEffet: "Étourdissement", Puissance: 2},
	},
}

// --- Utilitaires --- //
var grenadeFumigene = Arme{
	Nom: "Grenade Fumigène", DegatsPhysiques: 0, DegatsMagiques: 0,
	Precision: 1.0, TauxCritique: 0.0, Durabilite: 1,
	Competences: []Competence{
		{Nom: "Écran de Fumée", Description: "Aveugle modérément.", Degats: 0, CoutMana: 0, Type: "magique", TypeEffet: "Nébulation", Puissance: 3},
	},
}

var parcheminDispersion = Arme{
	Nom: "Parchemin de Dispersion", DegatsPhysiques: 0, DegatsMagiques: 0,
	Precision: 1.0, TauxCritique: 0.0, Durabilite: 1,
	Competences: []Competence{
		// "Anti-Buff" remplacé par "Affaiblissement"
		{Nom: "Dissipation", Description: "Affaiblit la cible en dissipant ses forces.", Degats: 0, CoutMana: 0, Type: "magique", TypeEffet: "Affaiblissement", Puissance: 2},
	},
}

// --- Buffs --- //
var elixirDeForce = Arme{
	Nom: "Élixir de Force", DegatsPhysiques: 0, DegatsMagiques: 0,
	Precision: 1.0, TauxCritique: 0.0, Durabilite: 1,
	Competences: []Competence{
		{Nom: "Puissance Brute", Description: "Augmente fortement les dégâts.", Degats: 0, CoutMana: 0, Type: "magique", TypeEffet: "Augmentation de Dégâts", Puissance: 4},
	},
}

var elixirDeVitesse = Arme{
	Nom: "Élixir de Vitesse", DegatsPhysiques: 0, DegatsMagiques: 0,
	Precision: 1.0, TauxCritique: 0.0, Durabilite: 1,
	Competences: []Competence{
		// Pas "Augmentation de Vitesse" → remplacé par "Augmentation de Dégâts Magiques" (boost indirect)
		{Nom: "Agilité Accrue", Description: "Augmente la vitesse d’attaque (magique).", Degats: 0, CoutMana: 0, Type: "magique", TypeEffet: "Augmentation de Dégâts Magiques", Puissance: 3},
	},
}

var elixirDeCritique = Arme{
	Nom: "Élixir de Précision", DegatsPhysiques: 0, DegatsMagiques: 0,
	Precision: 1.0, TauxCritique: 0.0, Durabilite: 1,
	Competences: []Competence{
		// Pas "Augmentation de Critique" → remplacé par "Augmentation de Dégâts"
		{Nom: "Visée Parfaite", Description: "Augmente les chances de critique en boostant les dégâts.", Degats: 0, CoutMana: 0, Type: "magique", TypeEffet: "Augmentation de Dégâts", Puissance: 5},
	},
}

var antidote = Arme{
	Nom:             "Antidote",
	DegatsPhysiques: 0,
	DegatsMagiques:  0,
	Precision:       1.0,
	TauxCritique:    0.0,
	Durabilite:      1,
	Competences: []Competence{
		{
			Nom:         "Purification",
			Description: "Supprime immédiatement l’empoisonnement.",
			Degats:      0,
			CoutMana:    0,
			Type:        "magique",
			TypeEffet:   "Guérison Poison",
			Puissance:   1,
		},
	},
}

// --- Armes d'Erwann (thème informatique) ---
var macMini = Arme{
	Nom: "Mac Mini", DegatsPhysiques: 15, DegatsMagiques: 10, Precision: 0.82, TauxCritique: 0.10, Durabilite: 80,
	Competences: []Competence{
		{Nom: "Compilation Rapide", Description: "Petit burst de calcul.", Degats: 18, CoutMana: 3, Type: "magique"},
	},
}

var macAir = Arme{
	Nom: "Mac Air", DegatsPhysiques: 20, DegatsMagiques: 18, Precision: 0.85, TauxCritique: 0.15, Durabilite: 90,
	Competences: []Competence{
		{Nom: "Ventilation Silencieuse", Description: "Réduit légèrement la précision ennemie.", Degats: 0, CoutMana: 5, Type: "magique", TypeEffet: "Nébulation", Puissance: 1},
		{Nom: "Process M1", Description: "Cycle optimisé.", Degats: 25, CoutMana: 4, Type: "magique"},
	},
}

var macPro = Arme{
	Nom: "Mac Pro", DegatsPhysiques: 28, DegatsMagiques: 35, Precision: 0.88, TauxCritique: 0.22, Durabilite: 110,
	Competences: []Competence{
		{Nom: "Rendu 3D", Description: "Charge CPU massive.", Degats: 40, CoutMana: 6, Type: "magique"},
		{Nom: "Kernel Panic", Description: "Perturbe l'ennemi (précision -).", Degats: 0, CoutMana: 7, Type: "magique", TypeEffet: "Nébulation", Puissance: 2},
	},
}

var pcDuCDI = Arme{
	Nom: "PC du CDI", DegatsPhysiques: 40, DegatsMagiques: 55, Precision: 0.90, TauxCritique: 0.30, Durabilite: 140,
	Competences: []Competence{
		{Nom: "Ecran Bleu", Description: "Gèle net l'initiative adverse (étourdissement).", Degats: 0, CoutMana: 10, Type: "magique", TypeEffet: "Étourdissement", Puissance: 3},
		{Nom: "Torrent Interdit", Description: "Déluge de paquets.", Degats: 60, CoutMana: 8, Type: "magique"},
	},
}

// --- Armes de Gabriel (thème céleste / jugement) ---
var vergeCeleste = Arme{
	Nom: "Verge Céleste", DegatsPhysiques: 35, DegatsMagiques: 20, Precision: 0.88, TauxCritique: 0.20, Durabilite: 120,
	Competences: []Competence{
		{Nom: "Frappe Lumineuse", Description: "Coup de base sanctifié.", Degats: 32, CoutMana: 0, Type: "physique"},
		{Nom: "Marque Sacrée", Description: "Affaiblit la cible.", Degats: 18, CoutMana: 4, Type: "magique", TypeEffet: "Affaiblissement", Puissance: 2},
	},
}

var lanceArchange = Arme{
	Nom: "Lance de l'Archange", DegatsPhysiques: 50, DegatsMagiques: 40, Precision: 0.90, TauxCritique: 0.25, Durabilite: 140,
	Competences: []Competence{
		{Nom: "Perçée Divine", Description: "Brise fortement l'armure.", Degats: 55, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 3},
		{Nom: "Châtiment", Description: "Inflige des dégâts sacrés.", Degats: 48, CoutMana: 6, Type: "magique"},
	},
}

var trompeteJugement = Arme{
	Nom: "Trompette du Jugement", DegatsPhysiques: 20, DegatsMagiques: 95, Precision: 0.92, TauxCritique: 0.35, Durabilite: 160,
	Competences: []Competence{
		{Nom: "Onde Sacrée", Description: "Réduit la résistance magique.", Degats: 60, CoutMana: 8, Type: "magique", TypeEffet: "Brise-Armure Magique", Puissance: 3},
		{Nom: "Résonance Céleste", Description: "Diminue la précision ennemie.", Degats: 0, CoutMana: 8, Type: "magique", TypeEffet: "Nébulation", Puissance: 3},
	},
}

var glaiveApocalypse = Arme{
	Nom: "Glaive de l'Apocalypse", DegatsPhysiques: 85, DegatsMagiques: 120, Precision: 0.95, TauxCritique: 0.50, Durabilite: 220,
	Competences: []Competence{
		{Nom: "Sentence Finale", Description: "Annihile l'armure.", Degats: 90, CoutMana: 10, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 5},
		{Nom: "Chœur des Cieux", Description: "Désintègre la résistance magique.", Degats: 110, CoutMana: 12, Type: "magique", TypeEffet: "Brise-Armure Magique", Puissance: 5},
		{Nom: "Jugement Dernier", Description: "Saignement apocalyptique.", Degats: 70, CoutMana: 10, Type: "physique", TypeEffet: "Saignement", Puissance: 5},
	},
}

// --- Armes de Vitaly (thème vodka / force sauvage) ---
var flashDeVodka = Arme{
	Nom: "Flash de Vodka", DegatsPhysiques: 15, DegatsMagiques: 5, Precision: 0.95, TauxCritique: 0.30, Durabilite: 80,
	Competences: []Competence{
		{Nom: "Gorgée Brûlante", Description: "Petit choc alcoolisé (brûlure légère).", Degats: 12, CoutMana: 2, Type: "magique", TypeEffet: "Brûlure", Puissance: 1},
	},
}

var bouteilleDeVodka = Arme{
	Nom: "Bouteille de Vodka", DegatsPhysiques: 30, DegatsMagiques: 10, Precision: 0.97, TauxCritique: 0.45, Durabilite: 110,
	Competences: []Competence{
		{Nom: "Coup de Bouteille", Description: "Brise légèrement l'armure.", Degats: 28, CoutMana: 0, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 1},
		{Nom: "Jet Enflammé", Description: "Brûlure modérée.", Degats: 18, CoutMana: 4, Type: "magique", TypeEffet: "Brûlure", Puissance: 2},
	},
}

var griffeDOurs = Arme{
	Nom: "Griffe d'Ours", DegatsPhysiques: 55, DegatsMagiques: 15, Precision: 0.98, TauxCritique: 0.55, Durabilite: 150,
	Competences: []Competence{
		{Nom: "Déchirure", Description: "Saignement sévère.", Degats: 40, CoutMana: 5, Type: "physique", TypeEffet: "Saignement", Puissance: 4},
		{Nom: "Rugissement Glacial", Description: "Réduit la précision ennemie.", Degats: 0, CoutMana: 6, Type: "magique", TypeEffet: "Nébulation", Puissance: 3},
	},
}

var apocalypseVodka = Arme{
	Nom: "Apocalypse Vodka", DegatsPhysiques: 95, DegatsMagiques: 80, Precision: 0.99, TauxCritique: 0.70, Durabilite: 240,
	Competences: []Competence{
		{Nom: "Explosion Alcoolisée", Description: "Dégâts massifs et brûlure.", Degats: 120, CoutMana: 12, Type: "magique", TypeEffet: "Brûlure", Puissance: 5},
		{Nom: "Morsure de l'Hiver", Description: "Affaiblit l'ennemi et réduit sa précision.", Degats: 30, CoutMana: 8, Type: "magique", TypeEffet: "Défavorisation", Puissance: 4},
		{Nom: "Éclats de Verre", Description: "Brise fortement l'armure.", Degats: 60, CoutMana: 6, Type: "physique", TypeEffet: "Brise-Armure", Puissance: 4},
	},
}

// Potion spéciale — régénère toute la vie mais pénalise la précision
var vodkaDeVitaly = Arme{
	Nom:             "Vodka de Vitaly",
	DegatsPhysiques: 0,
	DegatsMagiques:  0,
	Precision:       1.0,
	TauxCritique:    0.0,
	Durabilite:      1,
	Competences: []Competence{
		{
			Nom:         "Coup de Fouet",
			Description: "Régénère toute la vie mais trouble la vision (-30% précision pendant 3 tours).",
			Degats:      0,
			CoutMana:    0,
			Type:        "magique",
			TypeEffet:   "Ivresse",
			Puissance:   0,
		},
	},
}
