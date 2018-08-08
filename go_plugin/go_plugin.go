package go_plugin

import (
	"io/ioutil"
	"path"
	"plugin"
	"log"
)

type GoPlugin struct {
	PluginPath          string
	CustomRunPluginName []string
	RunPlugins          []string
}

func NewGoPlugin(pluginPath string) *GoPlugin {
	return &GoPlugin{
		PluginPath: pluginPath,
	}
}

func (g *GoPlugin) Load() error {
	// 获取Plugin目录下所有插件
	osFiles, err := ioutil.ReadDir(g.PluginPath)
	if err != nil {
		return err
	}
	pluginSoMap := map[string]int{}
	for _, f := range osFiles {
		if path.Ext(f.Name()) == ".so" {
			pluginSoMap[f.Name()] = 1
			g.RunPlugins = append(g.RunPlugins, f.Name())
		}
	}
	// 检查是否有自动运行插件，如果有，则按自定义插件
	if len(g.CustomRunPluginName) > 0 {
		// 重置已经允许执行的插件
		g.RunPlugins = []string{}
		for _, pluginName := range g.CustomRunPluginName {
			if _, ok := pluginSoMap[pluginName]; ok {
				g.RunPlugins = append(g.RunPlugins, pluginName)
			}
		}
	}
	return nil
}

func (g *GoPlugin) Run() {
	for _, pluginName := range g.RunPlugins {
		p, err := plugin.Open(g.PluginPath + pluginName)
		if err != nil {
			log.Println(err)
			continue
		}
		funcSymbol, err := p.Lookup("Run")
		if err != nil {
			log.Println(err)
			continue
		}
		funcSymbol.(func([]interface{}))([]interface{}{"hello", "world"})
	}
}
