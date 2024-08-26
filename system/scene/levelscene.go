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
	currentPlayerWeapon := component.Weapon.Get(currentPlayerEntry)
	currentPlayerArmor := component.Armor.Get(currentPlayerEntry)

	playerEntry := archetype.PlayerTag.MustFirst(newWorld)
	component.Health.SetValue(playerEntry, component.HealthData{
		MaxHealth:     currentPlayerHealth.MaxHealth,
		CurrentHealth: currentPlayerHealth.CurrentHealth,
	})
	component.Weapon.SetValue(playerEntry, component.WeaponData{
		Name:          currentPlayerWeapon.Name,
		ActionText:    currentPlayerWeapon.ActionText,
		MinimumDamage: currentPlayerWeapon.MinimumDamage,
		MaximumDamage: currentPlayerWeapon.MaximumDamage,
		ToHitBonus:    currentPlayerWeapon.ToHitBonus,
	})
	component.Armor.SetValue(playerEntry, component.ArmorData{
		Name:       currentPlayerArmor.Name,
		Defense:    currentPlayerArmor.Defense,
		ArmorClass: currentPlayerArmor.ArmorClass,
	})

}
