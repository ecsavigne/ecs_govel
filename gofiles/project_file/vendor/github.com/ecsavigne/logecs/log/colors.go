package log

import "fmt"

/*
Los códigos de colores de la consola (ANSI escape codes) son:

1. La Estructura Real
El formato es \033[Xm, donde X es el código del atributo.

Rango de colores básicos (Texto): del 30 al 37.

Rango de colores brillantes (Texto): del 90 al 97.

El código 39: No es un color nuevo, es el código para resetear solo el color de primer plano al color por defecto de la terminal.

2. Tabla de Colores Rápidos
Para que los uses en tu proyecto:

Color	Código Estándar	Código Brillante
Negro	\033[30m	\033[90m
Rojo	\033[31m	\033[91m
Verde	\033[32m	\033[92m
Amarillo	\033[33m	\033[93m
Azul	\033[34m	\033[94m
Magenta	\033[35m	\033[95m
Cian	\033[36m	\033[96m
Blanco	\033[37m	\033[97m
RESET	\033[0m	(Resetea todo: color, negrita, etc.)
*/
type ColorText string
type ColorBackground string
type ColorLight string

var (
	Black  ColorText = "\033[30m"
	Red    ColorText = "\033[31m"
	Green  ColorText = "\033[32m"
	Yellow ColorText = "\033[33m"
	Blue   ColorText = "\033[34m"
	Purple ColorText = "\033[35m"
	Cyan   ColorText = "\033[36m"
	White  ColorText = "\033[37m"
	Reset  ColorText = "\033[0m"
)

var (
	BlackLight  ColorLight = "\033[90m"
	RedLight    ColorLight = "\033[91m"
	GreenLight  ColorLight = "\033[92m"
	YellowLight ColorLight = "\033[93m"
	BlueLight   ColorLight = "\033[94m"
	PurpleLight ColorLight = "\033[95m"
	CyanLight   ColorLight = "\033[96m"
	WhiteLight  ColorLight = "\033[97m"
)

var (
	BlackBackground  ColorBackground = "\033[40m"
	RedBackground    ColorBackground = "\033[41m"
	GreenBackground  ColorBackground = "\033[42m"
	YellowBackground ColorBackground = "\033[43m"
	BlueBackground   ColorBackground = "\033[44m"
	PurpleBackground ColorBackground = "\033[45m"
	CyanBackground   ColorBackground = "\033[46m"
	WhiteBackground  ColorBackground = "\033[47m"
)

type Color interface {
	ColorText | ColorBackground | ColorLight
}

func PrintColor[T Color](color T, text string) string {
	return fmt.Sprintf("%s%s%s", string(color), text, string(Reset))
}
