package app

import (
	"fmt"

	"github.com/alibaba/ioc-golang"
	"github.com/alibaba/ioc-golang/autowire/singleton"
	conf "github.com/alibaba/ioc-golang/config"
	"github.com/alibaba/ioc-golang/extension/config"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type App struct {
	Name *config.ConfigString  `config:",autowire.singleton.<github.com/alibaba/helloiocgo/cmd/app.App>.param.name"`
}

func (a *App) Run() {
	fmt.Println("Hello world")
	fmt.Printf("app.name: %s \n", a.Name.Value())
}

func Run() {
	if err := loadIoC(); err != nil {
		panic(err)
	}

	// We can get objects by ths id '$(PkgName).$(StructName)'
	appInterface, err := singleton.GetImpl("github.com/alibaba/helloiocgo/cmd/app.App", nil)
	if err != nil {
		panic(err)
	}
	app := appInterface.(*App)
	app.Run()
}

func loadIoC() error {
	nameOpt := conf.WithConfigName("ioc_golang")
	typeOpt := conf.WithConfigType("yaml")
	err := ioc.Load(nameOpt, typeOpt)

	return err
}
