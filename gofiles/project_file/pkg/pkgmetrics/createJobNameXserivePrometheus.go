package pkgmetrics

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/spf13/viper"
)

// CreateJobNameXserivePrometheus: Genera un archivo de configuración de Prometheus.
// La configuración se divide en 4 secciones:
// 1. Configuración Global
// 2. Agregar jobs estáticos manualmente
// 3. Generar jobs dinámicos mediante un ciclo
// 4. Convertir a YAML
// 5. Guardar en el archivo
func CreateJobNameXserivePrometheus() {
	pathConfigPrometheus := viper.GetString("PROMETHEUS_CONFIG_PATH")

	// Estructura raíz del archivo Prometheus
	type GlobalConfig struct {
		ScrapeInterval string `yaml:"scrape_interval"`
	}

	type StaticConfig struct {
		// Targets []string `yaml:"targets,flow"`
		Targets []string `yaml:"targets"`
	}

	type JobConfig struct {
		JobName       string         `yaml:"job_name"`
		StaticConfigs []StaticConfig `yaml:"static_configs"`
	}

	type PrometheusConfig struct {
		Global        GlobalConfig `yaml:"global"`
		ScrapeConfigs []JobConfig  `yaml:"scrape_configs"`
	}

	// 1. Configuración Global
	config := PrometheusConfig{
		Global: GlobalConfig{
			ScrapeInterval: "10s",
		},
	}

	// 2. Agregar jobs estáticos manualmente
	config.ScrapeConfigs = append(config.ScrapeConfigs, JobConfig{
		JobName: "app-crud-mongo",
		StaticConfigs: []StaticConfig{
			{
				Targets: []string{"my.host.local:1111"},
			},
		},
	})

	// 3. Generar jobs dinámicos mediante un ciclo
	puertos := []int{13371}

	config.ScrapeConfigs = append(config.ScrapeConfigs, JobConfig{
		JobName:       "Services",
		StaticConfigs: make([]StaticConfig, 0),
	})

	services := make([]string, 0)
	for i, p := range puertos {
		serv := fmt.Sprintf("%s", fmt.Sprintf("servicex%d.socialhub.pro:%d", i+1, p))
		services = append(services, serv)
	}

	config.ScrapeConfigs[1].StaticConfigs = make([]StaticConfig, 0)
	config.ScrapeConfigs[1].StaticConfigs = append(config.ScrapeConfigs[1].StaticConfigs, StaticConfig{
		Targets: services,
	})

	// 4. Convertir a YAML
	data, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// 5. Guardar en el archivo
	err = os.WriteFile(pathConfigPrometheus, data, 0644)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
