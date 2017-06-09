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

type config struct {
	Redis    map[string]server
	Memcache map[string]server
}

type server struct {
	Address string
}

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

func (p *program) Start() error {
	opts := vedad.NewOptions()
	flagSet := vedadFlagSet(opts)
	flagSet.Parse(os.Args[1:])

	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Println(version.String())
		os.Exit(0)
	}

	var cfg vedad.Config
	configFile := flagSet.Lookup("config").Value.String()
	if configFile != "" {
		_, err := toml.DecodeFile(configFile, &cfg)
		if err != nil {
			log.Fatalf("ERROR: failed to load config file %s - %s", configFile, err.Error())
		}
	}
	for serverName, server := range cfg.Redis {
		fmt.Printf("Server: %s (%s)\n", serverName, server.Address)
	}
	os.Exit(0)
	return nil
}

func (p *program) Stop() error {
	if p.vedad != nil {
		p.vedad.Exit()
	}
	return nil
}
