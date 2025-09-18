package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Path to saves directory
const savesDir = "Play"

// Remember last used save slot for in-game quick save
var lastUsedSaveSlot = -1

// Optional spawn override for load
var pendingSpawn struct {
	Active bool
	X      int
	Y      int
}

type bossRoomsStateSnapshot struct {
	DefeatedMini map[string]bool
	Fragments    int
	SpawnerSpawn bool
	BossDefeated bool
}

type SaveData struct {
	Version int       `json:"version"`
	Slot    int       `json:"slot"`
	When    time.Time `json:"when"`

	Player          Personnage                 `json:"player"`
	PlayerInventory map[string]int             `json:"playerInventory"`
	CurrentMap      string                     `json:"currentMap"`
	PlayerX         int                        `json:"playerX"`
	PlayerY         int                        `json:"playerY"`
	EnemiesDefeated map[string]map[string]bool `json:"enemiesDefeated"`
	PnjTransformed  map[string]map[string]bool `json:"pnjTransformed"`
	RewardsGiven    map[string]map[string]bool `json:"rewardsGiven"`
	ChestOpened     map[string]bool            `json:"chestOpened"`
	SecretChests    map[string]bool            `json:"secretChests"`
	StoneBroken     bool                       `json:"stoneBroken"`
	CutTrees        map[string]map[string]bool `json:"cutTrees"`
	MapTrees        map[string][]TilePlacement `json:"mapTrees"`

	// Boss systems
	Salle12 struct {
		DefeatedMini map[string]bool
		Fragments    int
		SpawnerSpawn bool
		BossDefeated bool
	} `json:"salle12"`
	BossRooms map[string]bossRoomsStateSnapshot `json:"bossRooms"`

	// Dynamic mobs and flags
	RandomMobsSalle10 []struct{ X, Y int }       `json:"randomMobsSalle10"`
	SuperEnemyFlags   map[string]map[string]bool `json:"superEnemyFlags"`
}

func slotFilePath(slot int) string {
	return filepath.Join(savesDir, fmt.Sprintf("slot%d.json", slot))
}

func ensureSavesDir() error {
	return os.MkdirAll(savesDir, 0o755)
}

func firstEmptySlot() int {
	for i := 1; i <= 4; i++ {
		if _, err := os.Stat(slotFilePath(i)); os.IsNotExist(err) {
			return i
		}
	}
	return -1
}

// SaveToSlot persists the current game state into the given slot (1..4).
func SaveToSlot(slot int) error {
	if slot < 1 || slot > 4 {
		return fmt.Errorf("slot invalide: %d", slot)
	}
	if err := ensureSavesDir(); err != nil {
		return err
	}

	// Derive current position
	px, py := -1, -1
	if mapDataGlobalRef != nil {
		px, py = findPlayer(mapDataGlobalRef)
	}
	// Fallback positions if not running loop yet
	if px == -1 || py == -1 {
		// Default spawn in salle1
		px, py = 8, 5
	}

	// Snapshot boss rooms state
	brStates := make(map[string]bossRoomsStateSnapshot)
	for name, cfg := range bossRooms {
		st := cfg.state
		snap := bossRoomsStateSnapshot{
			DefeatedMini: map[string]bool{},
			Fragments:    st.fragments,
			SpawnerSpawn: st.spawnerSpawn,
			BossDefeated: st.bossDefeated,
		}
		for k, v := range st.defeatedMini {
			snap.DefeatedMini[k] = v
		}
		brStates[name] = snap
	}

	// Snapshot salle12
	salle12Snap := struct {
		DefeatedMini map[string]bool
		Fragments    int
		SpawnerSpawn bool
		BossDefeated bool
	}{
		DefeatedMini: map[string]bool{},
		Fragments:    salle12BossState.fragments,
		SpawnerSpawn: salle12BossState.spawnerSpawn,
		BossDefeated: salle12BossState.bossDefeated,
	}
	for k, v := range salle12BossState.defeatedMini {
		salle12Snap.DefeatedMini[k] = v
	}

	// Random mobs salle10 snapshot
	var mobs10 []struct{ X, Y int }
	for _, m := range randomMobsSalle10 {
		mobs10 = append(mobs10, struct{ X, Y int }{X: m.x, Y: m.y})
	}

	// Deep copy helpers for maps
	copy2DMap := func(src map[string]map[string]bool) map[string]map[string]bool {
		out := make(map[string]map[string]bool, len(src))
		for k, v := range src {
			inner := make(map[string]bool, len(v))
			for ik, iv := range v {
				inner[ik] = iv
			}
			out[k] = inner
		}
		return out
	}

	sd := SaveData{
		Version:           1,
		Slot:              slot,
		When:              time.Now(),
		Player:            currentPlayer,
		PlayerInventory:   map[string]int{},
		CurrentMap:        currentMapGlobalRef,
		PlayerX:           px,
		PlayerY:           py,
		EnemiesDefeated:   copy2DMap(enemiesDefeated),
		PnjTransformed:    copy2DMap(pnjTransformed),
		RewardsGiven:      copy2DMap(rewardsGiven),
		ChestOpened:       map[string]bool{},
		SecretChests:      map[string]bool{},
		StoneBroken:       stoneBroken,
		CutTrees:          copy2DMap(cutTrees),
		MapTrees:          mapTrees,
		Salle12:           salle12Snap,
		BossRooms:         brStates,
		RandomMobsSalle10: mobs10,
		SuperEnemyFlags:   superEnemyFlags,
	}
	for k, v := range playerInventory {
		sd.PlayerInventory[k] = v
	}
	for k, v := range chestOpened {
		sd.ChestOpened[k] = v
	}
	for k, v := range secretChestsOpened {
		sd.SecretChests[k] = v
	}

	b, err := json.MarshalIndent(sd, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(slotFilePath(slot), b, 0o644); err != nil {
		return err
	}
	lastUsedSaveSlot = slot
	addHUDMessage(fmt.Sprintf("💾 Sauvegardé dans l'emplacement %d", slot))
	return nil
}

// Save initial state for a new game into the first empty slot; returns chosen slot or -1.
func SaveInitialNewGame(currentMap string, px, py int) int {
	slot := firstEmptySlot()
	if slot == -1 {
		return -1
	}
	// Prepare globals minimally for the save call
	prevMap := currentMapGlobalRef
	currentMapGlobalRef = currentMap
	defer func() { currentMapGlobalRef = prevMap }()
	if err := SaveToSlot(slot); err != nil {
		return -1
	}
	lastUsedSaveSlot = slot
	return slot
}

// LoadFromSlot loads state and launches the game loop from the saved map.
func LoadFromSlot(appStop func(), slot int) error {
	if slot < 1 || slot > 4 {
		return fmt.Errorf("slot invalide: %d", slot)
	}
	data, err := os.ReadFile(slotFilePath(slot))
	if err != nil {
		return err
	}
	var sd SaveData
	if err := json.Unmarshal(data, &sd); err != nil {
		return err
	}
	// Restore basics
	currentPlayer = sd.Player
	playerInventory = map[string]int{}
	for k, v := range sd.PlayerInventory {
		playerInventory[k] = v
	}
	enemiesDefeated = sd.EnemiesDefeated
	pnjTransformed = sd.PnjTransformed
	rewardsGiven = sd.RewardsGiven
	chestOpened = sd.ChestOpened
	secretChestsOpened = sd.SecretChests
	stoneBroken = sd.StoneBroken
	cutTrees = sd.CutTrees
	mapTrees = sd.MapTrees
	// Salle12 state
	salle12BossState.defeatedMini = sd.Salle12.DefeatedMini
	salle12BossState.fragments = sd.Salle12.Fragments
	salle12BossState.spawnerSpawn = sd.Salle12.SpawnerSpawn
	salle12BossState.bossDefeated = sd.Salle12.BossDefeated
	// Boss rooms state (preserve layout/codes)
	for name, snap := range sd.BossRooms {
		if cfg, ok := bossRooms[name]; ok {
			cfg.state.defeatedMini = snap.DefeatedMini
			cfg.state.fragments = snap.Fragments
			cfg.state.spawnerSpawn = snap.SpawnerSpawn
			cfg.state.bossDefeated = snap.BossDefeated
		}
	}
	// Random mobs and flags
	randomMobsSalle10 = []struct{ x, y int }{}
	for _, m := range sd.RandomMobsSalle10 {
		randomMobsSalle10 = append(randomMobsSalle10, struct{ x, y int }{m.X, m.Y})
	}
	superEnemyFlags = sd.SuperEnemyFlags

	lastUsedSaveSlot = slot
	// Set pending spawn and run the loop
	pendingSpawn = struct {
		Active bool
		X      int
		Y      int
	}{Active: true, X: sd.PlayerX, Y: sd.PlayerY}

	if appStop != nil {
		appStop()
	}
	RunGameLoop(sd.CurrentMap)
	return nil
}

// Delete a slot save file
func DeleteSlot(slot int) error {
	if slot < 1 || slot > 4 {
		return fmt.Errorf("slot invalide: %d", slot)
	}
	if err := os.Remove(slotFilePath(slot)); err != nil {
		return err
	}
	if lastUsedSaveSlot == slot {
		lastUsedSaveSlot = -1
	}
	return nil
}

// Read a slot summary for UI display
func ReadSlotSummary(slot int) (exists bool, line string) {
	p := slotFilePath(slot)
	data, err := os.ReadFile(p)
	if err != nil {
		return false, fmt.Sprintf("Slot %d: Vide", slot)
	}
	var sd SaveData
	if err := json.Unmarshal(data, &sd); err != nil {
		return true, fmt.Sprintf("Slot %d: [corrompu]", slot)
	}
	name := sd.Player.Nom
	when := sd.When.Format("2006-01-02 15:04")
	line = fmt.Sprintf("Slot %d: %s @ %s (%s)", slot, name, sd.CurrentMap, when)
	return true, line
}
