package distro

import (
	"bufio"
	"github/Claudye/autoserver/utils"
	"os"
	"runtime"
	"strings"
)

func GetOs() string {
	return runtime.GOOS
}

func GetPackageManager() string {
	switch GetDistroName() {
	case "ubuntu", "debian":
		return "apt"
	case "centos", "rhel":
		return "yum"
	default:
		return ""
	}
}

func GetSudo() string {
	if GetOs() == "linux" {
		return "sudo"
	}
	return ""
}

func GetDistroName() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			return strings.Trim(strings.SplitN(line, "=", 2)[1], `"`)
		}
	}
	return "unknown"
}

func Update(yes bool) {
	pkg := GetPackageManager()

	if pkg == "" {
		utils.Run("echo", "Distribution non supportée pour update")
		return
	}

	args := []string{pkg, "update"}
	if yes {
		args = append(args, "-y")
	}

	utils.Run(args[0], args[1:]...)
}

func Upgrade(yes bool) {
	sudo := GetSudo()
	pkg := GetPackageManager()

	if pkg == "" {
		utils.Run("echo", "Distribution non supportée pour upgrade")
		return
	}

	args := []string{pkg, "upgrade"}
	if yes {
		args = append(args, "-y")
	}
	if sudo != "" {
		args = append([]string{sudo}, args...)
	}

	utils.Run(args[0], args[1:]...)
}
