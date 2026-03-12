package unbloat

import (
	"os"
	"os/exec"
)


func installPackage(packageName string) error {
	cmd := exec.Command("npm", "install", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}