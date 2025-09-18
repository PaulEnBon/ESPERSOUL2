package main

import "fmt"

// Gestion des artefacts (objets spéciaux conférant des bonus passifs permanents)
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
	if slot < 0 || slot >= MaxArtefactsEquipes {
		return fmt.Errorf("slot d'artefact invalide: %d", slot)
	}
	// S'assure que le slice a la bonne taille
	if len(p.ArtefactsEquipes) < MaxArtefactsEquipes {
		// étend ou réinitialise proprement
		newSlice := make([]*Artefact, MaxArtefactsEquipes)
		copy(newSlice, p.ArtefactsEquipes)
		p.ArtefactsEquipes = newSlice
	}
	// Retire l'ancien artefact si présent
	if p.ArtefactsEquipes[slot] != nil {
		DesequiperArtefact(p, p.ArtefactsEquipes[slot].Nom)
	}
	p.ArtefactsEquipes[slot] = &a
	EquiperArtefact(p, a)
	// Ajoute à la collection possédée si pas déjà
	possede := false
	for _, owned := range p.ArtefactsPossedes {
		if owned.Nom == a.Nom {
			possede = true
			break
		}
	}
	if !possede {
		p.ArtefactsPossedes = append(p.ArtefactsPossedes, a)
	}
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

// Ajoute un artefact à la collection possédée (sans l'équiper)
func AjouterArtefactPossede(p *Personnage, a Artefact) {
	for _, owned := range p.ArtefactsPossedes {
		if owned.Nom == a.Nom {
			return // déjà possédé
		}
	}
	p.ArtefactsPossedes = append(p.ArtefactsPossedes, a)
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

// Définition individuelle de chaque artefact (variables réutilisables)
// Nom des variables en camelCase sans accents pour simplifier l'utilisation dans le code.
var (
	insigneSergent         = Artefact{Nom: "Insigne du Sergent", Description: "Emblème du soldat discipliné: +15% dégâts physiques, +5% précision.", Effet: Effet{ToursRestants: 999, ModifDegats: 0.15, ModifPrecision: 0.05}}
	gantAntiEmeute         = Artefact{Nom: "Gant Anti-Émeute", Description: "Standard CRS: +10% armure et +03% précision.", Effet: Effet{ToursRestants: 999, ModifArmure: 0.10, ModifPrecision: 0.03}}
	talismanBrasier        = Artefact{Nom: "Talisman du Brasier", Description: "Pyromane: +10% dégâts magiques, +5% dégâts physiques.", Effet: Effet{ToursRestants: 999, ModifDegatsMag: 0.10, ModifDegats: 0.05}}
	carquoisMilleFleches   = Artefact{Nom: "Carquois des Mille Flèches", Description: "Robin des Bois: +12% précision, +5% taux critique.", Effet: Effet{ToursRestants: 999, ModifPrecision: 0.12, ModifCritique: 0.05}}
	coutelasRunique        = Artefact{Nom: "Coutelas Runique", Description: "Boucher: +15% dégâts physiques, +5% critique.", Effet: Effet{ToursRestants: 999, ModifDegats: 0.15, ModifCritique: 0.05}}
	dentMammouth           = Artefact{Nom: "Dent de Mammouth", Description: "Cro-Magnon: +20% dégâts physiques.", Effet: Effet{ToursRestants: 999, ModifDegats: 0.20}}
	anneauTempetes         = Artefact{Nom: "Anneau des Tempêtes", Description: "Zeus: +10% dégâts magiques et +10% taux critique.", Effet: Effet{ToursRestants: 999, ModifDegatsMag: 0.10, ModifCritique: 0.10}}
	bandeauRonin           = Artefact{Nom: "Bandeau du Ronin", Description: "Samouraï: +8% précision, +8% critique.", Effet: Effet{ToursRestants: 999, ModifPrecision: 0.08, ModifCritique: 0.08}}
	perleAether            = Artefact{Nom: "Perle d'Æther", Description: "Gandalf: +12% dégâts magiques, +5% précision.", Effet: Effet{ToursRestants: 999, ModifDegatsMag: 0.12, ModifPrecision: 0.05}}
	peauBananeSacree       = Artefact{Nom: "Peau de Banane Sacrée", Description: "Singe: +15% précision.", Effet: Effet{ToursRestants: 999, ModifPrecision: 0.15}}
	puceVisee              = Artefact{Nom: "Puce de Visée", Description: "AR/SCAR (Bambi): +8% précision, +7% critique.", Effet: Effet{ToursRestants: 999, ModifPrecision: 0.08, ModifCritique: 0.07}}
	conqueProfondeurs      = Artefact{Nom: "Conque des Profondeurs", Description: "Poséidon: +10% résistance magique, +5% précision.", Effet: Effet{ToursRestants: 999, ModifResistMag: 0.10, ModifPrecision: 0.05}}
	medaillonFoudrePure    = Artefact{Nom: "Médaillon de Foudre Pure", Description: "Zeus: +18% dégâts magiques.", Effet: Effet{ToursRestants: 999, ModifDegatsMag: 0.18}}
	reliqueSylve           = Artefact{Nom: "Relique de la Sylve", Description: "Robin des Bois: +5% armure, +5% précision.", Effet: Effet{ToursRestants: 999, ModifArmure: 0.05, ModifPrecision: 0.05}}
	pierreIgnition         = Artefact{Nom: "Pierre d'Ignition", Description: "Pyromane: +8% dégâts magiques, +4% critique.", Effet: Effet{ToursRestants: 999, ModifDegatsMag: 0.08, ModifCritique: 0.04}}
	oeilLynx               = Artefact{Nom: "Œil de Lynx", Description: "Visée parfaite: +12% précision, +6% critique.", Effet: Effet{ToursRestants: 999, ModifPrecision: 0.12, ModifCritique: 0.06}}
	runeTrempe             = Artefact{Nom: "Rune de Trempe", Description: "Acier trempé: +10% armure, +6% résistance magique.", Effet: Effet{ToursRestants: 999, ModifArmure: 0.10, ModifResistMag: 0.06}}
	coquilleAbyssale       = Artefact{Nom: "Coquille Abyssale", Description: "Bénédiction des abysses: +15% résistance magique.", Effet: Effet{ToursRestants: 999, ModifResistMag: 0.15}}
	medaillonChasseurMages = Artefact{Nom: "Médaillon du Chasseur de Mages", Description: "Pourfendeur d'arcanes: +10% dégâts physiques, +10% résistance magique.", Effet: Effet{ToursRestants: 999, ModifDegats: 0.10, ModifResistMag: 0.10}}
	glypheParade           = Artefact{Nom: "Glyphe de Parade", Description: "Maîtrise défensive: +15% armure.", Effet: Effet{ToursRestants: 999, ModifArmure: 0.15}}
	boussoleChasseur       = Artefact{Nom: "Boussole de Chasseur", Description: "Toujours sur la cible: +6% précision, +6% dégâts physiques.", Effet: Effet{ToursRestants: 999, ModifPrecision: 0.06, ModifDegats: 0.06}}
	eclatFoudreGelee       = Artefact{Nom: "Éclat de Foudre Gelée", Description: "Magie stable: +10% dégâts magiques, +3% précision.", Effet: Effet{ToursRestants: 999, ModifDegatsMag: 0.10, ModifPrecision: 0.03}}
	// Dissipation / utilitaires
	antidoteEternel       = Artefact{Nom: "Antidote Éternel", Description: "Supprime le Poison au début de chaque tour.", Effet: Effet{ToursRestants: 999}}
	talismanEteigneflamme = Artefact{Nom: "Talisman Éteigneflamme", Description: "Supprime la Brûlure au début de chaque tour.", Effet: Effet{ToursRestants: 999}}
	sceauHemostatique     = Artefact{Nom: "Sceau Hémostatique", Description: "Supprime le Saignement au début de chaque tour.", Effet: Effet{ToursRestants: 999}}
	pendentifCourage      = Artefact{Nom: "Pendentif de Courage", Description: "Supprime la Peur au début de chaque tour.", Effet: Effet{ToursRestants: 999}}
	talismanVigilance     = Artefact{Nom: "Talisman de Vigilance", Description: "Supprime l'Étourdissement au début de chaque tour.", Effet: Effet{ToursRestants: 999}}
	sceauFocalisation     = Artefact{Nom: "Sceau de Focalisation", Description: "Supprime Nébulation et Défavorisation au début de chaque tour.", Effet: Effet{ToursRestants: 999}}
	glypheBastion         = Artefact{Nom: "Glyphe de Bastion", Description: "Supprime Brise-Armure et Brise-Armure Magique au début de chaque tour.", Effet: Effet{ToursRestants: 999}}
	cachetDetermination   = Artefact{Nom: "Cachet de Détermination", Description: "Supprime les débuffs d'attaque (ex: Peur) au début de chaque tour.", Effet: Effet{ToursRestants: 999}}
)

// Slice des artefacts disponibles construit à partir des variables ci-dessus
var ArtefactsDisponibles = []Artefact{
	insigneSergent,
	gantAntiEmeute,
	talismanBrasier,
	carquoisMilleFleches,
	coutelasRunique,
	dentMammouth,
	anneauTempetes,
	bandeauRonin,
	perleAether,
	peauBananeSacree,
	puceVisee,
	conqueProfondeurs,
	medaillonFoudrePure,
	reliqueSylve,
	pierreIgnition,
	oeilLynx,
	runeTrempe,
	coquilleAbyssale,
	medaillonChasseurMages,
	glypheParade,
	boussoleChasseur,
	eclatFoudreGelee,
	antidoteEternel,
	talismanEteigneflamme,
	sceauHemostatique,
	pendentifCourage,
	talismanVigilance,
	sceauFocalisation,
	glypheBastion,
	cachetDetermination,
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
