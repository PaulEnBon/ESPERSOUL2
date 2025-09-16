package main

func (p *Personnage) ArmeActuelle() *Arme {
	if p.NiveauArme < len(p.ArmesDisponibles) {
		return &p.ArmesDisponibles[p.NiveauArme]
	}
	return nil // ou arme par défaut
}

func (p *Personnage) ArmureActuelle() *Armure {
	if p.NiveauArmure < len(p.ArmuresDisponibles) {
		return &p.ArmuresDisponibles[p.NiveauArmure]
	}
	return nil // ou armure par défaut
}

func (p *Personnage) AmeliorerArme() bool {
	if p.NiveauArme+1 < len(p.ArmesDisponibles) {
		p.NiveauArme++
		return true
	}
	return false // déjà max
}

func (p *Personnage) AmeliorerArmure() bool {
	if p.NiveauArmure+1 < len(p.ArmuresDisponibles) {
		p.NiveauArmure++
		return true
	}
	return false // déjà max
}

/*
if joueur.AmeliorerArme() {
    fmt.Println("Amélioration de l'arme réussie :", joueur.ArmeActuelle().Nom)
} else {
    fmt.Println("L'arme est déjà au niveau maximum.")
}

if joueur.AmeliorerArmure() {
    fmt.Println("Amélioration de l'armure réussie :", joueur.ArmureActuelle().Nom)
} else {
    fmt.Println("L'armure est déjà au niveau maximum.")
}*/
