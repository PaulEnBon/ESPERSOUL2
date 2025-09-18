package main

// Mapping des noms d'affichage personnalisés pour les salles
// Vous pouvez modifier librement ces noms.
var mapDisplayNames = map[string]string{
	"salle1":  "Sanctuaire du Réveil",
	"salle2":  "Forêt des Murmures",
	"salle3":  "Arène Maudite",
	"salle4":  "Marché Souterrain",
	"salle5":  "Forge des Titans",
	"salle6":  "Salle aux Coffres",
	"salle7":  "Casino des Ombres",
	"salle8":  "Antre des Créateurs",
	"salle9":  "Terres Sauvages",
	"salle10": "Crypte des Âmes",
	"salle11": "Sanctuaire de Repos",
	"salle12": "Colisée des Épreuves",
	"salle13": "Forteresse de l'Ascension",
	"salle14": "Trône des Ténèbres",
	"salle15": "Chambre des Plaisirs",
}

// Retourne le nom d'affichage pour une carte, ou son identifiant si non mappée
func displayNameFor(mapName string) string {
	if n, ok := mapDisplayNames[mapName]; ok && n != "" {
		return n
	}
	return mapName
}
