package utilidades

import "strings"

func ExtensionPermitida(extensionArchivo string) bool {
	extensionesValidas := []string{".png", ".jpg", ".jpeg"}
	for _, v := range extensionesValidas {
		if strings.ToLower(extensionArchivo) == v {
			return true
		}
	}
	return false
}
