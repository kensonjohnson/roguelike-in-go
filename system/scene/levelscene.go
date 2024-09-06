package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kensonjohnson/roguelike-game-go/archetype"
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/event"
	"github.com/kensonjohnson/roguelike-game-go/internal/config"
	"github.com/kensonjohnson/roguelike-game-go/internal/logger"
	"github.com/kensonjohnson/roguelike-game-go/system"
	"github.com/kensonjohnson/roguelike-game-go/system/layer"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type LevelScene struct {
	ecs   ecs.ECS
	ready bool
}

func (ls *LevelScene) Update() {
	ls.ecs.Update()
	event.ProgressLevelEvent.ProcessEvents(ls.ecs.World)
}

func (ls *LevelScene) Draw(screen *ebiten.Image) {
	ls.ecs.Draw(screen)
}

func (ls *LevelScene) Ready() bool {
	return ls.ready
}

func (ls *LevelScene) Setup(world donburi.World) {
	ls.ready = false

	if logger.DebugOn {
		logger.DebugLogger.Println("LevelScene setup")
	}

	go func() {

		levelData := archetype.GenerateLevel(world)

		if _, ok := archetype.UITag.First(world); !ok {
			archetype.CreateNewUI(world)
		}

		playerEntry := archetype.PlayerTag.MustFirst(world)
		playerPosition := component.Position.Get(playerEntry)
		startingRoom := levelData.Rooms[0]
		playerPosition.X, playerPosition.Y = startingRoom.Center()

		playerSprite := component.Sprite.Get(playerEntry)
		playerSprite.OffestX = 0
		playerSprite.OffestY = 0

		component.Fov.Get(playerEntry).
			VisibleTiles.Compute(levelData, playerPosition.X, playerPosition.Y, 8)

		cameraEntry := archetype.CameraTag.MustFirst(world)
		camera := component.Camera.Get(cameraEntry)
		camera.MainCamera.Lerp = false
		camera.MainCamera.LookAt(
			float64((playerPosition.X*config.TileWidth)+config.TileWidth/2),
			float64((playerPosition.Y*config.TileHeight)+config.TileHeight/2),
		)
		camera.MainCamera.Lerp = true

		ls.configureECS(world)

		ls.ready = true
	}()
}

func (ls *LevelScene) Teardown() {
	ls.ready = false

	if logger.DebugOn {
		logger.DebugLogger.Println("LevelScene teardown")
	}

	go func() {
		archetype.LevelTag.MustFirst(ls.ecs.World).Remove()

		for entry := range archetype.MonsterTag.Iter(ls.ecs.World) {
			if logger.DebugOn {
				logger.DebugLogger.Println("Removing entry: ", entry.String())
			}
			entry.Remove()
		}

		for entry := range archetype.PickupTag.Iter(ls.ecs.World) {
			if logger.DebugOn {
				logger.DebugLogger.Println("Removing entry: ", entry.String())
			}
			entry.Remove()
		}

		ls.ready = true
	}()
}

func progressLevel(world donburi.World, eventData event.ProgressLevel) {
	newLevelScene := &LevelScene{}
	SceneManager.GoTo(newLevelScene)
}

func (ls *LevelScene) configureECS(world donburi.World) {
	ls.ecs = *ecs.NewECS(world)
	// Add systems
	ls.ecs.AddSystem(system.Camera.Update)
	ls.ecs.AddSystem(system.Turn.Update)
	ls.ecs.AddSystem(system.UI.Update)

	// Add renderers
	ls.ecs.AddRenderer(layer.Background, system.Render.DrawBackground)
	ls.ecs.AddRenderer(layer.Foreground, system.Render.Draw)
	ls.ecs.AddRenderer(layer.UI, system.UI.Draw)
	ls.ecs.AddRenderer(layer.UI, system.DrawMinimap)
	if system.Debug.On {
		ls.ecs.AddRenderer(layer.UI, system.Debug.Draw)
	}

	// Add event listeners
	event.ProgressLevelEvent.Subscribe(ls.ecs.World, progressLevel)
}

// func copyPlayerInstance(
// 	oldWorld donburi.World,
// 	newWorld donburi.World,
// ) {
// 	currentPlayerEntry := archetype.PlayerTag.MustFirst(oldWorld)
// 	currentPlayerHealth := component.Health.Get(currentPlayerEntry)
// 	currentPlayerEquipment := component.Equipment.Get(currentPlayerEntry)

// 	playerEntry := archetype.PlayerTag.MustFirst(newWorld)
// 	component.Health.SetValue(playerEntry, *currentPlayerHealth)
// 	component.Equipment.SetValue(playerEntry, component.EquipmentData{
// 		Weapon: copyWeapon(newWorld, currentPlayerEquipment.Weapon),
// 		// Sheild: currentPlayerEquipment.Sheild,
// 		// Gloves: currentPlayerEquipment.Gloves,
// 		Armor: copyArmor(newWorld, currentPlayerEquipment.Armor),
// 		// Boots:  currentPlayerEquipment.Boots,
// 	})

// }

// func copyWeapon(newWorld donburi.World, entry *donburi.Entry) *donburi.Entry {
// 	if entry == nil {
// 		logger.ErrorLogger.Panic("Entry is nil when copying weapon data")
// 	}
// 	name := component.Name.Get(entry).Value
// 	sprite := component.Sprite.Get(entry).Image
// 	if sprite == nil {
// 		logger.ErrorLogger.Panic("Sprite missing when copying weapon data")
// 	}
// 	actionText := component.ActionText.Get(entry).Value
// 	attack := component.Attack.Get(entry)
// 	if attack == nil {
// 		logger.ErrorLogger.Panic("Attack data missing when copying weapon data")
// 	}

// 	newWeaponData := items.WeaponData{
// 		ItemData: items.ItemData{
// 			Name:   name,
// 			Sprite: sprite,
// 		},
// 		ActionText:    actionText,
// 		MinimumDamage: attack.MinimumDamage,
// 		MaximumDamage: attack.MaximumDamage,
// 		ToHitBonus:    attack.ToHitBonus,
// 	}

// 	newWeapon := archetype.CreateNewWeapon(newWorld, newWeaponData)
// 	if newWeapon == nil {
// 		logger.ErrorLogger.Panic("Failed to create new weapon when copying weapon data")
// 	}
// 	return newWeapon
// }

// func copyArmor(newWorld donburi.World, entry *donburi.Entry) *donburi.Entry {
// 	if entry == nil {
// 		logger.ErrorLogger.Panic("Entry is nil when copying armor data")
// 	}
// 	name := component.Name.Get(entry).Value
// 	sprite := component.Sprite.Get(entry).Image
// 	if sprite == nil {
// 		logger.ErrorLogger.Panic("Sprite missing when copying armor data")
// 	}
// 	defense := component.Defense.Get(entry)
// 	if defense == nil {
// 		logger.ErrorLogger.Panic("Defense data missing when copying armor data")
// 	}

// 	newArmorData := items.ArmorData{
// 		ItemData: items.ItemData{
// 			Name:   name,
// 			Sprite: sprite,
// 		},
// 		Defense:    defense.Defense,
// 		ArmorClass: defense.ArmorClass,
// 	}

// 	newArmor := archetype.CreateNewArmor(newWorld, newArmorData)
// 	if newArmor == nil {
// 		logger.ErrorLogger.Panic("Failed to create new armor when copying armor data")
// 	}
// 	return newArmor
// }
