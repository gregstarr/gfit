package main

import (
	"encoding/json"
	"gfit/generator"
	"log"
	"os"
)

func main() {
	f, err := os.Open(generator.ApiDefFile)
	if err != nil {
		log.Fatalln("can't open api def", err)
	}
	def := generator.ApiDef{}
	dec := json.NewDecoder(f)
	dec.Decode(&def)
	def.Generate()
}
