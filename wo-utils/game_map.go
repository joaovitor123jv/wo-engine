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

func (gmt *GameMapTileSet) getTileRect(tileId int32) sdl.Rect {
	tileId -= gmt.minTileId
	column := tileId % gmt.columns
	row := tileId / gmt.columns

	return sdl.Rect{
		X: column * gmt.tileWidth,
		Y: row * gmt.tileHeight,
		W: gmt.tileWidth,
		H: gmt.tileHeight,
	}
}

type GameMap struct {
	tileWidth  int32
	tileHeight int32
	mapWidth   int32
	mapHeight  int32
	layers     []GameMapLayer
	tileSets   []*GameMapTileSet // Maps tileset firstgid to tileset
}

func NewGameMap(context *GameContext, mapName string, tmxFilePath string) GameMap {
	tileMap := NewTiledMap(tmxFilePath)

	tileSets := make([]*GameMapTileSet, len(tileMap.TmxMap.TileSets))
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

	layers := make([]GameMapLayer, len(tileMap.TmxMap.Layers))
	for layerIndex := range tileMap.TmxMap.Layers {
		layer := tileMap.TmxMap.Layers[layerIndex] // Using pointer to update the original struct

		layers[layerIndex] = GameMapLayer{
			layerName: layer.Name,
			tiles:     layer.Data.Tiles,
		}
	}

	for tileSetIndex := range tileSets {
		tileSet := tileSets[tileSetIndex] // Using pointer to update the original struct
		if err := loadTextures(context.GetRenderer(), tileSet); err != nil {
			log.Fatalf("Failed to load textures for tileset ID %d, image_source: %s\n", tileSetIndex, tileSet.textureSourcePath)
		}
	}

	return GameMap{
		tileWidth:  int32(tileMap.TmxMap.TileWidth),
		tileHeight: int32(tileMap.TmxMap.TileHeight),
		mapWidth:   int32(tileMap.TmxMap.Width),
		mapHeight:  int32(tileMap.TmxMap.Height),
		layers:     layers,
		tileSets:   tileSets,
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

func rectsOverlap(a, b *sdl.Rect) bool {
	return a.X+a.W > b.X && a.X < b.X+b.W && a.Y+a.H > b.Y && a.Y < b.Y+b.H
}

func (gm *GameMap) getTilesetFromTileId(tileId int32) *GameMapTileSet {
	for _, tileSet := range gm.tileSets {
		if (tileId >= tileSet.minTileId) && (tileId <= tileSet.maxTileId) {
			return tileSet
		}
	}
	return nil
}

func (gm *GameMap) getTileCoordinates(index int) (x, y int32) {
	x = int32(index) % gm.mapWidth * gm.tileWidth
	y = int32(index) / gm.mapWidth * gm.tileHeight

	x, y = CartesianToIsometric(x, y)
	return x, y
}

func (gm *GameMap) Render(gc *GameContext) {
	var currentTileset *GameMapTileSet
	renderer := gc.GetRenderer()

	gc.InitRenderZoom()
	defer gc.ResetRenderZoom()

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

			if currentTileset = gm.getTilesetFromTileId(tileID); currentTileset == nil {
				log.Fatalln("Couldn't find tileset. Are the tilemaps and tilesets properly configured?")
			}

			x, y := gm.getTileCoordinates(i)
			tileRect := sdl.Rect{
				X: x,
				Y: y,
				W: gm.tileWidth,
				H: gm.tileHeight,
			}
			gc.Camera.TranslateSDLRect(&tileRect)

			// If tile will be rendered in visible area of the screen
			if rectsOverlap(&tileRect, &viewport) {
				tileSetRect := currentTileset.getTileRect(tileID)
				renderer.Copy(currentTileset.texture, &tileSetRect, &tileRect)
			}
		}
	}
}
