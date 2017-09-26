package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/hawkingrei/veda/internal/version"
	"github.com/hawkingrei/veda/vedad"
	"github.com/judwhite/go-svc/svc"
)

type program struct {
	vedad *vedad.VEDAD
}

func vedadFlagSet(opts *vedad.Options) *flag.FlagSet {
	flagSet := flag.NewFlagSet("vedad", flag.ExitOnError)
	flagSet.Bool("version", false, "print version string")
	flagSet.String("config", "", "path to config file")
	return flagSet
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func loadmeta(configFile string) (meta vedad.Meta, err error) {
	if configFile != "" {
		_, err = toml.DecodeFile(configFile, &meta)
		if err != nil {
			return
		}
	}
	return
}

func (p *program) Start() error {
	opts := vedad.NewOptions()
	flagSet := vedadFlagSet(opts)
	flagSet.Parse(os.Args[1:])

	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Println(version.String())
		os.Exit(0)
	}

	vedad, err := vedad.New(opts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	configFile := flagSet.Lookup("config").Value.String()
	meta, err := loadmeta(configFile)
	if err != nil {
		log.Fatalf("ERROR: failed to load config file %s - %s", configFile, err.Error())
	}
	//for k,v := range meta.Topics {
	//	fmt.Println("--")
	//	fmt.Println(k)
	//	for kk,vv := range v {
	//		fmt.Println("-")
	//		fmt.Println(kk)
	//		fmt.Println(vv.Address)
	//		fmt.Println(vv.Interval)
	//	}
	//}

	err = vedad.Loadmeta(meta)
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	vedad.Main()
	p.vedad = vedad

	return nil
}

func (p *program) Stop() error {
	if p.vedad != nil {
		p.vedad.Exit()
		log.Fatalf("INFO: stop the veda")
	}
	return nil
}
