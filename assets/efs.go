package assets

import (
	"bytes"
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kensonjohnson/roguelike-game-go/config"
	"github.com/kensonjohnson/roguelike-game-go/internal/logger"
)

var (
	//go:embed "*"
	assetsFS embed.FS

	// Tiles
	Floor       *ebiten.Image
	Wall        *ebiten.Image
	StairsUp    *ebiten.Image
	StairsDown  *ebiten.Image
	ChestOpen   *ebiten.Image
	ChestClosed *ebiten.Image

	// UI
	UIPanel               *ebiten.Image
	UIPanelWithMinimap    *ebiten.Image
	KenneyMiniFont        *text.GoTextFace
	KenneyMiniSquaredFont *text.GoTextFace
	KenneyPixelFont       *text.GoTextFace

	// Characters
	Player *ebiten.Image
	Skelly *ebiten.Image
	Orc    *ebiten.Image

	// Weapons
	RustyAxe      *ebiten.Image
	GreatAxe      *ebiten.Image
	RoyalGreatAxe *ebiten.Image

	RustyBattleAxe *ebiten.Image
	BattleAxe      *ebiten.Image
	RoyalBattleAxe *ebiten.Image

	Club       *ebiten.Image
	HeavyClub  *ebiten.Image
	SpikedClub *ebiten.Image

	RustyDagger *ebiten.Image
	Dagger      *ebiten.Image
	RoyalDagger *ebiten.Image

	Spear      *ebiten.Image
	GreatSpear *ebiten.Image
	RoyalSpear *ebiten.Image

	RustyGreatSword *ebiten.Image
	GreatSword      *ebiten.Image
	RoyalGreatSword *ebiten.Image

	RustyLongSword *ebiten.Image
	LongSword      *ebiten.Image
	RoyalLongSword *ebiten.Image

	RustyShortSword *ebiten.Image
	ShortSword      *ebiten.Image
	RoyalShortSword *ebiten.Image

	Mace      *ebiten.Image
	RoyalMace *ebiten.Image

	Scythe       *ebiten.Image
	ReaperScythe *ebiten.Image

	// Armor

	LinenShirt     *ebiten.Image
	PaddedArmor    *ebiten.Image
	HalfPlateArmor *ebiten.Image
	PlateArmor     *ebiten.Image
	RoyalArmor     *ebiten.Image

	ClothGloves     *ebiten.Image
	LeatherGloves   *ebiten.Image
	ChainmailGloves *ebiten.Image
	RoyalGloves     *ebiten.Image

	LeatherCap   *ebiten.Image
	Helmet       *ebiten.Image
	KnightHelmet *ebiten.Image
	RoyalHelmet  *ebiten.Image

	WornBoots    *ebiten.Image
	LeatherBoots *ebiten.Image
	KnightBoots  *ebiten.Image
	RoyalBoots   *ebiten.Image

	// Shields
	WoodenShield     *ebiten.Image
	WoodenTallShield *ebiten.Image
	Shield           *ebiten.Image
	SpikedShield     *ebiten.Image
	HeavyShield      *ebiten.Image
	RoyalShield      *ebiten.Image

	// Consumables
	HealthPotion      *ebiten.Image
	GreatHealthPotion *ebiten.Image
	RoyalHealthPotion *ebiten.Image

	Apple  *ebiten.Image
	Bread  *ebiten.Image
	Carrot *ebiten.Image
	Cheese *ebiten.Image
	Egg    *ebiten.Image
	Fish   *ebiten.Image
	Ham    *ebiten.Image
	Milk   *ebiten.Image
	Pear   *ebiten.Image
	Steak  *ebiten.Image

	// World Items / Pickups
	WorldAlcohol           *ebiten.Image
	WorldSmallCoin         *ebiten.Image
	WorldCoinStack         *ebiten.Image
	WorldHealthPotion      *ebiten.Image
	WorldGreatHealthPotion *ebiten.Image
	WorldRoyalHealthPotion *ebiten.Image
)

// Loads all required assets, panics if any one fails.
func MustLoadAssets() {
	/*-----------------------
	--------- Tiles ---------
	-----------------------*/
	Floor = mustLoadImage("images/tiles/floor.png")
	Wall = mustLoadImage("images/tiles/wall.png")
	StairsUp = mustLoadImage("images/tiles/stairs-up.png")
	StairsDown = mustLoadImage("images/tiles/stairs-down.png")
	ChestOpen = mustLoadImage("images/tiles/chest_open.png")
	ChestClosed = mustLoadImage("images/tiles/chest_closed.png")

	/*-----------------------
	---------- UI -----------
	-----------------------*/
	UIPanel = mustLoadImage("images/ui/UIPanel.png")
	UIPanelWithMinimap = mustLoadImage("images/ui/UIPanelWithMinimap.png")

	kenneyMiniFontBytes, err := assetsFS.ReadFile("fonts/KenneyMini.ttf")
	if err != nil {
		logger.Fatal(err)
	}
	KenneyMiniFont = mustLoadFont(kenneyMiniFontBytes)
	kenneyMiniSquaredFontBytes, err := assetsFS.ReadFile("fonts/KenneyMiniSquared.ttf")
	if err != nil {
		logger.Fatal(err)
	}
	KenneyMiniSquaredFont = mustLoadFont(kenneyMiniSquaredFontBytes)
	kenneyPixelFontBytes, err := assetsFS.ReadFile("fonts/KenneyPixel.ttf")
	if err != nil {
		logger.Fatal(err)
	}
	KenneyPixelFont = mustLoadFont(kenneyPixelFontBytes)
	// For some reason, the KenneyPixel shows up as half the size of the other fonts.
	KenneyPixelFont.Size = float64(config.FontSize) * 1.5

	/*-----------------------
	------ Characters -------
	-----------------------*/
	Player = mustLoadImage("images/characters/player.png")
	Skelly = mustLoadImage("images/enemies/skelly.png")
	Orc = mustLoadImage("images/enemies/orc.png")

	/*-----------------------
	------ Inventory --------
	-----------------------*/
	// Weapons
	RustyAxe = mustLoadImage("images/items/weapons/rusty_axe.png")
	GreatAxe = mustLoadImage("images/items/weapons/great_axe.png")
	RoyalGreatAxe = mustLoadImage("images/items/weapons/royal_great_axe.png")

	RustyBattleAxe = mustLoadImage("images/items/weapons/rusty_battle_axe.png")
	BattleAxe = mustLoadImage("images/items/weapons/battle_axe.png")
	RoyalBattleAxe = mustLoadImage("images/items/weapons/royal_battle_axe.png")

	Club = mustLoadImage("images/items/weapons/club.png")
	HeavyClub = mustLoadImage("images/items/weapons/heavy_club.png")
	SpikedClub = mustLoadImage("images/items/weapons/spiked_club.png")

	RustyDagger = mustLoadImage("images/items/weapons/rusty_dagger.png")
	Dagger = mustLoadImage("images/items/weapons/dagger.png")
	RoyalDagger = mustLoadImage("images/items/weapons/royal_dagger.png")

	Spear = mustLoadImage("images/items/weapons/spear.png")
	GreatSpear = mustLoadImage("images/items/weapons/great_spear.png")
	RoyalSpear = mustLoadImage("images/items/weapons/royal_spear.png")

	RustyGreatSword = mustLoadImage("images/items/weapons/rusty_great_sword.png")
	GreatSword = mustLoadImage("images/items/weapons/great_sword.png")
	RoyalGreatSword = mustLoadImage("images/items/weapons/royal_great_sword.png")

	RustyLongSword = mustLoadImage("images/items/weapons/rusty_long_sword.png")
	LongSword = mustLoadImage("images/items/weapons/long_sword.png")
	RoyalLongSword = mustLoadImage("images/items/weapons/royal_long_sword.png")

	RustyLongSword = mustLoadImage("images/items/weapons/rusty_short_sword.png")
	LongSword = mustLoadImage("images/items/weapons/short_sword.png")
	RoyalLongSword = mustLoadImage("images/items/weapons/royal_short_sword.png")

	Mace = mustLoadImage("images/items/weapons/mace.png")
	RoyalMace = mustLoadImage("images/items/weapons/royal_mace.png")

	Scythe = mustLoadImage("images/items/weapons/scythe.png")
	ReaperScythe = mustLoadImage("images/items/weapons/reaper_scythe.png")

	// Armor
	LinenShirt = mustLoadImage("images/items/armor/linen_shirt.png")
	PaddedArmor = mustLoadImage("images/items/armor/padded_armor.png")
	HalfPlateArmor = mustLoadImage("images/items/armor/half_plate_armor.png")
	PlateArmor = mustLoadImage("images/items/armor/plate_armor.png")
	RoyalArmor = mustLoadImage("images/items/armor/royal_armor.png")

	ClothGloves = mustLoadImage("images/items/armor/cloth_gloves.png")
	LeatherGloves = mustLoadImage("images/items/armor/leather_gloves.png")
	ChainmailGloves = mustLoadImage("images/items/armor/chainmail_gloves.png")
	RoyalGloves = mustLoadImage("images/items/armor/royal_gloves.png")

	LeatherCap = mustLoadImage("images/items/armor/leather_cap.png")
	Helmet = mustLoadImage("images/items/armor/helmet.png")
	KnightHelmet = mustLoadImage("images/items/armor/knight_helmet.png")
	RoyalHelmet = mustLoadImage("images/items/armor/royal_helmet.png")

	WornBoots = mustLoadImage("images/items/armor/worn_boots.png")
	LeatherBoots = mustLoadImage("images/items/armor/leather_boots.png")
	KnightBoots = mustLoadImage("images/items/armor/knight_boots.png")
	RoyalBoots = mustLoadImage("images/items/armor/royal_boots.png")

	// Shields
	WoodenShield = mustLoadImage("images/items/shields/wooden_shield.png")
	WoodenTallShield = mustLoadImage("images/items/shields/wooden_tall_shield.png")
	Shield = mustLoadImage("images/items/shields/shield.png")
	SpikedShield = mustLoadImage("images/items/shields/spiked_shield.png")
	HeavyShield = mustLoadImage("images/items/shields/heavy_shield.png")
	RoyalShield = mustLoadImage("images/items/shields/royal_shield.png")

	// Consumables
	HealthPotion = mustLoadImage("images/items/consumables/health_potion.png")
	GreatHealthPotion = mustLoadImage("images/items/consumables/great_health_potion.png")
	RoyalHealthPotion = mustLoadImage("images/items/consumables/royal_health_potion.png")

	Apple = mustLoadImage("images/items/consumables/food/apple.png")
	Bread = mustLoadImage("images/items/consumables/food/bread.png")
	Carrot = mustLoadImage("images/items/consumables/food/carrot.png")
	Cheese = mustLoadImage("images/items/consumables/food/cheese.png")
	Egg = mustLoadImage("images/items/consumables/food/egg.png")
	Fish = mustLoadImage("images/items/consumables/food/fish.png")
	Ham = mustLoadImage("images/items/consumables/food/ham.png")
	Milk = mustLoadImage("images/items/consumables/food/milk.png")
	Pear = mustLoadImage("images/items/consumables/food/pear.png")
	Steak = mustLoadImage("images/items/consumables/food/steak.png")

	/*-----------------------
	------ World Items ------
	-----------------------*/
	WorldAlcohol = mustLoadImage("images/items/world_alcohol.png")
	WorldSmallCoin = mustLoadImage("images/items/world_small_coin.png")
	WorldCoinStack = mustLoadImage("images/items/world_coin_stack.png")
	WorldHealthPotion = mustLoadImage("images/items/world_health_potion.png")
	WorldGreatHealthPotion = mustLoadImage("images/items/world_great_health_potion.png")
	WorldRoyalHealthPotion = mustLoadImage("images/items/world_royal_health_potion.png")

}

// Loads image at specified path, panics if it fails.
func mustLoadImage(filePath string) *ebiten.Image {
	imgSource, err := assetsFS.ReadFile(filePath)
	if err != nil {
		logger.Fatal(err)
	}
	image, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(imgSource))
	if err != nil {
		logger.Fatal(err)
	}
	return image
}

// Loads font at specified path, panics if it fails.
func mustLoadFont(font []byte) *text.GoTextFace {
	source, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		logger.Fatal(err)
	}
	return &text.GoTextFace{
		Source: source,
		Size:   float64(config.FontSize),
	}
}
