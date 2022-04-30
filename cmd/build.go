package cmd

import (
	"fmt"

	"github.com/tatthien/giraffe/engine"
)

func Build() {
	engine := engine.New()

	engine.ScanContent()

	engine.GenerateIndexPage()
	engine.GenerateTagIndexPage()
	engine.GenerateSingluarPages()
	engine.GenerateTagPages()
	engine.GenerateRSS()
	engine.GeneratePostTypeArchive()

	engine.CopyStaticFiles()

	fmt.Printf("Generated: %d posts\n", len(engine.Posts))
	fmt.Printf("Generated: %d tags\n", len(engine.Tags))
}
