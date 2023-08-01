package main

import (
	"context"
	"fmt"
	"gfit/gfit"
	"log"
)

func main() {
	cfg, err := gfit.LoadConfig(gfit.Sleep, gfit.Weight, gfit.Activity)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := gfit.GetClient(context.Background(), cfg)
	if err != nil {
		log.Fatalln(err)
	}
	b, err := gfit.GetSleepGoal(client)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}
