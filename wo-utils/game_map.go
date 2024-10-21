package woutils

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type GameMapLayer struct {
	layerName string
	tiles     []int32 // Tiles ID
}

type GameMapTileSet struct {
	minTileId         int32
	maxTileId         int32
	texture           *sdl.Texture
	textureSourcePath string
	columns           int32
	tileWidth         int32
	tileHeight        int32
}

type GameMap struct {
	tileWidth    int32
	tileHeight   int32
	mapWidth     int32
	mapHeight    int32
	layers       map[int]GameMapLayer
	tileSets     map[int]*GameMapTileSet // Maps tileset firstgid to tileset
	translationX int32
	translationY int32
	zoom         float32
}

func NewGameMap(renderer *sdl.Renderer, mapName string, tmxFilePath string) GameMap {
	tileMap := NewTiledMap(tmxFilePath)

	tileSets := make(map[int]*GameMapTileSet)
	for index, tileSet := range tileMap.TmxMap.TileSets {
		tileSets[index] = &GameMapTileSet{
			minTileId:         int32(tileSet.FirstGid),
			maxTileId:         int32(tileSet.FirstGid + tileSet.TsxData.TileCount - 1),
			texture:           nil,
			textureSourcePath: AppendOnPath(GetDirFromPath(tileSet.TsxPath), tileSet.TsxData.Image.Source),
			tileWidth:         int32(tileSet.TsxData.TileWidth),
			tileHeight:        int32(tileSet.TsxData.TileHeight),
			columns:           int32(tileSet.TsxData.Columns),
		}
	}

	layers := make(map[int]GameMapLayer)
	for layerIndex := range tileMap.TmxMap.Layers {
		layer := tileMap.TmxMap.Layers[layerIndex] // Using pointer to update the original struct

		layers[layer.Id] = GameMapLayer{
			layerName: layer.Name,
			tiles:     ListIntToListInt32(layer.Data.Tiles),
		}
	}

	for tileSetIndex := range tileSets {
		tileSet := tileSets[tileSetIndex] // Using pointer to update the original struct
		if err := loadTextures(renderer, tileSet); err != nil {
			log.Fatalf("Failed to load textures for tileset ID %d, image_source: %s\n", tileSetIndex, tileSet.textureSourcePath)
		}
	}

	return GameMap{
		tileWidth:    int32(tileMap.TmxMap.TileWidth),
		tileHeight:   int32(tileMap.TmxMap.TileHeight),
		mapWidth:     int32(tileMap.TmxMap.Width),
		mapHeight:    int32(tileMap.TmxMap.Height),
		layers:       layers,
		tileSets:     tileSets,
		translationX: 0,
		translationY: 0,
		zoom:         1,
	}
}

func loadTextures(renderer *sdl.Renderer, tileSet *GameMapTileSet) error {
	if tileSet.texture != nil {
		log.Printf("Texture already loaded for tileset %s\n", tileSet.textureSourcePath)
		return nil
	}

	texture, err := LoadTexture(renderer, tileSet.textureSourcePath)
	if err != nil {
		return err
	}

	tileSet.texture = texture
	return nil
}

func (gm *GameMap) Destroy() {
	for _, tileSet := range gm.tileSets {
		if tileSet.texture != nil {
			tileSet.texture.Destroy()
		}
	}
}

func (gm *GameMap) Translate(x, y int32) {
	gm.translationX += x
	gm.translationY += y
}

// SetZoom alters the zoom in wich the map is rendered.
// 1 is the default zoom (100%).
// Param "zoom" must be between 0.1 and 10.0
func (gm *GameMap) SetZoom(zoom float32) {
	if zoom <= 0.1 || zoom > 10.0 {
		log.Println("Zoom must be between 0.1 (10%) and 10.0 (10x)")
		return
	}
	gm.zoom = zoom
}

func rectsOverlap(a, b *sdl.Rect) bool {
	return a.X+a.W > b.X && a.X < b.X+b.W && a.Y+a.H > b.Y && a.Y < b.Y+b.H
}

func (gm *GameMap) Render(renderer *sdl.Renderer) {
	var currentTileset *GameMapTileSet

	prevScaleX, prevScaleY := renderer.GetScale()
	renderer.SetScale(gm.zoom, gm.zoom)

	offsetX, offsetY := gm.tileWidth, gm.tileHeight
	viewport := renderer.GetViewport()
	viewport.X = viewport.X - offsetX
	viewport.Y = viewport.Y - offsetY
	viewport.W = viewport.W + offsetX
	viewport.H = viewport.H + offsetY

	for _, layer := range gm.layers {
		for i, tileID := range layer.tiles {
			if tileID == 0 {
				continue
			}

			currentTileset = nil

			for _, tileSet := range gm.tileSets {
				if (tileID >= tileSet.minTileId) && (tileID <= tileSet.maxTileId) {
					currentTileset = tileSet
					break
				}
			}

			if currentTileset == nil {
				log.Fatalln("Couldn't find tileset. Are the tilemaps and tilesets properly configured?")
			}

			x := int32(i) % gm.mapWidth * gm.tileWidth
			y := int32(i) / gm.mapWidth * gm.tileHeight

			x, y = CartesianToIsometric(x, y)
			x += gm.translationX
			y += gm.translationY

			tileRect := &sdl.Rect{
				X: x,
				Y: y,
				W: gm.tileWidth,
				H: gm.tileHeight,
			}

			if !rectsOverlap(tileRect, &viewport) {
				continue
			}

			tileSetTileRow := (tileID - currentTileset.minTileId) / currentTileset.columns
			tileSetTileColumn := (tileID - currentTileset.minTileId) % currentTileset.columns

			tileSetRect := &sdl.Rect{
				X: tileSetTileColumn * currentTileset.tileWidth,
				Y: tileSetTileRow * currentTileset.tileHeight,
				W: currentTileset.tileWidth,
				H: currentTileset.tileHeight,
			}

			renderer.Copy(currentTileset.texture, tileSetRect, tileRect)
		}
	}
	renderer.SetScale(prevScaleX, prevScaleY)
}
