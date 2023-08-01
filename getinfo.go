package main

import (
	"encoding/json"
	"fmt"
	"gfit/generator"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open(generator.ApiDefFile)
	if err != nil {
		log.Fatalln("can't open api def", err)
	}
	def := generator.ApiDef{}
	dec := json.NewDecoder(f)
	dec.Decode(&def)
	for _, tag := range def.Tags {
		fmt.Println(tag.Name)
		for _, reqDef := range def.Paths {
			if !generator.Contains(reqDef.Get.Tags, tag.Name) {
				continue
			}
			for _, p := range reqDef.Get.Parameters {
				if len(p.Enum) > 0 {
					fmt.Println(p.Name, p.Type, p.Enum)
				} else if strings.Contains(strings.ToLower(p.Description), "option") {
					fmt.Println(p.Name, p.Type, p.Description)
				}

			}
		}
		fmt.Println()
	}

}
