package servers

import (
	"fmt"
	"github/Claudye/autoserver/distro"
	"github/Claudye/autoserver/utils"
	"strings"
)

type serverConfig struct {
	WebServer  string // "nginx" or "apache"
	Database   string // "mysql" or "postgres"
	PHPVersion string
}

func updateSytem() {
	upgraded := utils.AskYesNo("Do you want to upgrade the system ?")

	if upgraded {
		distro.Upgrade(true)
		distro.Update(true)
	}

	if !upgraded && utils.AskYesNo("Do you want to update the system ?") {
		distro.Update(true)
	}
}

// Check if the provided webserver is supported, and is not installed yet
// Process to installation
func installWebServer() {
	// Supported web servers and their aliases
	supportedWebServers := map[string][]string{
		"nginx":  {"nginx"},
		"apache": {"apache2", "apache"},
	}

	webServer := utils.AskForString("Which web server do you want to use? (nginx/apache)")
	webServer = strings.ToLower(strings.TrimSpace(webServer))

	aliases, ok := supportedWebServers[webServer]
	if !ok {
		fmt.Println("Unsupported web server.")
		return
	}

	for name, alts := range supportedWebServers {
		if name != webServer {
			for _, alt := range alts {
				if utils.IsInstalled(alt) {
					fmt.Printf("⚠️ Warning: Another web server (%s) seems to be installed.\n", alt)
					fmt.Println("Having multiple web servers running may cause port conflicts (e.g., port 80).")
					fmt.Println("If needed, uninstall the other one manually before using this one.")
					break
				}
			}
		}
	}

	// Check if selected server is already installed
	for _, alias := range aliases {
		if utils.IsInstalled(alias) {
			fmt.Printf("%s is already installed.\n", alias)
			return
		}
	}

	fmt.Printf("Installing %s...\n", webServer)

	pkgManager := distro.GetPackageManager()
	if pkgManager == "" {
		fmt.Println("No supported package manager found for this distro.")
		return
	}

	utils.Run(pkgManager, "install", aliases[0], "-y")

	fmt.Printf("%s has been installed.\n", webServer)
}

func installDatabase() {
	//User shouldb able to choose his database,
	// Depends the system, he will create the database, the user, give the password,etc

}

func installPHP() {

}
func afterInstallingWebServer() {

}

func afterInstallingDatabase() {

}

func afterInstallingPHP() {

}

func Start() {
	updateSytem()

	installWebServer()
	afterInstallingWebServer()

	installDatabase()
	afterInstallingDatabase()

	installPHP()
	afterInstallingPHP()

}
