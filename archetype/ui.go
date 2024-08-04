package archetype

import (
	"github.com/kensonjohnson/roguelike-game-go/component"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/yohamta/donburi"
)

var UITag = donburi.NewTag("ui")

func CreateNewUI(world donburi.World) {
	entity := world.Create(
		UITag,
		component.UI,
	)
	entry := world.Entry(entity)

	ui := component.UIData{
		MessageBox: component.UserMessageBoxData{},
		PlayerHUD:  component.PlayerHUDData{},
	}

	// Configure message box
	topPixel := (config.ScreenHeight - config.UIHeight) * config.TileHeight
	ui.MessageBox.Position = component.PositionData{
		X: 0,
		Y: topPixel,
	}
	ui.MessageBox.FontX = config.FontSize
	ui.MessageBox.FontY = topPixel + 10
	ui.MessageBox.LastText = make([]string, 0, 5)

	// Configure player HUD
	playerEntry := PlayerTag.MustFirst(world)
	ui.PlayerHUD.Position = component.PositionData{
		X: config.ScreenWidth * config.TileWidth / 2,
		Y: topPixel,
	}
	ui.PlayerHUD.FontX = ui.PlayerHUD.Position.X + config.FontSize
	ui.PlayerHUD.FontY = topPixel + 12
	ui.PlayerHUD.Health = component.Health.Get(playerEntry)
	ui.PlayerHUD.Armor = component.Armor.Get(playerEntry)
	ui.PlayerHUD.Weapon = component.Weapon.Get(playerEntry)

	component.UI.SetValue(entry, ui)
}
