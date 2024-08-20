package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/event"
	"github.com/kensonjohnson/roguelike-game-go/layer"
	"github.com/kensonjohnson/roguelike-game-go/system"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type Level struct {
	ecs ecs.ECS
}

func (level *Level) Update() {
	level.ecs.Update()
	event.ProgressLevelEvent.ProcessEvents(level.ecs.World)

}

func (level *Level) Draw(screen *ebiten.Image) {
	level.ecs.Draw(screen)
}

func CreateFirstLevel() *Level {
	level := &Level{}
	level.ecs = *ecs.NewECS(createWorld())

	// Add systems
	level.ecs.AddSystem(system.Camera.Update)
	level.ecs.AddSystem(system.Turn.Update)
	level.ecs.AddSystem(system.UI.Update)

	// Add renderers
	level.ecs.AddRenderer(layer.Background, system.Render.DrawBackground)
	level.ecs.AddRenderer(layer.Foreground, system.Render.Draw)
	level.ecs.AddRenderer(layer.UI, system.UI.Draw)
	level.ecs.AddRenderer(layer.UI, system.DrawMinimap)
	if system.Debug.On {
		level.ecs.AddRenderer(layer.UI, system.Debug.Draw)
	}

	// Add event listeners
	event.ProgressLevelEvent.Subscribe(level.ecs.World, progressLevel)

	return level
}

func createWorld() donburi.World {
	world := donburi.NewWorld()

	// Create current level
	level := archetype.GenerateLevel(world)

	for index, room := range level.Rooms {
		if index == 0 {
			archetype.CreateNewPlayer(world)
		} else {
			archetype.CreateMonster(world, level, room)
		}
	}

	// Create the UI
	archetype.CreateNewUI(world)

	// Create the camera
	archetype.CreateNewCamera(world)

	return world
}

func progressLevel(world donburi.World, eventData event.ProgressLevel) {

	// Grab the current player's data
	playerEntry := archetype.PlayerTag.MustFirst(world)
	playerHealth := component.Health.Get(playerEntry)
	playerWeapon := component.Weapon.Get(playerEntry)
	playerArmor := component.Armor.Get(playerEntry)

	// Create a new world
	newWorld := createWorld()

	// Apply the player's data to the new world
	copyPlayerInstance(newWorld, playerHealth, playerWeapon, playerArmor)

	level := &Level{}
	level.ecs = *ecs.NewECS(newWorld)
	level.ecs.AddSystem(system.Camera.Update)
	level.ecs.AddSystem(system.Turn.Update)
	level.ecs.AddSystem(system.UI.Update)
	level.ecs.AddRenderer(layer.Background, system.Render.DrawBackground)
	level.ecs.AddRenderer(layer.Foreground, system.Render.Draw)
	level.ecs.AddRenderer(layer.UI, system.UI.Draw)
	level.ecs.AddRenderer(layer.UI, system.DrawMinimap)
	if system.Debug.On {
		level.ecs.AddRenderer(layer.UI, system.Debug.Draw)
	}

	event.ProgressLevelEvent.Subscribe(level.ecs.World, progressLevel)

	SceneManager.GoTo(level)
}

func copyPlayerInstance(
	newWorld donburi.World,
	health *component.HealthData,
	weapon *component.WeaponData,
	armor *component.ArmorData,
) {
	playerEntry := archetype.PlayerTag.MustFirst(newWorld)
	component.Health.SetValue(playerEntry, component.HealthData{MaxHealth: health.MaxHealth, CurrentHealth: health.CurrentHealth})
	component.Weapon.SetValue(playerEntry, component.WeaponData{
		Name:          weapon.Name,
		ActionText:    weapon.ActionText,
		MinimumDamage: weapon.MinimumDamage,
		MaximumDamage: weapon.MaximumDamage,
		ToHitBonus:    weapon.ToHitBonus,
	})
	component.Armor.SetValue(playerEntry, component.ArmorData{
		Name:       armor.Name,
		Defense:    armor.Defense,
		ArmorClass: armor.ArmorClass,
	})

}
