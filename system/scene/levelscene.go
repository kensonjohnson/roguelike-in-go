package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/event"
	"github.com/kensonjohnson/roguelike-game-go/internal/logger"
	"github.com/kensonjohnson/roguelike-game-go/items"
	"github.com/kensonjohnson/roguelike-game-go/layer"
	"github.com/kensonjohnson/roguelike-game-go/system"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type LevelScene struct {
	ecs ecs.ECS
}

func (level *LevelScene) Update() {
	level.ecs.Update()
	event.ProgressLevelEvent.ProcessEvents(level.ecs.World)
}

func (level *LevelScene) Draw(screen *ebiten.Image) {
	level.ecs.Draw(screen)
}

func CreateFirstLevel() *LevelScene {
	level := &LevelScene{}
	level.configureECS(createWorld())
	return level
}

func createWorld() donburi.World {
	world := donburi.NewWorld()

	// Create current level
	archetype.GenerateLevel(world)

	// Create the UI
	archetype.CreateNewUI(world)

	// Create the camera
	archetype.CreateNewCamera(world)

	return world
}

func progressLevel(world donburi.World, eventData event.ProgressLevel) {

	// Create a new world
	newWorld := createWorld()

	// Apply the player's data to the new world
	copyPlayerInstance(world, newWorld)

	level := &LevelScene{}
	level.configureECS(newWorld)

	SceneManager.GoTo(level)
}

func (l *LevelScene) configureECS(world donburi.World) {
	l.ecs = *ecs.NewECS(world)
	// Add systems
	l.ecs.AddSystem(system.Camera.Update)
	l.ecs.AddSystem(system.Turn.Update)
	l.ecs.AddSystem(system.UI.Update)

	// Add renderers
	l.ecs.AddRenderer(layer.Background, system.Render.DrawBackground)
	l.ecs.AddRenderer(layer.Foreground, system.Render.Draw)
	l.ecs.AddRenderer(layer.UI, system.UI.Draw)
	l.ecs.AddRenderer(layer.UI, system.DrawMinimap)
	if system.Debug.On {
		l.ecs.AddRenderer(layer.UI, system.Debug.Draw)
	}

	// Add event listeners
	event.ProgressLevelEvent.Subscribe(l.ecs.World, progressLevel)
}

func copyPlayerInstance(
	oldWorld donburi.World,
	newWorld donburi.World,
) {
	currentPlayerEntry := archetype.PlayerTag.MustFirst(oldWorld)
	currentPlayerHealth := component.Health.Get(currentPlayerEntry)
	currentPlayerEquipment := component.Equipment.Get(currentPlayerEntry)

	playerEntry := archetype.PlayerTag.MustFirst(newWorld)
	component.Health.SetValue(playerEntry, *currentPlayerHealth)
	component.Equipment.SetValue(playerEntry, component.EquipmentData{
		Weapon: copyWeapon(newWorld, currentPlayerEquipment.Weapon),
		// Sheild: currentPlayerEquipment.Sheild,
		// Gloves: currentPlayerEquipment.Gloves,
		Armor: copyArmor(newWorld, currentPlayerEquipment.Armor),
		// Boots:  currentPlayerEquipment.Boots,
	})

}

func copyWeapon(newWorld donburi.World, entry *donburi.Entry) *donburi.Entry {
	if entry == nil {
		logger.ErrorLogger.Fatal("Entry is nil when copying weapon data")
	}
	name := component.Name.Get(entry).Value
	sprite := component.Sprite.Get(entry).Image
	if sprite == nil {
		logger.ErrorLogger.Fatal("Sprite missing when copying weapon data")
	}
	actionText := component.ActionText.Get(entry).Value
	attack := component.Attack.Get(entry)
	if attack == nil {
		logger.ErrorLogger.Fatal("Attack data missing when copying weapon data")
	}

	newWeaponData := items.WeaponData{
		ItemData: items.ItemData{
			Name:   name,
			Sprite: sprite,
		},
		ActionText:    actionText,
		MinimumDamage: attack.MinimumDamage,
		MaximumDamage: attack.MaximumDamage,
		ToHitBonus:    attack.ToHitBonus,
	}

	newWeapon := archetype.CreateNewWeapon(newWorld, newWeaponData)
	if newWeapon == nil {
		logger.ErrorLogger.Fatal("Failed to create new weapon when copying weapon data")
	}
	return newWeapon
}

func copyArmor(newWorld donburi.World, entry *donburi.Entry) *donburi.Entry {
	if entry == nil {
		logger.ErrorLogger.Fatal("Entry is nil when copying armor data")
	}
	name := component.Name.Get(entry).Value
	sprite := component.Sprite.Get(entry).Image
	if sprite == nil {
		logger.ErrorLogger.Fatal("Sprite missing when copying armor data")
	}
	defense := component.Defense.Get(entry)
	if defense == nil {
		logger.ErrorLogger.Fatal("Defense data missing when copying armor data")
	}

	newArmorData := items.ArmorData{
		ItemData: items.ItemData{
			Name:   name,
			Sprite: sprite,
		},
		Defense:    defense.Defense,
		ArmorClass: defense.ArmorClass,
	}

	newArmor := archetype.CreateNewArmor(newWorld, newArmorData)
	if newArmor == nil {
		logger.ErrorLogger.Fatal("Failed to create new armor when copying armor data")
	}
	return newArmor
}
