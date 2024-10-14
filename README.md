# Roguelike in Go

This is a simple roguelike game written in Go. Original inspiration was the tutorial by [Fat Old Yeti](https://www.fatoldyeti.com/categories/roguelike-tutorial/), now heavily modified and expanded.

You can [try the game here](https://kensonjohnson.github.io/roguelike-in-go/), via GitHub Pages and WebAssembly.

## Features

- Built entirely in [Go](https://go.dev/)
- Uses [Ebiten](https://ebitengine.org/) for graphics and input

## Roadmap

- [x] Randomly generated levels
- [x] Simple combat system
- [x] Basic AI for enemies
- [ ] Inventory system
- [ ] Equipment system
- [ ] Town and NPCs
- [ ] Character progression
- [ ] Quests (Maybe)

## Getting started

To run the game, you need to have Go installed. You can download it from the [official website](https://golang.org/) or install it using your package manager.

Obviously, you also need to clone this repository:

```bash
git clone git@github.com:kensonjohnson/roguelike-in-go.git
```

Fetch the dependencies:

```bash
go mod tidy
```

Then, you can run the game using the following command:

```bash
go run .
```
