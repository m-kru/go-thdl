package vet

import (
	"github.com/m-kru/go-thdl/internal/args"
	"github.com/m-kru/go-thdl/internal/utils"
	"github.com/m-kru/go-thdl/internal/vet/vhdl"
	"strings"
	"sync"
)

var vetArgs args.VetArgs

func Vet(args args.VetArgs) {
	vetArgs = args

	var wg sync.WaitGroup
	defer wg.Wait()

	vhdlFiles := []string{}

	if vetArgs.Filepath != "" {
		if strings.HasSuffix(vetArgs.Filepath, ".vhd") || strings.HasSuffix(vetArgs.Filepath, ".vhdl") {
			vhdlFiles = append(vhdlFiles, vetArgs.Filepath)
		}
	} else {
		vhdlFiles = utils.GetVHDLFilePaths()
	}

	wg.Add(1)
	vhdl.Vet(vhdlFiles, &wg)
}
