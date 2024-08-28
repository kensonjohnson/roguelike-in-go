package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/event"
	"github.com/kensonjohnson/roguelike-game-go/items/armors"
	"github.com/kensonjohnson/roguelike-game-go/items/weapons"
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

	weaponId := component.ItemId.Get(currentPlayerEquipment.Weapon)
	armorId := component.ItemId.Get(currentPlayerEquipment.Armor)

	playerEntry := archetype.PlayerTag.MustFirst(newWorld)
	component.Health.SetValue(playerEntry, *currentPlayerHealth)
	component.Equipment.SetValue(playerEntry, component.EquipmentData{
		Weapon: archetype.CreateNewWeapon(newWorld, weapons.WeaponId(weaponId.Id)),
		// Sheild: currentPlayerEquipment.Sheild,
		// Gloves: currentPlayerEquipment.Gloves,
		Armor: archetype.CreateNewArmor(newWorld, armors.ArmorId(armorId.Id)),
		// Boots:  currentPlayerEquipment.Boots,
	})

}
