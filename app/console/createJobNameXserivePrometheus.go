package console

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

// CreateJobNameXserivePrometheus: Genera un archivo de configuración de Prometheus.
// La configuración se divide en 4 secciones:
// 1. Configuración Global
// 2. Agregar jobs estáticos manualmente
// 3. Generar jobs dinámicos mediante un ciclo
// 4. Convertir a YAML
// 5. Guardar en el archivo
func CreateJobNameXserivePrometheus() {
	pathConfigPrometheus := "sdkopentelemetry/config_prometheus/prometheus.yml"

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
	puertos := []int{13371, 13372, 13373, 13374, 13375, 13376, 13377, 13378, 13379, 13310, 13311, 13312, 13313, 13314, 13315, 13316, 13317, 13318, 13319,
		13320, 13321, 13322, 13323, 13324, 13325, 13326, 13327, 13328, 13329, 13330, 13331, 13332, 13333, 13334, 13335, 13336, 13337, 13338, 13339,
		13340, 13341, 13342, 13343, 13344, 13345, 13346, 13347, 13348, 13349, 13350, 13351, 13352, 13353, 13354, 13355, 13356, 13357, 13358, 13359,
		13360, 13361, 13362, 13363, 13364, 13365, 13366, 13367, 13368, 13369, 13370}

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
