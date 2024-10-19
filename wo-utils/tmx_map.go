package woutils

import (
	"encoding/xml"
	"log"
	"strconv"
	"strings"
)

type TsxTileSet struct {
	Version      string `xml:"version,attr"`
	TiledVersion string `xml:"tiledversion,attr"`
	Name         string `xml:"name,attr"`
	TileWidth    int    `xml:"tilewidth,attr"`
	TileHeight   int    `xml:"tileheight,attr"`
	TileCount    int    `xml:"tilecount,attr"`
	Columns      int    `xml:"columns,attr"`
	Image        struct {
		Source string `xml:"source,attr"`
		Width  int    `xml:"width,attr"`
		Height int    `xml:"height,attr"`
	} `xml:"image"`
}

type TmxMap struct {
	XMLName      xml.Name `xml:"map"`
	Version      string   `xml:"version,attr"`
	TiledVersion string   `xml:"tiledversion,attr"`
	Orientation  string   `xml:"orientation,attr"`
	RenderOrder  string   `xml:"renderorder,attr"`
	Width        int      `xml:"width,attr"`
	Height       int      `xml:"height,attr"`
	TileWidth    int      `xml:"tilewidth,attr"`
	TileHeight   int      `xml:"tileheight,attr"`
	Infinite     int      `xml:"infinite,attr"`
	NextLayerId  int      `xml:"nextlayerid,attr"`
	NextObjectId int      `xml:"nextobjectid,attr"`

	TileSets []struct {
		FirstGid int    `xml:"firstgid,attr"`
		Source   string `xml:"source,attr"`
		TsxPath  string
		TsxData  *TsxTileSet
	} `xml:"tileset"`
	Layers []struct {
		Id     int    `xml:"id,attr"`
		Name   string `xml:"name,attr"`
		Width  int    `xml:"width,attr"`
		Height int    `xml:"height,attr"`
		Data   struct {
			Encoding string `xml:"encoding,attr"`
			Tiles    []int  `xml:"-"` // Filled with data from the Data.Content after UnmarshalXML
			Content  string `xml:",chardata"`
		} `xml:"data"`
	} `xml:"layer"`
}

type TiledMap struct {
	path   string
	TmxMap TmxMap
}

func NewTiledMap(path string) TiledMap {
	var tmxMap TmxMap
	if err := ReadXml(path, &tmxMap); err != nil {
		log.Fatalln(err)
	}

	if err := processTiles(&tmxMap); err != nil {
		log.Fatalln("Failed to process tilemap: ", err)
	}

	for tilesetIndex := range tmxMap.TileSets {
		tileset := &tmxMap.TileSets[tilesetIndex] // Using pointer to update the original struct

		if tileset.Source != "" {
			tileset.TsxPath = AppendOnPath(GetDirFromPath(path), tileset.Source)
			var tsxTileSet TsxTileSet
			if err := ReadXml(tileset.TsxPath, &tsxTileSet); err != nil {
				log.Fatalln(err)
			}

			tileset.TsxData = &tsxTileSet
		}
	}

	return TiledMap{
		path:   path,
		TmxMap: tmxMap,
	}
}

func processTiles(tmxMap *TmxMap) error {
	for layerIndex := range tmxMap.Layers {
		layer := &tmxMap.Layers[layerIndex] // Using pointer to update the original struct

		if layer.Data.Encoding == "csv" {
			content := strings.TrimSpace(layer.Data.Content)
			tileStrings := strings.Split(content, ",")

			tiles := make([]int, 0, len(tileStrings))

			for _, tile := range tileStrings {
				tile = strings.TrimSpace(tile)
				if tile == "" {
					continue
				}

				tileID, err := strconv.Atoi(tile)
				if err != nil {
					return err
				}
				tiles = append(tiles, tileID)
			}

			layer.Data.Tiles = tiles
		}
	}
	return nil
}
