package main

import (
	"fmt"
	"log"
	"path/filepath"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	fmt.Println(filepath.Abs("."))
	fmt.Println(filepath.Abs("ent/schema"))
	err := entc.Generate("./ent/schema", &gen.Config{
		Templates: entgql.AllTemplates,
		Target:    "ent",
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
