package main

import "fmt"

type Artefact struct {
	Nom         string
	Description string
	Effet       Effet
}

// Équipe un artefact en l'appliquant comme effet permanent
func EquiperArtefact(p *Personnage, a Artefact) {
	// Retire l'effet précédent du même artefact s'il existe (évite les doublons)
	DesequiperArtefact(p, a.Nom)

	eff := a.Effet
	// Nom interne distinct pour l'affichage/gestion
	eff.Nom = "Artefact: " + a.Nom
	// Rendre permanent et déterministe
	if eff.ToursRestants < 999 {
		eff.ToursRestants = 999
	}
	eff.ChanceAppliquer = 1.0

	// Ajout direct comme effet actif
	p.EffetsActifs = append(p.EffetsActifs, EffetActif{Effet: eff, ToursRestants: eff.ToursRestants})
}

// Équipe un artefact dans un des 2 slots (0 ou 1) et applique son effet
func EquiperArtefactDansSlot(p *Personnage, a Artefact, slot int) error {
	if slot < 0 || slot >= len(p.ArtefactsEquipes) {
		return fmt.Errorf("slot d'artefact invalide: %d", slot)
	}
	// Retire l'ancien artefact du slot (effet associé)
	if p.ArtefactsEquipes[slot] != nil {
		DesequiperArtefact(p, p.ArtefactsEquipes[slot].Nom)
	}
	// Équipe le nouveau
	p.ArtefactsEquipes[slot] = &a
	EquiperArtefact(p, a)
	return nil
}

// Déséquipe l'artefact d'un slot
func DesequiperArtefactDuSlot(p *Personnage, slot int) error {
	if slot < 0 || slot >= len(p.ArtefactsEquipes) {
		return fmt.Errorf("slot d'artefact invalide: %d", slot)
	}
	if p.ArtefactsEquipes[slot] != nil {
		DesequiperArtefact(p, p.ArtefactsEquipes[slot].Nom)
		p.ArtefactsEquipes[slot] = nil
	}
	return nil
}

// Retourne la liste des artefacts équipés (non-nil)
func ListeArtefactsEquipes(p *Personnage) []Artefact {
	var res []Artefact
	for _, ptr := range p.ArtefactsEquipes {
		if ptr != nil {
			res = append(res, *ptr)
		}
	}
	return res
}

// Retire l'effet d'un artefact par son nom
func DesequiperArtefact(p *Personnage, nom string) {
	SupprimerEffet(p, "Artefact: "+nom)
}

// Récupère un artefact par son nom (retourne bool si trouvé)
func GetArtefactParNom(nom string) (Artefact, bool) {
	for _, a := range ArtefactsDisponibles {
		if a.Nom == nom {
			return a, true
		}
	}
	return Artefact{}, false
}

// Une quinzaine d'artefacts, proches des thèmes des classes et de l'univers
var ArtefactsDisponibles = []Artefact{
	{
		Nom:         "Insigne du Sergent",
		Description: "Emblème du soldat discipliné: +15% dégâts physiques, +5% précision.",
		Effet:       Effet{ToursRestants: 999, ModifDegats: 0.15, ModifPrecision: 0.05},
	},
	{
		Nom:         "Gant Anti-Émeute",
		Description: "Standard CRS: +10% armure et +03% précision.",
		Effet:       Effet{ToursRestants: 999, ModifArmure: 0.10, ModifPrecision: 0.03},
	},
	{
		Nom:         "Talisman du Brasier",
		Description: "Pyromane: +10% dégâts magiques, +5% dégâts physiques.",
		Effet:       Effet{ToursRestants: 999, ModifDegatsMag: 0.10, ModifDegats: 0.05},
	},
	{
		Nom:         "Carquois des Mille Flèches",
		Description: "Robin des Bois: +12% précision, +5% taux critique.",
		Effet:       Effet{ToursRestants: 999, ModifPrecision: 0.12, ModifCritique: 0.05},
	},
	{
		Nom:         "Coutelas Runique",
		Description: "Boucher: +15% dégâts physiques, +5% critique.",
		Effet:       Effet{ToursRestants: 999, ModifDegats: 0.15, ModifCritique: 0.05},
	},
	{
		Nom:         "Dent de Mammouth",
		Description: "Cro-Magnon: +20% dégâts physiques.",
		Effet:       Effet{ToursRestants: 999, ModifDegats: 0.20},
	},
	{
		Nom:         "Anneau des Tempêtes",
		Description: "Zeus: +10% dégâts magiques et +10% taux critique.",
		Effet:       Effet{ToursRestants: 999, ModifDegatsMag: 0.10, ModifCritique: 0.10},
	},
	{
		Nom:         "Bandeau du Ronin",
		Description: "Samouraï: +8% précision, +8% critique.",
		Effet:       Effet{ToursRestants: 999, ModifPrecision: 0.08, ModifCritique: 0.08},
	},
	{
		Nom:         "Perle d'Æther",
		Description: "Gandalf: +12% dégâts magiques, +5% précision.",
		Effet:       Effet{ToursRestants: 999, ModifDegatsMag: 0.12, ModifPrecision: 0.05},
	},
	{
		Nom:         "Peau de Banane Sacrée",
		Description: "Singe: +15% précision.",
		Effet:       Effet{ToursRestants: 999, ModifPrecision: 0.15},
	},
	{
		Nom:         "Puce de Visée",
		Description: "AR/SCAR (Bambi): +8% précision, +7% critique.",
		Effet:       Effet{ToursRestants: 999, ModifPrecision: 0.08, ModifCritique: 0.07},
	},
	{
		Nom:         "Conque des Profondeurs",
		Description: "Poséidon: +10% résistance magique, +5% précision.",
		Effet:       Effet{ToursRestants: 999, ModifResistMag: 0.10, ModifPrecision: 0.05},
	},
	{
		Nom:         "Médaillon de Foudre Pure",
		Description: "Zeus: +18% dégâts magiques.",
		Effet:       Effet{ToursRestants: 999, ModifDegatsMag: 0.18},
	},
	{
		Nom:         "Relique de la Sylve",
		Description: "Robin des Bois: +5% armure, +5% précision.",
		Effet:       Effet{ToursRestants: 999, ModifArmure: 0.05, ModifPrecision: 0.05},
	},
	{
		Nom:         "Pierre d'Ignition",
		Description: "Pyromane: +8% dégâts magiques, +4% critique.",
		Effet:       Effet{ToursRestants: 999, ModifDegatsMag: 0.08, ModifCritique: 0.04},
	},
	// --- Nouveaux artefacts (utilitaires et variés) ---
	{
		Nom:         "Amulette Anti-Poison",
		Description: "Protège des toxines: supprime le Poison au début de chaque tour.",
		Effet:       Effet{ToursRestants: 999}, // pas de modif de stats
	},
	{
		Nom:         "Pendentif Purificateur",
		Description: "Purifie l'âme: retire tous les débuffs au début de chaque tour.",
		Effet:       Effet{ToursRestants: 999},
	},
	{
		Nom:         "Talisman Stoïque",
		Description: "Donne du courage: dissipe Peur et Étourdissement au début de chaque tour.",
		Effet:       Effet{ToursRestants: 999},
	},
	{
		Nom:         "Totem de Refroidissement",
		Description: "Éteint les flammes: supprime Brûlure et Saignement au début de chaque tour.",
		Effet:       Effet{ToursRestants: 999},
	},
	{
		Nom:         "Œil de Lynx",
		Description: "Visée parfaite: +12% précision, +6% critique.",
		Effet:       Effet{ToursRestants: 999, ModifPrecision: 0.12, ModifCritique: 0.06},
	},
	{
		Nom:         "Rune de Trempe",
		Description: "Acier trempé: +10% armure, +6% résistance magique.",
		Effet:       Effet{ToursRestants: 999, ModifArmure: 0.10, ModifResistMag: 0.06},
	},
	{
		Nom:         "Coquille Abyssale",
		Description: "Bénédiction des abysses: +15% résistance magique.",
		Effet:       Effet{ToursRestants: 999, ModifResistMag: 0.15},
	},
	{
		Nom:         "Médaillon du Chasseur de Mages",
		Description: "Pourfendeur d'arcanes: +10% dégâts physiques, +10% résistance magique.",
		Effet:       Effet{ToursRestants: 999, ModifDegats: 0.10, ModifResistMag: 0.10},
	},
	{
		Nom:         "Glyphe de Parade",
		Description: "Maîtrise défensive: +15% armure.",
		Effet:       Effet{ToursRestants: 999, ModifArmure: 0.15},
	},
	{
		Nom:         "Boussole de Chasseur",
		Description: "Toujours sur la cible: +6% précision, +6% dégâts physiques.",
		Effet:       Effet{ToursRestants: 999, ModifPrecision: 0.06, ModifDegats: 0.06},
	},
	{
		Nom:         "Éclat de Foudre Gelée",
		Description: "Magie stable: +10% dégâts magiques, +3% précision.",
		Effet:       Effet{ToursRestants: 999, ModifDegatsMag: 0.10, ModifPrecision: 0.03},
	},
	// Artefacts unitaires de dissipation
	{Nom: "Antidote Éternel", Description: "Supprime le Poison au début de chaque tour.", Effet: Effet{ToursRestants: 999}},
	{Nom: "Talisman Éteigneflamme", Description: "Supprime la Brûlure au début de chaque tour.", Effet: Effet{ToursRestants: 999}},
	{Nom: "Sceau Hémostatique", Description: "Supprime le Saignement au début de chaque tour.", Effet: Effet{ToursRestants: 999}},
	{Nom: "Pendentif de Courage", Description: "Supprime la Peur au début de chaque tour.", Effet: Effet{ToursRestants: 999}},
	{Nom: "Talisman de Vigilance", Description: "Supprime l'Étourdissement au début de chaque tour.", Effet: Effet{ToursRestants: 999}},
	{Nom: "Sceau de Focalisation", Description: "Supprime Nébulation et Défavorisation au début de chaque tour.", Effet: Effet{ToursRestants: 999}},
	{Nom: "Glyphe de Bastion", Description: "Supprime Brise-Armure et Brise-Armure Magique au début de chaque tour.", Effet: Effet{ToursRestants: 999}},
	{Nom: "Cachet de Détermination", Description: "Supprime les débuffs d'attaque (ex: Peur) au début de chaque tour.", Effet: Effet{ToursRestants: 999}},
}

// --- Helpers supplémentaires ---

// Indique si un artefact précis est actuellement équipé (présent comme effet actif)
func PossedeArtefact(p *Personnage, nom string) bool {
	tag := "Artefact: " + nom
	for _, e := range p.EffetsActifs {
		if e.Effet.Nom == tag {
			return true
		}
	}
	return false
}

// À appeler en début de tour d'un personnage pour appliquer les protections passives
func AppliquerProtectionsArtefactsDebutTour(p *Personnage) {
	// Artefacts unitaires
	if PossedeArtefact(p, "Antidote Éternel") {
		GuerirPoison(p)
	}
	if PossedeArtefact(p, "Talisman Éteigneflamme") {
		SupprimerEffet(p, "Brûlure")
	}
	if PossedeArtefact(p, "Sceau Hémostatique") {
		SupprimerEffet(p, "Saignement")
	}
	if PossedeArtefact(p, "Pendentif de Courage") {
		SupprimerEffet(p, "Peur")
	}
	if PossedeArtefact(p, "Talisman de Vigilance") {
		SupprimerEffet(p, "Étourdissement")
	}
	if PossedeArtefact(p, "Sceau de Focalisation") {
		SupprimerDebuffsCritPrecision(p)
	}
	if PossedeArtefact(p, "Glyphe de Bastion") {
		SupprimerDebuffsArmure(p)
	}
	if PossedeArtefact(p, "Cachet de Détermination") {
		SupprimerDebuffsAttaque(p)
	}
}

// Supprime tous les débuffs de critique et de précision
func SupprimerDebuffsCritPrecision(p *Personnage) {
	SupprimerEffet(p, "Nébulation")
	SupprimerEffet(p, "Défavorisation")
}

// Supprime tous les débuffs d'armure (physique et magique)
func SupprimerDebuffsArmure(p *Personnage) {
	SupprimerEffet(p, "Brise-Armure")
	SupprimerEffet(p, "Brise-Armure Magique")
}

// Supprime les débuffs qui affectent directement l'attaque (dégâts)
func SupprimerDebuffsAttaque(p *Personnage) {
	// Supprime la Peur (réduction des dégâts via modificateur)
	SupprimerEffet(p, "Peur")
}
