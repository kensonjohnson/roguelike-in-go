package config

// Holds values for the size of elements within the game
type config struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
	UIHeight     int
	FontSize     int
}

var Config *config

// Creates a fully populated Config struct
func init() {
	Config = &config{
		ScreenWidth:  80,
		ScreenHeight: 60,
		TileWidth:    16,
		TileHeight:   16,
		UIHeight:     10,
		FontSize:     16,
	}
}
