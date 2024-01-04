//Author:github.com/Azumi67
//This script is for educational use and for my own learning, but I'd be happy if you find it useful too.
//This script simplifies the configuration of FEC & UDP2raw. You can also use IP6IP6 & ICMP next to the tunnel.
//You can send me feedback so I can use it to learn more.
//This script comes without any warranty
//Thank you.
package main

import (
    "time"
	"strconv"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/fatih/color"
	"log"
	"github.com/AlecAivazis/survey/v2"
	"path/filepath"
	"runtime"
	"syscall"
	"net"
	"io/ioutil"
)
func getIPv4() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, iface := range interfaces {
		name := iface.Name
		if strings.HasPrefix(name, "eth") || strings.HasPrefix(name, "en") {
			addresses, err := iface.Addrs()
			if err != nil {
				continue
			}

			for _, addr := range addresses {
				ipnet, ok := addr.(*net.IPNet)
				if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}

	return ""
}
func displayProgress(total, current int) {
	width := 40
	percentage := current * 100 / total
	completed := width * current / total
	remaining := width - completed

	fmt.Printf("\r[%s>%s] %d%%", strings.Repeat("=", completed), strings.Repeat(" ", remaining), percentage)
}

func displayError(message string) {
	fmt.Printf("\u2718 Error: %s\n", message)
}

func displayNotification(message string) {
	fmt.Printf("\033[93m%s\033[0m\n", message)
}

func displayCheckmark(message string) {
	fmt.Printf("\033[92m\u2714 \033[0m%s\n", message)
}

func displayLoading() {
    frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
    delay := 100 * time.Millisecond
    duration := 5 * time.Second

    endTime := time.Now().Add(duration)

    for time.Now().Before(endTime) {
        for _, frame := range frames {
            fmt.Printf("\r[%s] Loading... ", frame)
            time.Sleep(delay)
        }
    }
    fmt.Println()
}
func displayLogo2() error {
	cmd := exec.Command("bash", "-c", "/etc/./logo.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
func displayLogo() {
	logo := `
   ______    _______    __      _______          __       _____  ___  
  /    " \  |   __ "\  |" \    /"      \        /""\      (\"   \|" \ 
 // ____  \ (. |__) :) ||  |  |:        |      /    \     |.\\   \   |
/  /    ) :)|:  ____/  |:  |  |_____/   )     /' /\  \    |: \.   \\ |
(: (____/ // (|  /     |.  |   //       /    //  __'  \   |.  \    \ |
\        // |__/ \     /\  |\  |:  __   \   /   /  \\  \  |    \    \|
 \"_____ / (_______)  (__\_|_) |__|  \___) (___/    \___) \___|\____\)
`
	
    cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
    blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	green := color.New(color.FgHiGreen, color.Bold).SprintFunc()      
    yellow := color.New(color.FgHiYellow, color.Bold).SprintFunc()   
    red := color.New(color.FgHiRed, color.Bold).SprintFunc()        


	

	    logo = cyan("  ______   ") + blue(" _______  ") + green("  __    ") + yellow("   ______   ") + red("     __      ") + cyan("  _____  ___  \n") +
		cyan(" /     \" \\  ") + blue("|   __ \" ") + green(" |\" \\  ") + yellow("   /\"     \\   ") + red("   /\"\"\\     ") + cyan(" (\\\"   \\|\"  \\ \n") +
		cyan("//  ____  \\ ")  + blue("(. |__) :)") + green("||  |  ") + yellow(" |:       |  ") + red("  /    \\   ") + cyan("  |.\\\\   \\   |\n") +
		cyan("/  /    ) :)") + blue("|:  ____/ ") + green("|:  |  ") + yellow(" |_____/  )  ") + red(" /' /\\  \\   ") + cyan(" |: \\.   \\\\ |\n") +
		cyan("(: (____/ / ") + blue("(|  /     ") + green("|.  | ") + yellow("  //      /  ") + red("//   __'  \\  ") + cyan(" |.  \\    \\ |\n") +
		cyan("\\        / ") + blue("/|__/ \\   ") + green(" /\\  |\\ ") + yellow(" |:  __  \\ ") + red(" /   /  \\\\   ") + cyan ("  |    \\    \\|\n") +
		cyan(" \"_____ / ") + blue("(_______)") + green("  (__\\_|_)") + yellow(" (__) \\___)") + red("(___/    \\___)") + cyan(" \\___|\\____\\)\n")


	fmt.Println(logo)
}

func main() {
	if os.Geteuid() != 0 {
		fmt.Println("\033[91mThis script must be run as root. Please use sudo -i.\033[0m")
		os.Exit(1)
	}

	mainMenu()
}
func runCmd(cmd string) error {
	command := exec.Command("sh", "-c", cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
func azumicmd(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Azumi has failed to run the command: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}
func userInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func writeanFile(filePath string, lines []string) error {
	content := []byte(strings.Join(lines, "\n"))
	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
func readInput() {
	fmt.Print("Press Enter to continue..")
	fmt.Scanln()
	mainMenu()
}
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func mainMenu() {
	for {
		err := displayLogo2()
		if err != nil {
			log.Fatalf("Failed to display logo: %v", err)
		}
		displayLogo()
		border := "\033[93m+" + strings.Repeat("=", 70) + "+\033[0m"
		content := "\033[93m║            ▌║█║▌│║▌│║▌║▌█║ \033[92mMain Menu\033[93m  ▌│║▌║▌│║║▌█║▌                  ║"
		footer := " \033[92m            Join Opiran Telegram \033[34m@https://t.me/OPIranClub\033[0m "

		borderLength := len(border) - 2
		centeredContent := fmt.Sprintf("%[1]*s", -borderLength, content)

		fmt.Println(border)
		fmt.Println(centeredContent)
		fmt.Println(border)

		fmt.Println(border)
		fmt.Println(footer)
		fmt.Println(border)
		prompt := &survey.Select{
			Message: "Enter your choice Please:",
			Options: []string{"0. \033[91mSTATUS Menu\033[0m", "1. \033[96mUDP2RAW \033[92mNO FEC \033[0m", "2. \033[93mUDP2RAW FEC \033[92mIPV4\033[0m", "3. \033[96mUDP2RAW FEC \033[92mIPV6\033[0m", "4. \033[93mUDP2RAW FEC \033[92mICMP\033[0m", "5. \033[96mUDP2RAW FEC \033[92mIP6IP6\033[0m", "6. \033[92mStop | Restart Service\033[0m", "7. \033[91mUninstall\033[0m", "q. Exit"},
		
		}
		fmt.Println("\033[93m╰─────────────────────────────────────────────────────────────────────╯\033[0m")

		var choice string
		err = survey.AskOne(prompt, &choice)
		if err != nil {
			log.Fatalf("\033[91muser input is wrong:\033[0m %v", err)
		}
		switch choice {
		case "0. \033[91mSTATUS Menu\033[0m":
			status()
		case "1. \033[96mUDP2RAW \033[92mNO FEC \033[0m":
			udp2raw()
		case "2. \033[93mUDP2RAW FEC \033[92mIPV4\033[0m":
			udp2rawsingle()
		case "3. \033[96mUDP2RAW FEC \033[92mIPV6\033[0m":
			udp2raw6()
		case "4. \033[93mUDP2RAW FEC \033[92mICMP\033[0m":
			udp2rawIcmp()
		case "5. \033[96mUDP2RAW FEC \033[92mIP6IP6\033[0m":
			udp2rawPri()
		case "6. \033[92mStop | Restart Service\033[0m":
			startMain()
		case "7. \033[91mUninstall\033[0m":
			UniMenu()
		case "q. Exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}

		
		readInput()
	}
}
func deleteCron() {
	entriesToDelete := []string{
		"0 */1 * * * /etc/udp.sh",
		"0 */2 * * * /etc/udp.sh",
		"0 */3 * * * /etc/udp.sh",
		"0 */4 * * * /etc/udp.sh",
		"0 */5 * * * /etc/udp.sh",
		"0 */6 * * * /etc/udp.sh",
		"0 */7 * * * /etc/udp.sh",
		"0 */8 * * * /etc/udp.sh",
		"0 */9 * * * /etc/udp.sh",
		"0 */10 * * * /etc/udp.sh",
		"0 */11 * * * /etc/udp.sh",
		"0 */12 * * * /etc/udp.sh",
		"0 */13 * * * /etc/udp.sh",
		"0 */14 * * * /etc/udp.sh",
		"0 */15 * * * /etc/udp.sh",
		"0 */16 * * * /etc/udp.sh",
		"0 */17 * * * /etc/udp.sh",
		"0 */18 * * * /etc/udp.sh",
		"0 */19 * * * /etc/udp.sh",
		"0 */20 * * * /etc/udp.sh",
		"0 */21 * * * /etc/udp.sh",
		"0 */22 * * * /etc/udp.sh",
		"0 */23 * * * /etc/udp.sh",
		"0 */24 * * * /etc/udp.sh",
	}

	existingCrontab, err := exec.Command("crontab", "-l").Output()
	if err != nil {
		displayError("\033[91mNo existing cron found!\033[0m")
		return
	}

	newCrontab := string(existingCrontab)
	for _, entry := range entriesToDelete {
		if strings.Contains(newCrontab, entry) {
			newCrontab = strings.Replace(newCrontab, entry, "", -1)
		}
	}

	if newCrontab != string(existingCrontab) {
		cmd := exec.Command("crontab")
		cmd.Stdin = strings.NewReader(newCrontab)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		displayNotification("\033[92mDeleting Previous Crons..\033[0m")
	} else {
		displayError("\033[91mNothing Found, moving on..!\033[0m")
	}
}

const (
	crontabFilePath = "/var/spool/cron/crontabs/root"
	azumiudpKharej  = "azumiudp_kharej"
	azumifecKharej  = "azumifec_kharej"
	azumiudpIran    = "azumiudp_iran"
	azumifecIran    = "azumifec_iran"
)

func resKharej() {
	deleteCron()
	if _, err := os.Stat("/etc/udp.sh"); err == nil {
		os.Remove("/etc/udp.sh")
	}

	file, err := os.Create("/etc/udp.sh")
	if err != nil {
		log.Fatalf("\033[91mbash creation error:\033[0m %v", err)
	}
	defer file.Close()

	file.WriteString("#!/bin/bash\n")
	file.WriteString("sudo systemctl daemon-reload\n")
	file.WriteString("sudo sync; echo 1 > /proc/sys/vm/drop_caches\n")
	file.WriteString("sudo journalctl --vacuum-size=1M\n")

	cmd := exec.Command("chmod", "+x", "/etc/udp.sh")
	if err := cmd.Run(); err != nil {
		log.Fatalf("\033[91mchmod cmd error:\033[0m %v", err)
	}

	fmt.Println("╭──────────────────────────────────────╮")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter number of \033[92mConfigs\033[96m [2 Hours Reset Timer]\033[93m: \033[0m")
	configCountStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("\033[91minvalid input: %v\033[0m", err)
	}
	configCountStr = strings.TrimSpace(configCountStr)
	fmt.Println("╰──────────────────────────────────────╯")

	configCount, err := strconv.Atoi(configCountStr)
	if err != nil {
		log.Fatalf("\033[91minvalid input for Configs number:\033[0m %v", err)
	}

	hours := configCount * 2

	cronEntry := fmt.Sprintf("0 */%d * * * /etc/udp.sh", hours)

	crontabFile, err := os.OpenFile(crontabFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("\033[91mCouldn't open Cron:\033[0m %v", err)
	}
	defer crontabFile.Close()

	var crontabContent strings.Builder
	scanner := bufio.NewScanner(crontabFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == cronEntry {
			fmt.Println("\033[92mOh .. Cron entry already exists!\033[0m")
			return
		}
		crontabContent.WriteString(line)
		crontabContent.WriteString("\n")
	}

	for i := 1; i <= configCount; i++ {
		configAzumiudp := fmt.Sprintf("%s%d", azumiudpKharej, i)
		configAzumifec := fmt.Sprintf("%s%d", azumifecKharej, i)
		file.WriteString(fmt.Sprintf("sudo systemctl restart %s\n", configAzumiudp))
		file.WriteString(fmt.Sprintf("sudo systemctl restart %s\n", configAzumifec))
	}

	crontabContent.WriteString(cronEntry)
	crontabContent.WriteString("\n")

	if err := scanner.Err(); err != nil {
		log.Fatalf("\033[91mcrontab Reading error:\033[0m %v", err)
	}

	if err := crontabFile.Truncate(0); err != nil {
		log.Fatalf("\033[91mcouldn't truncate cron file:\033[0m %v", err)
	}

	if _, err := crontabFile.Seek(0, 0); err != nil {
		log.Fatalf("\033[91mcouldn't find cron file: \033[0m%v", err)
	}

	if _, err := crontabFile.WriteString(crontabContent.String()); err != nil {
		log.Fatalf("\033[91mCouldn't write cron file:\033[0m %v", err)
	}

	fmt.Println("\033[92mCron entry added successfully!\033[0m")
}
func resIran() {
	deleteCron()
	if _, err := os.Stat("/etc/udp.sh"); err == nil {
		os.Remove("/etc/udp.sh")
	}

	file, err := os.Create("/etc/udp.sh")
	if err != nil {
		log.Fatalf("\033[91mbash creation error:\033[0m %v", err)
	}
	defer file.Close()

	file.WriteString("#!/bin/bash\n")
	file.WriteString("sudo systemctl daemon-reload\n")
	file.WriteString("sudo sync; echo 1 > /proc/sys/vm/drop_caches\n")
	file.WriteString("sudo journalctl --vacuum-size=1M\n")

	cmd := exec.Command("chmod", "+x", "/etc/udp.sh")
	if err := cmd.Run(); err != nil {
		log.Fatalf("\033[91mchmod cmd error:\033[0m %v", err)
	}

	fmt.Println("╭──────────────────────────────────────╮")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter number of \033[92mConfigs\033[96m [2 Hours Reset Timer]\033[93m: \033[0m")
	configCountStr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("\033[91minvalid input: %v\033[0m", err)
	}
	configCountStr = strings.TrimSpace(configCountStr)
	fmt.Println("╰──────────────────────────────────────╯")

	configCount, err := strconv.Atoi(configCountStr)
	if err != nil {
		log.Fatalf("\033[91mInvalid input for Configs number:\033[0m %v", err)
	}

	hours := configCount * 2

	cronEntry := fmt.Sprintf("0 */%d * * * /etc/udp.sh", hours)

	crontabFile, err := os.OpenFile(crontabFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("\033[91mCouldn't open Cron:\033[0m %v", err)
	}
	defer crontabFile.Close()

	var crontabContent strings.Builder
	scanner := bufio.NewScanner(crontabFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == cronEntry {
			fmt.Println("\033[92mOh .. Cron entry already exists!\033[0m")
			return
		}
		crontabContent.WriteString(line)
		crontabContent.WriteString("\n")
	}

	for i := 1; i <= configCount; i++ {
		configAzumiudp := fmt.Sprintf("%s-%d", azumiudpIran, i)
		configAzumifec := fmt.Sprintf("%s-%d", azumifecIran, i)
		file.WriteString(fmt.Sprintf("sudo systemctl restart %s\n", configAzumiudp))
		file.WriteString(fmt.Sprintf("sudo systemctl restart %s\n", configAzumifec))
	}

	crontabContent.WriteString(cronEntry)
	crontabContent.WriteString("\n")

	if err := scanner.Err(); err != nil {
		log.Fatalf("\033[91mcrontab Reading error:\033[0m %v", err)
	}


	if err := crontabFile.Truncate(0); err != nil {
		log.Fatalf("\033[91mcouldn't truncate cron file:\033[0m %v", err)
	}

	if _, err := crontabFile.Seek(0, 0); err != nil {
		log.Fatalf("\033[91mcouldn't find cron file: \033[0m%v", err)
	}

	if _, err := crontabFile.WriteString(crontabContent.String()); err != nil {
		log.Fatalf("\033[91mCouldn't write cron file:\033[0m %v", err)
	}

	fmt.Println("\033[92mCron entry added successfully!\033[0m")
}
func udp2raw() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m UDP2RAW \033[92m NO FEC ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mIPV4\033[0m", "2. \033[93mIPV6\033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mIPV4\033[0m":
		udp2rawip4()
	case "2. \033[93mIPV6\033[0m":
		udp2rawip6()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func udp2rawip4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m UDP2RAW \033[92m NO FEC \033[96m IPV4 ")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[93mKharej\033[0m", "2. \033[92mIRAN\033[0m", "0. \033[94mBack to the previous menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[93mKharej\033[0m":
		kharejip4()
	case "2. \033[92mIRAN\033[0m":
		iranip4()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		udp2raw()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func kharejip4() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Kharej \033[96mIPV4\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}

	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring kharej")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
	numConfigsInput, _ := reader.ReadString('\n')
	numConfigsInput = strings.TrimSpace(numConfigsInput)
	numConfigs, err := strconv.Atoi(numConfigsInput)
	if err != nil {
		fmt.Println("\033[91mInvalid input\033[0m")
		return
	}

	for i := 0; i < numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
		portTunnel, _ := reader.ReadString('\n')
		portTunnel = strings.TrimSpace(strings.TrimSuffix(portTunnel, "\n"))


		fmt.Print("\033[93mEnter Password: \033[0m")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(strings.TrimSuffix(password, "\n"))

		fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
		wireguardPort, _ := reader.ReadString('\n')
		wireguardPort = strings.TrimSpace(strings.TrimSuffix(wireguardPort, "\n"))
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m")      
		rawMode := userInput("Enter your choice: ")      
		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		default:
			fmt.Println("\033[91mInvalid choice.\033[0m")
			return
		}

		
		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_kharej%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -s -l0.0.0.0:%s -r127.0.0.1:%s -k %s %s -a",
				portTunnel, wireguardPort, password, rawAzumi),
		}
		err := createService(azumiudpData)
		if err != nil {
			log.Fatal(err)
		}

		
                resKharej()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		readInput()
	}
}
func iranip4() {
    clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m IRAN \033[96mIPV4\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	var clientIP, portTunnel, password, wireguardPort, rawMode string

	var numConfigs int
	for {
		fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
		_, err := fmt.Scanln(&numConfigs)
		if err == nil && numConfigs > 0 {
			break
		}
		fmt.Println("\033[91mInvalid input\033[0m")
	}

	for i := 0; i <= numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("\033[93mEnter \033[92mKHAREJ\033[96m IPV4\033[93m: \033[0m")
			clientIP, _ = reader.ReadString('\n')
			clientIP = strings.TrimSpace(clientIP)
			if clientIP != "" {
				break
			}
			fmt.Println("\033[91mClient IP can't be empty, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
			portTunnel, _ = reader.ReadString('\n')
			portTunnel = strings.TrimSpace(portTunnel)
			if portTunnel != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}


		for {
			fmt.Print("\033[93mEnter Password: \033[0m")
			password, _ = reader.ReadString('\n')
			password = strings.TrimSpace(password)
			if password != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
			wireguardPort, _ = reader.ReadString('\n')
			wireguardPort = strings.TrimSpace(wireguardPort)
			if wireguardPort != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
        fmt.Println("\033[93m─────────────────\033[0m")
		
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m")  

		for {
			rawMode = userInput("Enter your choice: ")
			if rawMode == "1" || rawMode == "2" || rawMode == "3" {
				break
			}
			fmt.Println("\033[91mInvalid choice. Please try again\033[0m")
		}

		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		}

		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_iran-%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -c -l0.0.0.0:%s -r %s:%s -k \"%s\" %s -a",
				wireguardPort, clientIP, portTunnel, password, rawAzumi),
		}
		if err := createService(azumiudpData); err != nil {
			fmt.Printf("\033[91mFailed creating Azumiudp %d:\033[0m %v\n", i, err)
		}

                resIran()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		currentIPv4 := getIPv4()

	
	    if currentIPv4 != "" {
		        displayAddress(currentIPv4, wireguardPort)
	}
		readInput()
	}
}
func udp2rawip6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m UDP2RAW \033[92m NO FEC \033[96mIPV6\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[93mKharej\033[0m", "2. \033[92mIRAN \033[0m", "0. \033[94mBack to the previous menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[93mKharej\033[0m":
		kharejip6()
	case "2. \033[92mIRAN \033[0m":
		iranip6()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		udp2raw()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func kharejip6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Kharej \033[96mIPV6\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
		downl("udpspeed")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring kharej")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
	numConfigsInput, _ := reader.ReadString('\n')
	numConfigsInput = strings.TrimSpace(numConfigsInput)
	numConfigs, err := strconv.Atoi(numConfigsInput)
	if err != nil {
		fmt.Println("\033[91mInvalid input\033[0m")
		return
	}

	for i := 0; i < numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
		portTunnel, _ := reader.ReadString('\n')
		portTunnel = strings.TrimSpace(strings.TrimSuffix(portTunnel, "\n"))

		fmt.Print("\033[93mEnter Password: \033[0m")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(strings.TrimSuffix(password, "\n"))

		fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
		wireguardPort, _ := reader.ReadString('\n')
		wireguardPort = strings.TrimSpace(strings.TrimSuffix(wireguardPort, "\n"))
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m")      
		rawMode := userInput("Enter your choice: ")      
		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		default:
			fmt.Println("\033[91mInvalid choice.\033[0m")
			return
		}
				azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_kharej%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -s -l[::]:%s -r127.0.0.1:%s -k %s %s -a",
				portTunnel, wireguardPort, password, rawAzumi),
		}
		err := createService(azumiudpData)
		if err != nil {
			log.Fatal(err)
		}
                resKharej()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		readInput()
	}
}
func iranip6() {
    clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m IRAN \033[96mIPV6\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
		downl("udpspeed")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	var clientIP, portTunnel, password, wireguardPort, rawMode string

	var numConfigs int
	for {
		fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
		_, err := fmt.Scanln(&numConfigs)
		if err == nil && numConfigs > 0 {
			break
		}
		fmt.Println("\033[91mInvalid input\033[0m")
	}

	for i := 0; i <= numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("\033[93mEnter \033[92mKHAREJ\033[96m IPV6\033[93m: \033[0m")
			clientIP, _ = reader.ReadString('\n')
			clientIP = strings.TrimSpace(clientIP)
			if clientIP != "" {
				break
			}
			fmt.Println("\033[91mClient IP can't be empty, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
			portTunnel, _ = reader.ReadString('\n')
			portTunnel = strings.TrimSpace(portTunnel)
			if portTunnel != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter Password: \033[0m")
			password, _ = reader.ReadString('\n')
			password = strings.TrimSpace(password)
			if password != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
			wireguardPort, _ = reader.ReadString('\n')
			wireguardPort = strings.TrimSpace(wireguardPort)
			if wireguardPort != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
        fmt.Println("\033[93m─────────────────\033[0m")
		
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m") 

		for {
			rawMode = userInput("Enter your choice: ")
			if rawMode == "1" || rawMode == "2" || rawMode == "3" {
				break
			}
			fmt.Println("\033[91mInvalid choice. Please try again\033[0m")
		}

		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		}

		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_iran-%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -c -l[::]:%s -r [%s]:%s -k \"%s\" %s -a",
				wireguardPort, clientIP, portTunnel, password, rawAzumi),
		}
		if err := createService(azumiudpData); err != nil {
			fmt.Printf("\033[91mFailed creating Azumiudp %d:\033[0m %v\n", i, err)
		}

                resIran()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		currentIPv4 := getIPv4()

	
	    if currentIPv4 != "" {
		        displayAddress(currentIPv4, wireguardPort)
	}
		readInput()
	}
}
func startMain() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Service \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mRestart\033[0m", "2. \033[93mStop \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mRestart\033[0m":
		start()
	case "2. \033[93mStop \033[0m":
		stop()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func start() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Restart \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mUDP2RAW\033[0m", "2. \033[93mFEC \033[0m", "0. \033[94mBack to the previous menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mUDP2RAW\033[0m":
		restartudp()
	case "2. \033[93mFEC \033[0m":
		restartfec()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		startMain()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func restartudp() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	displayNotification("\033[93mRestarting UDP2raw \033[93m..\033[0m")
	fmt.Println("\033[93m╭─────────────────────────────────────────────╮\033[0m")

	fmt.Printf("\033[93mHow many \033[92mConfigs\033[93m do you have? \033[0m")
	var numConfigs int
	_, err := fmt.Scanf("%d", &numConfigs)
	if err != nil {
		fmt.Printf("\033[91mnumber of configs input error: %s\033[0m\n", err)
		return
	}

	for i := 1; i <= numConfigs; i++ {
		serviceName := fmt.Sprintf("azumiudp_kharej%d", i)
		cmd := exec.Command("systemctl", "restart", serviceName)
		cmd.Run()
		time.Sleep(1 * time.Second)
	}

	for i := 1; i <= numConfigs; i++ {
		serviceName := fmt.Sprintf("azumiudp_iran-%d", i+1)
		cmd := exec.Command("systemctl", "restart", serviceName)
		cmd.Run()
		time.Sleep(1 * time.Second)
	}

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 0.1
	duration := 1.0
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
		}
	}

	displayCheckmark("\033[92mRestart completed!\033[0m")
}
func restartfec() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	displayNotification("\033[93mRestarting.. \033[93m..\033[0m")
	fmt.Println("\033[93m╭─────────────────────────────────────────────╮\033[0m")

	fmt.Printf("\033[93mHow many \033[92mConfigs\033[93m do you have? \033[0m")
	var numConfigs int
	_, err := fmt.Scanf("%d", &numConfigs)
	if err != nil {
		fmt.Printf("\033[91mnumber of configs input error: %s\033[0m\n", err)
		return
	}

	for i := 1; i <= numConfigs; i++ {
		serviceName := fmt.Sprintf("azumifec_kharej%d", i)
		cmd := exec.Command("systemctl", "restart", serviceName)
		cmd.Run()
		time.Sleep(1 * time.Second)
	}

	for i := 1; i <= numConfigs; i++ {
		serviceName := fmt.Sprintf("azumifec_iran-%d", i+1)
		cmd := exec.Command("systemctl", "restart", serviceName)
		cmd.Run()
		time.Sleep(1 * time.Second)
	}

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 0.1
	duration := 1.0
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
		}
	}

	displayCheckmark("\033[92mRestart completed!\033[0m")
}
func stop() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Stop \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mUDP2RAW\033[0m", "2. \033[93mFEC \033[0m", "0. \033[94mBack to the previous menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mUDP2RAW\033[0m":
		stopudp()
	case "2. \033[93mFEC \033[0m":
		stopfec()
	case "0. \033[94mBack to the previous menu\033[0m":
	    clearScreen()
		startMain()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func stopudp() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	displayNotification("\033[93mStopping UDP2raw \033[93m..\033[0m")
	fmt.Println("\033[93m╭─────────────────────────────────────────────╮\033[0m")

	fmt.Printf("\033[93mHow many \033[92mConfigs\033[93m do you have? \033[0m")
	var numConfigs int
	_, err := fmt.Scanf("%d", &numConfigs)
	if err != nil {
		fmt.Printf("\033[91mnumber of configs input error: %s\033[0m\n", err)
		return
	}

	for i := 1; i <= numConfigs; i++ {
		serviceName := fmt.Sprintf("azumiudp_kharej%d", i)
		cmd := exec.Command("systemctl", "stop", serviceName)
		cmd.Run()
		time.Sleep(1 * time.Second)
	}

	for i := 1; i <= numConfigs; i++ {
		serviceName := fmt.Sprintf("azumiudp_iran-%d", i+1)
		cmd := exec.Command("systemctl", "stop", serviceName)
		cmd.Run()
		time.Sleep(1 * time.Second)
	}

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 0.1
	duration := 1.0
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
		}
	}

	displayCheckmark("\033[92mService/s Stopped!\033[0m")
}
func stopfec() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	displayNotification("\033[93mStopping.. \033[93m..\033[0m")
	fmt.Println("\033[93m╭─────────────────────────────────────────────╮\033[0m")

	fmt.Printf("\033[93mHow many \033[92mConfigs\033[93m do you have? \033[0m")
	var numConfigs int
	_, err := fmt.Scanf("%d", &numConfigs)
	if err != nil {
		fmt.Printf("\033[91mnumber of configs input error: %s\033[0m\n", err)
		return
	}

	for i := 1; i <= numConfigs; i++ {
		serviceName := fmt.Sprintf("azumifec_kharej%d", i)
		cmd := exec.Command("systemctl", "stop", serviceName)
		cmd.Run()
		time.Sleep(1 * time.Second)
	}

	for i := 1; i <= numConfigs; i++ {
		serviceName := fmt.Sprintf("azumifec_iran-%d", i+1)
		cmd := exec.Command("systemctl", "stop", serviceName)
		cmd.Run()
		time.Sleep(1 * time.Second)
	}

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 0.1
	duration := 1.0
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(time.Duration(delay * float64(time.Second)))
		}
	}

	displayCheckmark("\033[92mService/s Stopped!\033[0m")
}
func status() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Status \033[93mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mUDP2RAW\033[0m", "2. \033[93mFEC \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mUDP2RAW\033[0m":
		udpStatus()
	case "2. \033[93mFEC \033[0m":
		fecStatus()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func udpStatus() {
	var numConfigs int
	fmt.Printf("\033[93mEnter the \033[92mnumber\033[93m of \033[96mConfigs\033[93m:\033[0m ")
	_, err := fmt.Scanf("%d", &numConfigs)
	if err != nil {
		fmt.Printf("Error reading the number of configs: %s\n", err)
		return
	}

	services := []string{"azumiudp_kharej", "azumiudp_iran"}

	fmt.Println("\033[93m            ╔════════════════════════════════════════════╗\033[0m")
	fmt.Println("\033[93m            ║               \033[92mUDP2RAW Status\033[93m               ║\033[0m")
	fmt.Println("\033[93m            ╠════════════════════════════════════════════╣\033[0m")

	for _, service := range services {
		for i := 1; i <= numConfigs; i++ {
			configServiceName := fmt.Sprintf("%s%d", service, i)
			if service == "azumiudp_iran" {
				configServiceName = fmt.Sprintf("%s-%d", service, i)
			}
			cmd := exec.Command("systemctl", "is-active", "--quiet", configServiceName)
			err := cmd.Run()
			if err != nil {
				continue
			}

			status := "\033[92m✓ Active      \033[0m"
			displayName := ""
			switch service {
			case "azumiudp_iran":
				displayName = "\033[93mIRAN Server   \033[0m"
			case "azumiudp_kharej":
				displayName = "\033[93mKharej Server \033[0m"
			default:
				displayName = service
			}

			fmt.Printf("           \033[93m ║\033[0m    %s %d:   |    %s\033[93m ║\033[0m\n", displayName, i, status)
		}
	}

	fmt.Println("\033[93m            ╚════════════════════════════════════════════╝\033[0m")
}
func fecStatus() {
	var numConfigs int
	fmt.Printf("\033[93mEnter the \033[92mnumber\033[93m of \033[96mConfigs\033[93m:\033[0m ")
	_, err := fmt.Scanf("%d", &numConfigs)
	if err != nil {
		fmt.Printf("Error reading the number of configs: %s\n", err)
		return
	}

	services := []string{"azumifec_kharej", "azumifec_iran"}

	fmt.Println("\033[93m            ╔════════════════════════════════════════════╗\033[0m")
	fmt.Println("\033[93m            ║                  \033[92mFEC Status\033[93m                ║\033[0m")
	fmt.Println("\033[93m            ╠════════════════════════════════════════════╣\033[0m")

	for _, service := range services {
		for i := 1; i <= numConfigs; i++ {
			configServiceName := fmt.Sprintf("%s%d", service, i)
			if service == "azumifec_iran" {
				configServiceName = fmt.Sprintf("%s-%d", service, i)
			}
			cmd := exec.Command("systemctl", "is-active", "--quiet", configServiceName)
			err := cmd.Run()
			if err != nil {
				continue
			}

			status := "\033[92m✓ Active      \033[0m"
			displayName := ""
			switch service {
			case "azumifec_iran":
				displayName = "\033[93mIRAN Server   \033[0m"
			case "azumifec_kharej":
				displayName = "\033[93mKharej Server \033[0m"
			default:
				displayName = service
			}

			fmt.Printf("           \033[93m ║\033[0m    %s %d:   |    %s\033[93m ║\033[0m\n", displayName, i, status)
		}
	}

	fmt.Println("\033[93m            ╚════════════════════════════════════════════╝\033[0m")
}
var downlURLs = map[string]map[string]string{
	"udp2raw": {
		"amd64": "https://github.com/Azumi67/udp2raw/releases/download/latest/udp2raw_amd64",
		"arm64": "https://github.com/Azumi67/udp2raw/releases/download/latest/udp2raw_arm",
	},
	"udpspeed": {
		"amd64": "https://github.com/Azumi67/udp2raw/releases/download/speed/speederv2_amd64",
		"arm64": "https://github.com/Azumi67/udp2raw/releases/download/speed/speederv2_arm",
	},
}

var extractDirs = map[string]string{
	"udp2raw":  "udp",
	"udpspeed": "udp",
}

func cpuArch() string {
	machine := runtime.GOARCH
	if machine == "amd64" {
		return "amd64"
	} else if machine == "arm64" {
		return "arm64"
	} else {
		errorMessage := fmt.Sprintf("\033[91munsupported CPU arch: %s\033[0m", machine)
		panic(errorMessage)
	}
}

func downl(tool string) {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Downloading..")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	defer Panic()

	arch := cpuArch()

	urls, ok := downlURLs[tool]
	if !ok {
		displayError("\033[91m" + fmt.Sprintf("unknown URL: %s", tool) + "\033[0m")
		return
	}

	url, ok := urls[arch]
	if !ok {
		displayError("\033[91m" + fmt.Sprintf("unsupported CPU arch for %s: %s", tool, arch) + "\033[0m")
		return
	}

	fileName := tool

	displayNotification("\033[93m" + fmt.Sprintf("Downloading %s for %s ..", tool, arch) + "\033[0m")

	err := downlFile(url, fileName)
	if err != nil {
		displayError("\033[91m" + fmt.Sprintf("%s download error: %s", tool, err.Error()) + "\033[0m")
		return
	}

	err = os.MkdirAll(extractDirs[tool], 0755)
	if err != nil {
		displayError("\033[91m" + fmt.Sprintf("failed to create dir %s: %s", extractDirs[tool], err.Error()) + "\033[0m")
		return
	}

	err = os.Rename(fileName, filepath.Join(extractDirs[tool], fileName))
	if err != nil {
		displayError("\033[91m" + fmt.Sprintf("%s renaming error: %s", tool, err.Error()) + "\033[0m")
		return
	}

	err = os.Chmod(filepath.Join(extractDirs[tool], fileName), 0755)
	if err != nil {
		displayError("\033[91m" + fmt.Sprintf("failed to set permissionZz for %s: %s", tool, err.Error()) + "\033[0m")
		return
	}

	displayCheckmark(fmt.Sprintf("\033[92m%s Downloaded!\033[0m", strings.Title(tool)))
}

func downlFile(url string, fileName string) error {
	cmd := exec.Command("wget", "--quiet", "--show-progress", "-O", fileName, url)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Panic() {
	if r := recover(); r != nil {
		displayError(fmt.Sprintf("\033[91man error occurred: %v\033[0m", r))
		syscall.Exit(1)
	}
}

func udp2rawsingle() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m UDP2RAW \033[92mFEC \033[96mIPV4\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mKharej\033[0m", "2. \033[93mIRAN \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mKharej\033[0m":
		kharejSingle()
	case "2. \033[93mIRAN \033[0m":
		iranSingle()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}

func kharejSingle() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Kharej \033[96mIPV4\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
		downl("udpspeed")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}

	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring kharej")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
	numConfigsInput, _ := reader.ReadString('\n')
	numConfigsInput = strings.TrimSpace(numConfigsInput)
	numConfigs, err := strconv.Atoi(numConfigsInput)
	if err != nil {
		fmt.Println("\033[91mInvalid input\033[0m")
		return
	}

	for i := 0; i < numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
		portTunnel, _ := reader.ReadString('\n')
		portTunnel = strings.TrimSpace(strings.TrimSuffix(portTunnel, "\n"))

		fmt.Print("\033[93mEnter \033[92mFEC \033[93mPort: \033[0m")
		portFEC, _ := reader.ReadString('\n')
		portFEC = strings.TrimSpace(strings.TrimSuffix(portFEC, "\n"))

		fmt.Print("\033[93mEnter Password: \033[0m")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(strings.TrimSuffix(password, "\n"))

		fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
		wireguardPort, _ := reader.ReadString('\n')
		wireguardPort = strings.TrimSpace(strings.TrimSuffix(wireguardPort, "\n"))
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m")      
		rawMode := userInput("Enter your choice: ")      
		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		default:
			fmt.Println("\033[91mInvalid choice.\033[0m")
			return
		}

		
		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_kharej%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -s -l0.0.0.0:%s -r127.0.0.1:%s -k %s %s -a",
				portTunnel, portFEC, password, rawAzumi),
		}
		err := createService(azumiudpData)
		if err != nil {
			log.Fatal(err)
		}

		
		azumifecData := ServiceData{
			ServiceName: fmt.Sprintf("azumifec_kharej%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udpspeed -s -l0.0.0.0:%s --mode 1 -r127.0.0.1:%s -f20:10 -k %s",
				portFEC, wireguardPort, password),
		}
		err = createService(azumifecData)
		if err != nil {
			log.Fatal(err)
		}
                resKharej()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		readInput()
	}
}
func iranSingle() {
    clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m IRAN \033[96mIPV4\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
		downl("udpspeed")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	var clientIP, portTunnel, portFEC, password, wireguardPort, rawMode string

	var numConfigs int
	for {
		fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
		_, err := fmt.Scanln(&numConfigs)
		if err == nil && numConfigs > 0 {
			break
		}
		fmt.Println("\033[91mInvalid input\033[0m")
	}

	for i := 0; i <= numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("\033[93mEnter \033[92mKHAREJ\033[96m IPV4\033[93m: \033[0m")
			clientIP, _ = reader.ReadString('\n')
			clientIP = strings.TrimSpace(clientIP)
			if clientIP != "" {
				break
			}
			fmt.Println("\033[91mClient IP can't be empty, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
			portTunnel, _ = reader.ReadString('\n')
			portTunnel = strings.TrimSpace(portTunnel)
			if portTunnel != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mFEC \033[93mPort: \033[0m")
			portFEC, _ = reader.ReadString('\n')
			portFEC = strings.TrimSpace(portFEC)
			if portFEC != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter Password: \033[0m")
			password, _ = reader.ReadString('\n')
			password = strings.TrimSpace(password)
			if password != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
			wireguardPort, _ = reader.ReadString('\n')
			wireguardPort = strings.TrimSpace(wireguardPort)
			if wireguardPort != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
        fmt.Println("\033[93m─────────────────\033[0m")
		
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m")  

		for {
			rawMode = userInput("Enter your choice: ")
			if rawMode == "1" || rawMode == "2" || rawMode == "3" {
				break
			}
			fmt.Println("\033[91mInvalid choice. Please try again\033[0m")
		}

		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		}

		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_iran-%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -c -l0.0.0.0:%s -r %s:%s -k \"%s\" %s -a",
				wireguardPort, clientIP, portTunnel, password, rawAzumi),
		}
		if err := createService(azumiudpData); err != nil {
			fmt.Printf("\033[91mFailed creating Azumiudp %d:\033[0m %v\n", i, err)
		}

		azumifecData := ServiceData{
			ServiceName: fmt.Sprintf("azumifec_iran-%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udpspeed -c -l0.0.0.0:%s -r127.0.0.1:%s --mode 1 -f20:10 -k \"%s\"",
				portFEC, wireguardPort, password),
		}
		if err := createService(azumifecData); err != nil {
			fmt.Printf("\033[91mFailed creating Azumifec %d:\033[0m %v\n", i, err)
		}
                resIran()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		currentIPv4 := getIPv4()

	
	    if currentIPv4 != "" {
		        displayAddress(currentIPv4, portFEC)
	}
		readInput()
	}
}
func displayAddress(currentIPv4 string, portFEC string) {
	fmt.Println("\033[93m╭─────────────────────────────────────────────────────────╮\033[0m")
	fmt.Printf("\033[93m| Your Address & Port: %s : %s  \033[0m\n", currentIPv4, portFEC)
	fmt.Println("\033[93m╰─────────────────────────────────────────────────────────╯\033[0m")
}
func udp2raw6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m UDP2RAW \033[92mFEC \033[96mIPV6\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mKharej\033[0m", "2. \033[93mIRAN \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mKharej\033[0m":
		kharej6()
	case "2. \033[93mIRAN \033[0m":
		iran6()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func kharej6() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Kharej \033[96mIPV6\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
		downl("udpspeed")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring kharej")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
	numConfigsInput, _ := reader.ReadString('\n')
	numConfigsInput = strings.TrimSpace(numConfigsInput)
	numConfigs, err := strconv.Atoi(numConfigsInput)
	if err != nil {
		fmt.Println("\033[91mInvalid input\033[0m")
		return
	}

	for i := 0; i < numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
		portTunnel, _ := reader.ReadString('\n')
		portTunnel = strings.TrimSpace(strings.TrimSuffix(portTunnel, "\n"))

		fmt.Print("\033[93mEnter \033[92mFEC \033[93mPort: \033[0m")
		portFEC, _ := reader.ReadString('\n')
		portFEC = strings.TrimSpace(strings.TrimSuffix(portFEC, "\n"))

		fmt.Print("\033[93mEnter Password: \033[0m")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(strings.TrimSuffix(password, "\n"))

		fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
		wireguardPort, _ := reader.ReadString('\n')
		wireguardPort = strings.TrimSpace(strings.TrimSuffix(wireguardPort, "\n"))
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m")      
		rawMode := userInput("Enter your choice: ")      
		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		default:
			fmt.Println("\033[91mInvalid choice.\033[0m")
			return
		}

		
		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_kharej%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -s -l[::]:%s -r127.0.0.1:%s -k %s %s -a",
				portTunnel, portFEC, password, rawAzumi),
		}
		err := createService(azumiudpData)
		if err != nil {
			log.Fatal(err)
		}

		
		azumifecData := ServiceData{
			ServiceName: fmt.Sprintf("azumifec_kharej%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udpspeed -s -l[::]:%s -r127.0.0.1:%s --mode 1 -f20:10 -k %s",
				portFEC, wireguardPort, password),
		}
		err = createService(azumifecData)
		if err != nil {
			log.Fatal(err)
		}
                resKharej()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		readInput()
	}
}
func iran6() {
    clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m IRAN \033[96mIPV6\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
		downl("udpspeed")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	var clientIP, portTunnel, portFEC, password, wireguardPort, rawMode string

	var numConfigs int
	for {
		fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
		_, err := fmt.Scanln(&numConfigs)
		if err == nil && numConfigs > 0 {
			break
		}
		fmt.Println("\033[91mInvalid input\033[0m")
	}

	for i := 0; i <= numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("\033[93mEnter \033[92mKHAREJ\033[96m IPV6\033[93m: \033[0m")
			clientIP, _ = reader.ReadString('\n')
			clientIP = strings.TrimSpace(clientIP)
			if clientIP != "" {
				break
			}
			fmt.Println("\033[91mClient IP can't be empty, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
			portTunnel, _ = reader.ReadString('\n')
			portTunnel = strings.TrimSpace(portTunnel)
			if portTunnel != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mFEC \033[93mPort: \033[0m")
			portFEC, _ = reader.ReadString('\n')
			portFEC = strings.TrimSpace(portFEC)
			if portFEC != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter Password: \033[0m")
			password, _ = reader.ReadString('\n')
			password = strings.TrimSpace(password)
			if password != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
			wireguardPort, _ = reader.ReadString('\n')
			wireguardPort = strings.TrimSpace(wireguardPort)
			if wireguardPort != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
        fmt.Println("\033[93m─────────────────\033[0m")
		
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m") 

		for {
			rawMode = userInput("Enter your choice: ")
			if rawMode == "1" || rawMode == "2" || rawMode == "3" {
				break
			}
			fmt.Println("\033[91mInvalid choice. Please try again\033[0m")
		}

		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		}

		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_iran-%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -c -l[::]:%s -r [%s]:%s -k \"%s\" %s -a",
				wireguardPort, clientIP, portTunnel, password, rawAzumi),
		}
		if err := createService(azumiudpData); err != nil {
			fmt.Printf("\033[91mFailed creating Azumiudp %d:\033[0m %v\n", i, err)
		}

		azumifecData := ServiceData{
			ServiceName: fmt.Sprintf("azumifec_iran-%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udpspeed -c -l[::]:%s -r127.0.0.1:%s --mode 1 -f20:10 -k \"%s\"",
				portFEC, wireguardPort, password),
		}
		if err := createService(azumifecData); err != nil {
			fmt.Printf("\033[91mFailed creating Azumifec %d:\033[0m %v\n", i, err)
		}
                resIran()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		currentIPv4 := getIPv4()

	
	    if currentIPv4 != "" {
		        displayAddress(currentIPv4, portFEC)
	}
		readInput()
	}
}
func UniMenu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m Uninstallation \033[96mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mUDP2RAW + FEC\033[0m", "2. \033[93mUDP2RAW + ICMP \033[0m", "3. \033[96mUDP2RAW + IP6IP6 \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mUDP2RAW + FEC\033[0m":
		uni1()
	case "2. \033[93mUDP2RAW + ICMP \033[0m":
		removeICMPSingle()
	case "3. \033[96mUDP2RAW + IP6IP6 \033[0m":
		removeIpSingle()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func removeIPIP6() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving \033[92mIP6IP6\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	defer Panic()

	if _, err := os.Stat("/etc/ipip.sh"); err == nil {
		runCmd("rm /etc/ipip.sh")
	}
	if _, err := os.Stat("/etc/private.sh"); err == nil {
		runCmd("rm /etc/private.sh")
	}

	displayNotification("\033[93mRemoving cronjob...\033[0m")
	runCmd("crontab -l | grep -v \"@reboot /bin/bash /etc/ipip.sh\" | crontab -")
	runCmd("crontab -l | grep -v \"@reboot /bin/bash /etc/private.sh\" | crontab -")

	runCmd("sudo rm /etc/ping_v6.sh")
	runCmd("sudo rm /etc/ping_ip.sh")

	runCmd("systemctl disable ping_v6.service > /dev/null 2>&1")
	runCmd("systemctl stop ping_v6.service > /dev/null 2>&1")
	runCmd("rm /etc/systemd/system/ping_v6.service > /dev/null 2>&1")
	runCmd("systemctl disable ping_ip.service > /dev/null 2>&1")
	runCmd("systemctl stop ping_ip.service > /dev/null 2>&1")
	runCmd("rm /etc/systemd/system/ping_ip.service > /dev/null 2>&1")
	runCmd("systemctl daemon-reload")

	runCmd("ip link set dev azumip down > /dev/null")
	runCmd("ip tunnel del azumip > /dev/null")
	time.Sleep(1 * time.Second)
	runCmd("ip link set dev azumi down > /dev/null")
	runCmd("ip tunnel del azumi > /dev/null")


	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92mPrivateIP Uninstallation completed!\033[0m")
}
func removeIpSingle() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving \033[92mUDP + IP6IP6\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	defer Panic()
        deleteCron()
	removeIPIP6()
	removeSingle()
	if _, err := os.Stat("/etc/udp.sh"); err == nil {
		os.RemoveAll("/etc/udp.sh")
	}

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92mUDP + IP6IP6 Uninstallation completed!\033[0m")
}
func removeICMPSingle() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving \033[92mUDP + ICMP\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
        
	defer Panic()
	deleteCron()
        removeICMP()
	resetICMP()
	removeSingle()
	if _, err := os.Stat("/etc/udp.sh"); err == nil {
		os.RemoveAll("/etc/udp.sh")
	}
    
	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92mUDP + ICMP Uninstallation completed!\033[0m")
}
func removeSingle() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving UDP2RAW Config ..\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
        
	defer Panic()
        deleteCron()
	if _, err := os.Stat("/root/udp"); err == nil {
		os.RemoveAll("/root/udp")
	}
	if _, err := os.Stat("/etc/udp.sh"); err == nil {
		os.RemoveAll("/etc/udp.sh")
	}

	azumiServices := []string{
		"azumiudp_iran-1", "azumifec_iran-1", "azumifec_kharej1", "azumiudp_kharej1",
		"azumiudp_iran-2", "azumifec_iran-2", "azumifec_kharej2", "azumiudp_kharej2",
		"azumiudp_iran-3", "azumifec_iran-3", "azumifec_kharej3", "azumiudp_kharej3",
		"azumiudp_iran-4", "azumifec_iran-4", "azumifec_kharej4", "azumiudp_kharej4",
		"azumiudp_iran-5", "azumifec_iran-5", "azumifec_kharej5", "azumiudp_kharej5",
	}

	for _, serviceName := range azumiServices {
		runCmd(fmt.Sprintf("systemctl disable %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("systemctl stop %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("rm /etc/systemd/system/%s.service > /dev/null 2>&1", serviceName))
		
	}

	runCmd("systemctl daemon-reload")

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92mUDP Uninstallation completed!\033[0m")
}
func uni1() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m UDP2RAW + FEC \033[92mUninstallation \033[96mMenu\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. Config1", "2. Config2", "3. Config3", "4. Config4", "5. Config5", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. Config1":
		removeSingle1()
	case "2. Config2":
		removeSingle2()
	case "3. Config3":
		removeSingle3()
	case "4. Config4":
		removeSingle4()
	case "5. Config5m":
		removeSingle5()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func removeSingle1() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving UDP2RAW Config 1 ..\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
        
	defer Panic()
        deleteCron()
	if _, err := os.Stat("/root/udp"); err == nil {
		os.RemoveAll("/root/udp")
	}
	if _, err := os.Stat("/etc/udp.sh"); err == nil {
		os.RemoveAll("/etc/udp.sh")
	}

	azumiServices := []string{
		"azumiudp_iran-1", "azumifec_iran-1", "azumifec_kharej1", "azumiudp_kharej1",
	}

	for _, serviceName := range azumiServices {
		runCmd(fmt.Sprintf("systemctl disable %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("systemctl stop %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("rm /etc/systemd/system/%s.service > /dev/null 2>&1", serviceName))
		
	}

	runCmd("systemctl daemon-reload")

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92mUDP Uninstallation completed!\033[0m")
}
func removeSingle2() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving UDP2RAW Config 2 ..\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
        
	defer Panic()
        deleteCron()
	if _, err := os.Stat("/root/udp"); err == nil {
		os.RemoveAll("/root/udp")
	}
	if _, err := os.Stat("/etc/udp.sh"); err == nil {
		os.RemoveAll("/etc/udp.sh")
	}

	azumiServices := []string{
		"azumiudp_iran-2", "azumifec_iran-2", "azumifec_kharej2", "azumiudp_kharej2",

	}

	for _, serviceName := range azumiServices {
		runCmd(fmt.Sprintf("systemctl disable %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("systemctl stop %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("rm /etc/systemd/system/%s.service > /dev/null 2>&1", serviceName))
		
	}

	runCmd("systemctl daemon-reload")

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92mUDP Uninstallation completed!\033[0m")
}
func removeSingle3() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving UDP2RAW Config 3 ..\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
        deleteCron()
	defer Panic()

	if _, err := os.Stat("/root/udp"); err == nil {
		os.RemoveAll("/root/udp")
	}
	if _, err := os.Stat("/etc/udp.sh"); err == nil {
		os.RemoveAll("/etc/udp.sh")
	}

	azumiServices := []string{
		"azumiudp_iran-3", "azumifec_iran-3", "azumifec_kharej3", "azumiudp_kharej3",
	}

	for _, serviceName := range azumiServices {
		runCmd(fmt.Sprintf("systemctl disable %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("systemctl stop %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("rm /etc/systemd/system/%s.service > /dev/null 2>&1", serviceName))
		
	}

	runCmd("systemctl daemon-reload")

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92mUDP Uninstallation completed!\033[0m")
}
func removeSingle4() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving UDP2RAW Config 4 ..\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
        deleteCron()
	defer Panic()

	if _, err := os.Stat("/root/udp"); err == nil {
		os.RemoveAll("/root/udp")
	}
	if _, err := os.Stat("/etc/udp.sh"); err == nil {
		os.RemoveAll("/etc/udp.sh")
	}

	azumiServices := []string{
		"azumiudp_iran-4", "azumifec_iran-4", "azumifec_kharej4", "azumiudp_kharej4",
	}

	for _, serviceName := range azumiServices {
		runCmd(fmt.Sprintf("systemctl disable %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("systemctl stop %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("rm /etc/systemd/system/%s.service > /dev/null 2>&1", serviceName))
		
	}

	runCmd("systemctl daemon-reload")

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92mUDP Uninstallation completed!\033[0m")
}
func removeSingle5() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mRemoving UDP2RAW Config 5 ..\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
        deleteCron()
	defer Panic()

	if _, err := os.Stat("/root/udp"); err == nil {
		os.RemoveAll("/root/udp")
	}
	if _, err := os.Stat("/etc/udp.sh"); err == nil {
		os.RemoveAll("/etc/udp.sh")
	}

	azumiServices := []string{
		"azumiudp_iran-5", "azumifec_iran-5", "azumifec_kharej5", "azumiudp_kharej5",
	}

	for _, serviceName := range azumiServices {
		runCmd(fmt.Sprintf("systemctl disable %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("systemctl stop %s.service > /dev/null 2>&1", serviceName))
		runCmd(fmt.Sprintf("rm /etc/systemd/system/%s.service > /dev/null 2>&1", serviceName))
		
	}

	runCmd("systemctl daemon-reload")

	fmt.Print("Progress: ")

	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	delay := 100 * time.Millisecond
	duration := 1 * time.Second
	endTime := time.Now().Add(duration)

	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r[%s] Loading...  ", frame)
			time.Sleep(delay)
			fmt.Printf("\r[%s]             ", frame)
			time.Sleep(delay)
		}
	}

	displayCheckmark("\033[92mUDP Uninstallation completed!\033[0m")
}

func installICMP() {
	displayNotification("\033[93mInstalling \033[92mIcmptunnel\033[93m ...\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayLoading()

	ipv4ForwardStatus, _ := exec.Command("sysctl", "-n", "net.ipv4.ip_forward").Output()
	if string(ipv4ForwardStatus) != "1\n" {
		exec.Command("sysctl", "net.ipv4.ip_forward=1").Run()
	}

	if _, err := os.Stat("/root/icmptunnel"); err == nil {
		os.RemoveAll("/root/icmptunnel")
	}

	cloneCommand := "git clone https://github.com/jamesbarlow/icmptunnel.git icmptunnel"
	cloneResult := exec.Command("bash", "-c", cloneCommand).Run()
	if cloneResult != nil {
		fmt.Println("Error: Failed to clone Repo.")
		return
	}

	if _, err := os.Stat("/root/icmptunnel"); err == nil {
		os.Chdir("/root/icmptunnel")

		exec.Command("sudo", "apt", "install", "-y", "net-tools").Run()
		exec.Command("sudo", "apt", "install", "-y", "make").Run()
		exec.Command("sudo", "apt-get", "install", "-y", "libssl-dev").Run()
		exec.Command("sudo", "apt", "install", "-y", "g++").Run()

		exec.Command("make").Run()

		os.Chdir("..")
	} else {
		displayError("\033[91micmptunnel folder not found !\033[0m")
	}
}
func removeICMP() {
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
    displayNotification("\033[93mRemoving icmptunnel...\033[0m")
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")

    defer Panic()

    cmds := []string{
        "crontab -l | grep -v \"@reboot /bin/bash /etc/icmp.sh\" | crontab -",
        "crontab -l | grep -v \"@reboot /bin/bash /etc/icmp-iran.sh\" | crontab -",
    }

    for _, cmd := range cmds {
        err := runCmd(cmd)
        if err != nil {
            displayError(fmt.Sprintf("Error: %s", err.Error()))
            return
        }
    }

    fmt.Print("Progress: ")

    lsofCmd := exec.Command("lsof", "-t", "/root/icmptunnel/icmptunnel")
    lsofOutput, err := lsofCmd.Output()

    if err == nil {
        pids := strings.Split(strings.TrimSpace(string(lsofOutput)), "\n")
        for _, pid := range pids {
            killCmd := exec.Command("kill", pid)
            killCmd.Run()
        }

        rmCmd := exec.Command("rm", "-rf", "/root/icmptunnel")
        rmCmd.Run()
    } else if _, ok := err.(*exec.ExitError); ok {
        displayError(fmt.Sprintf("\033[91mError: This '/root/icmptunnel' doesn't exist.\033[0m"))
    } else {
        displayError(fmt.Sprintf("\033[91mError: %s\033[0m", err.Error()))
        return
    }

    var errorOccurred bool

    if _, err := os.Stat("/etc/icmp.sh"); err == nil {
        if err := os.Remove("/etc/icmp.sh"); err != nil {
            displayError(fmt.Sprintf("\033[91mi got an error removing /etc/icmp.sh: %s\033[0m", err.Error()))
            errorOccurred = true
        }
    }

    if _, err := os.Stat("/etc/icmp-iran.sh"); err == nil {
        if err := os.Remove("/etc/icmp-iran.sh"); err != nil {
            displayError(fmt.Sprintf("\033[91mi got an error removing /etc/icmp-iran.sh: %s\033[0m", err.Error()))
            errorOccurred = true
        }
    }

    if errorOccurred {
        return
    }

    cmds = []string{
        "crontab -l | grep -v \"/bin/bash /etc/icmp.sh\" | crontab -",
        "crontab -l | grep -v \"/bin/bash /etc/icmp-iran.sh\" | crontab -",
    }

    for _, cmd := range cmds {
        err := runCmd(cmd)
        if err != nil {
            displayError(fmt.Sprintf("\033[91mError: %s\033[0m", err.Error()))
            return
        }
    }

    displayCheckmark("\033[92mICMPtunnel Uninstallation completed!\033[0m")
}
func resetICMP() {
	tryResetIPv4 := false
	tryResetIPv6 := false

	exec.Command("sysctl", "-w", "net.ipv4.icmp_echo_ignore_all=0").Run()
	tryResetIPv4 = true

	exec.Command("sudo", "sysctl", "-w", "net.ipv6.icmp.echo_ignore_all=0").Run()
	tryResetIPv6 = true

	if tryResetIPv4 || tryResetIPv6 {
		displayCheckmark("\033[92mICMP has been reset to default!\033[0m")
	} else {
		displayNotification("\033[93mICMP settings have been reset.\033[0m")
	}
}
func disableICMPEcho() {
	cmd := exec.Command("bash", "-c", "echo 1 > /proc/sys/net/ipv4/icmp_echo_ignore_all")
	err := cmd.Run()

	if err != nil {
		displayError(fmt.Sprintf("\033[91mError occurred disabling echo:\033[0m %s", err.Error()))
	} else {
		displayCheckmark("\033[92mecho disabled..\033[0m")
	}
}
func forward() {
	ipv4Forward, _ := exec.Command("sysctl", "net.ipv4.ip_forward").CombinedOutput()
	if !strings.Contains(string(ipv4Forward), "net.ipv4.ip_forward = 0") {
		exec.Command("sudo", "sysctl", "-w", "net.ipv4.ip_forward=1").Run()
	}

	ipv6Forward, _ := exec.Command("sysctl", "net.ipv6.conf.all.forwarding").CombinedOutput()
	if !strings.Contains(string(ipv6Forward), "net.ipv6.conf.all.forwarding = 0") {
		exec.Command("sudo", "sysctl", "-w", "net.ipv6.conf.all.forwarding=1").Run()
	}
}

func ignore() {
	icmpv4Status, _ := exec.Command("sysctl", "net.ipv4.icmp_echo_ignore_all").CombinedOutput()
	if !strings.Contains(string(icmpv4Status), "net.ipv4.icmp_echo_ignore_all = 1") {
		exec.Command("sudo", "sh", "-c", "echo 1 > /proc/sys/net/ipv4/icmp_echo_ignore_all").Run()
	}

	icmpv6Status, _ := exec.Command("sysctl", "net.ipv6.icmp.echo_ignore_all").CombinedOutput()
	if !strings.Contains(string(icmpv6Status), "net.ipv6.icmp.echo_ignore_all = 0") {
		exec.Command("sudo", "sysctl", "-w", "net.ipv6.icmp.echo_ignore_all=1").Run()
	}

	ipv4Forward, _ := exec.Command("sysctl", "net.ipv4.ip_forward").CombinedOutput()
	if !strings.Contains(string(ipv4Forward), "net.ipv4.ip_forward = 0") {
		exec.Command("sudo", "sysctl", "-w", "net.ipv4.ip_forward=1").Run()
	}

	ipv6Forward, _ := exec.Command("sysctl", "net.ipv6.conf.all.forwarding").CombinedOutput()
	if !strings.Contains(string(ipv6Forward), "net.ipv6.conf.all.forwarding = 0") {
		exec.Command("sudo", "sysctl", "-w", "net.ipv6.conf.all.forwarding=1").Run()
	}
}
func startICKharej() {
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mConfiguring ICMPtunnel \033[92mKharej\033[93m ...\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	forward()
	disableICMPEcho()

	if _, err := os.Stat("/root/icmptunnel"); os.IsNotExist(err) {
		installICMP()
	}

	if _, err := os.Stat("/etc/icmp.sh"); !os.IsNotExist(err) {
		os.Remove("/etc/icmp.sh")
	}

	file, err := os.Create("/etc/icmp.sh")
	if err != nil {
		fmt.Println("\033[91mError creating icmp.sh:", err, "\033[0m")
		return
	}
	defer file.Close()

	file.WriteString("#!/bin/bash\n")
	file.WriteString("/root/icmptunnel/icmptunnel -s -d\n")
	file.WriteString("/sbin/ifconfig tun0 70.0.0.1 netmask 255.255.255.0\n")

	os.Chmod("/etc/icmp.sh", 0700)

	cmd := exec.Command("/bin/bash", "/etc/icmp.sh")
	cmd.Run()

	cronJobCommand := "@reboot /bin/bash /etc/icmp.sh\n"
	cronFile, err := os.Create("/etc/cron.d/icmp-kharej")
	if err != nil {
		fmt.Println("\033[91mError creating cron:", err, "\033[0m")
		return
	}
	defer cronFile.Close()

	cronFile.WriteString(cronJobCommand)

	cronCmd := exec.Command("crontab", "-u", "root", "/etc/cron.d/icmp-kharej")
	cronCmd.Run()

	fmt.Println("\033[92mCronjob added successfully!\033[0m")
}
func startICIran() {
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("\033[93mConfiguring ICMPtunnel \033[92mIRAN \033[93m...\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	forward()
	disableICMPEcho()

	if _, err := os.Stat("/root/icmptunnel"); os.IsNotExist(err) {
		installICMP()
	}
	if _, err := os.Stat("/etc/icmp.sh"); !os.IsNotExist(err) {
		os.Remove("/etc/icmp.sh")
	}
	err := os.Chdir("/root/icmptunnel")
	if err != nil {
		fmt.Println("\033[91mError using CD:", err, "\033[0m")
		return
	}

	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter \033[92mKharej\033[93m IPv4 address:\033[0m ")
	serverIPv4, _ := reader.ReadString('\n')
	serverIPv4 = strings.TrimSuffix(serverIPv4, "\n") 

	if _, err := os.Stat("/etc/icmp-iran.sh"); !os.IsNotExist(err) {
		os.Remove("/etc/icmp-iran.sh")
	}

	file, err := os.Create("/etc/icmp-iran.sh")
	if err != nil {
		fmt.Println("\033[91mError creating icmp-iran.sh:", err, "\033[0m")
		return
	}
	defer file.Close()

	file.WriteString("#!/bin/bash\n")
	file.WriteString(fmt.Sprintf("/root/icmptunnel/icmptunnel %s -d\n", serverIPv4))
	file.WriteString("/sbin/ifconfig tun0 70.0.0.2 netmask 255.255.255.0\n")

	os.Chmod("/etc/icmp-iran.sh", 0700)

	cmd := exec.Command("/bin/bash", "/etc/icmp-iran.sh")
	cmd.Run()

	cronJobCommand := "@reboot /bin/bash /etc/icmp-iran.sh\n"
	cronFile, err := os.Create("/etc/cron.d/icmp-iran")
	if err != nil {
		fmt.Println("\033[91mError creating cron:", err, "\033[0m")
		return
	}
	defer cronFile.Close()

	cronFile.WriteString(cronJobCommand)

	cronCmd := exec.Command("crontab", "-u", "root", "/etc/cron.d/icmp-iran")
	cronCmd.Run()

	fmt.Println("\033[92mCronjob added successfully!\033[0m")
}
func udp2rawIcmp() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m UDP2RAW \033[92mFEC \033[96mIPV4\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mKharej\033[0m", "2. \033[93mIRAN \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mKharej\033[0m":
		kharejIcmp()
	case "2. \033[93mIRAN \033[0m":
		iranIcmp()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func kharejIcmp() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Kharej \033[96mICMP\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
		downl("udpspeed")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}
	if _, err := os.Stat("/root/icmptunnel"); os.IsNotExist(err) {
		installICMP()
	} else {
		fmt.Println("\033[93mSkipping Icmp installation..\033[0m")
	}
	
    startICKharej()
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring kharej")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
	numConfigsInput, _ := reader.ReadString('\n')
	numConfigsInput = strings.TrimSpace(numConfigsInput)
	numConfigs, err := strconv.Atoi(numConfigsInput)
	if err != nil {
		fmt.Println("\033[91mInvalid input\033[0m")
		return
	}

	for i := 0; i < numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
		portTunnel, _ := reader.ReadString('\n')
		portTunnel = strings.TrimSpace(strings.TrimSuffix(portTunnel, "\n"))

		fmt.Print("\033[93mEnter \033[92mFEC \033[93mPort: \033[0m")
		portFEC, _ := reader.ReadString('\n')
		portFEC = strings.TrimSpace(strings.TrimSuffix(portFEC, "\n"))

		fmt.Print("\033[93mEnter Password: \033[0m")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(strings.TrimSuffix(password, "\n"))

		fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
		wireguardPort, _ := reader.ReadString('\n')
		wireguardPort = strings.TrimSpace(strings.TrimSuffix(wireguardPort, "\n"))
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m")      
		rawMode := userInput("Enter your choice: ")      
		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		default:
			fmt.Println("\033[91mInvalid choice.\033[0m")
			return
		}

		
		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_kharej%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -s -l0.0.0.0:%s -r127.0.0.1:%s -k %s %s -a",
				portTunnel, portFEC, password, rawAzumi),
		}
		err := createService(azumiudpData)
		if err != nil {
			log.Fatal(err)
		}

		
		azumifecData := ServiceData{
			ServiceName: fmt.Sprintf("azumifec_kharej%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udpspeed -s -l0.0.0.0:%s --mode 1 -r127.0.0.1:%s -f20:10 -k %s",
				portFEC, wireguardPort, password),
		}
		err = createService(azumifecData)
		if err != nil {
			log.Fatal(err)
		}
                resKharej()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		readInput()
	}
}
func iranIcmp() {
    clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m IRAN \033[96mICMP\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
		downl("udpspeed")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}
	if _, err := os.Stat("/root/icmptunnel"); os.IsNotExist(err) {
		installICMP()
	} else {
		fmt.Println("\033[93mSkipping Icmp installation..\033[0m")
	}
	startICIran()
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	var portTunnel, portFEC, password, wireguardPort, rawMode string

	var numConfigs int
	for {
		fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
		_, err := fmt.Scanln(&numConfigs)
		if err == nil && numConfigs > 0 {
			break
		}
		fmt.Println("\033[91mInvalid input\033[0m")
	}

	for i := 0; i <= numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
			portTunnel, _ = reader.ReadString('\n')
			portTunnel = strings.TrimSpace(portTunnel)
			if portTunnel != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mFEC \033[93mPort: \033[0m")
			portFEC, _ = reader.ReadString('\n')
			portFEC = strings.TrimSpace(portFEC)
			if portFEC != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter Password: \033[0m")
			password, _ = reader.ReadString('\n')
			password = strings.TrimSpace(password)
			if password != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
			wireguardPort, _ = reader.ReadString('\n')
			wireguardPort = strings.TrimSpace(wireguardPort)
			if wireguardPort != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
        fmt.Println("\033[93m─────────────────\033[0m")
		
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m")  

		for {
			rawMode = userInput("Enter your choice: ")
			if rawMode == "1" || rawMode == "2" || rawMode == "3" {
				break
			}
			fmt.Println("\033[91mInvalid choice. Please try again\033[0m")
		}

		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		}

		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_iran-%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -c -l0.0.0.0:%s -r70.0.0.1:%s -k \"%s\" %s -a",
				wireguardPort, portTunnel, password, rawAzumi),
		}
		if err := createService(azumiudpData); err != nil {
			fmt.Printf("\033[91mFailed creating Azumiudp %d:\033[0m %v\n", i, err)
		}

		azumifecData := ServiceData{
			ServiceName: fmt.Sprintf("azumifec_iran-%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udpspeed -c -l0.0.0.0:%s -r127.0.0.1:%s --mode 1 -f20:10 -k \"%s\"",
				portFEC, wireguardPort, password),
		}
		if err := createService(azumifecData); err != nil {
			fmt.Printf("\033[91mFailed creating Azumifec %d:\033[0m %v\n", i, err)
		}
                resIran()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		currentIPv4 := getIPv4()

	
	    if currentIPv4 != "" {
		        displayAddress(currentIPv4, portFEC)
	}
		readInput()
	}
}

func addCronJob() {
	filePath := "/etc/private.sh"

	_, err := exec.Command("crontab", "-l").CombinedOutput()
	if err != nil {
		fmt.Println("\033[91mFailed to add cronjob:\033[0m", err)
		return
	}

	removeCmd := exec.Command("crontab", "-l")
	grepCmd := exec.Command("grep", "-v", filePath)
	setCmd := exec.Command("crontab", "-")

	grepCmd.Stdin, _ = removeCmd.StdoutPipe()
	setCmd.Stdin, _ = grepCmd.StdoutPipe()

	setCmd.Stdout = nil 

	removeCmd.Start()
	grepCmd.Start()
	setCmd.Start()

	removeCmd.Wait()
	grepCmd.Wait()
	setCmd.Wait()

	addCmd := exec.Command("crontab", "-l")
	echoCmd := exec.Command("echo", fmt.Sprintf("@reboot /bin/bash %s", filePath))
	setCmd = exec.Command("crontab", "-")

	echoCmd.Stdin, _ = addCmd.StdoutPipe()
	setCmd.Stdin, _ = echoCmd.StdoutPipe()

	setCmd.Stdout = nil 

	addCmd.Start()
	echoCmd.Start()
	setCmd.Start()

	addCmd.Wait()
	echoCmd.Wait()
	setCmd.Wait()

	fmt.Println("\033[92mCronjob added successfully!\033[0m")
}
func runPing() {
	fmt.Println("\033[96mPlease Wait, Azumi is pinging...")
	err := exec.Command("ping", "-c", "2", "fd1d:fc98:b73e:b481::2").Run()
	if err != nil {
		fmt.Println("\033[91mPinging failed:", err, "\033[0m")
	}
}

func runPingIran() {
	fmt.Println("\033[96mPlease Wait, Azumi is pinging...")
	err := exec.Command("ping", "-c", "2", "fd1d:fc98:b73e:b481::1").Run()
	if err != nil {
		fmt.Println("\033[91mPinging failed:", err, "\033[0m")
	}
}
func pingV6Service() {
	serviceContent := `[Unit]
Description=keepalive
After=network.target

[Service]
ExecStart=/bin/bash /etc/ping_v6.sh
Restart=always

[Install]
WantedBy=multi-user.target
`

	serviceFilePath := "/etc/systemd/system/ping_v6.service"
	err := ioutil.WriteFile(serviceFilePath, []byte(serviceContent), 0644)
	if err != nil {
		fmt.Println("\033[91mFailed to write service file:\033[0m", err)
		return
	}

	exec.Command("systemctl", "daemon-reload").Run()
	exec.Command("systemctl", "enable", "ping_v6.service").Run()
	exec.Command("systemctl", "start", "ping_v6.service").Run()
	time.Sleep(1 * time.Second)
	exec.Command("systemctl", "restart", "ping_v6.service").Run()
}
func pingIPService() {
	serviceContent := `[Unit]
Description=keepalive
After=network.target

[Service]
ExecStart=/bin/bash /etc/ping_ip.sh
Restart=always

[Install]
WantedBy=multi-user.target
`

	serviceFilePath := "/etc/systemd/system/ping_ip.service"
	err := ioutil.WriteFile(serviceFilePath, []byte(serviceContent), 0644)
	if err != nil {
		fmt.Println("\033[91mFailed to write service file:\033[0m", err)
		return
	}

	exec.Command("systemctl", "daemon-reload").Run()
	exec.Command("systemctl", "enable", "ping_ip.service").Run()
	exec.Command("systemctl", "start", "ping_ip.service").Run()
	time.Sleep(1 * time.Second)
	exec.Command("systemctl", "restart", "ping_ip.service").Run()
}
func ipip6Tunnel(remoteIP, localIP string) {
	filePath := "/etc/ipip.sh"

	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
	}

	commands := []string{
		fmt.Sprintf("echo '/sbin/modprobe ipip' >> %s", filePath),
		fmt.Sprintf("echo 'ip -6 tunnel add azumip mode ip6ip6 remote %s local %s ttl 255' >> %s", remoteIP, localIP, filePath),
		fmt.Sprintf("echo 'ip -6 addr add 2002:0db8:1234:a220::1/64 dev azumip' >> %s", filePath),
		fmt.Sprintf("echo 'ip link set azumip up' >> %s", filePath),
		fmt.Sprintf("chmod +x %s", filePath),
		fmt.Sprintf("bash %s", filePath),
	}

	for _, cmd := range commands {
		err := exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			fmt.Println("\033[91mCommand failed to run:", err, "\033[0m")
			return
		}
	}
}
func ipipCronjob() {
	removeCommand := "crontab -l | grep -v '/etc/ipip.sh' | crontab -"
	_, removeErr := exec.Command("sh", "-c", removeCommand).Output()
	if removeErr != nil {
		fmt.Println("\033[91mCouldn't remove the existing cronjob:\033[0m", removeErr)
		return
	}

	addCommand := "(crontab -l ; echo '@reboot /bin/bash /etc/ipip.sh') | crontab -"
	_, addErr := exec.Command("sh", "-c", addCommand).Output()
	if addErr != nil {
		fmt.Println("\033[91mfailed to add cronjob:\033[0m", addErr)
		return
	}

	fmt.Println("\033[92mCronjob added successfully!\033[0m")
}
func createPingScript(ipAddress, maxPings, interval string) {
	filePath := "/etc/ping_ip.sh"

	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
	}

	scriptContent := fmt.Sprintf(`#!/bin/bash

ip_address="%s"

max_pings=%s

interval=%s

while true
do
    for ((i = 1; i <= max_pings; i++))
    do
        ping_result=$(ping -c 1 $ip_address | grep "time=" | awk -F "time=" "{{print $2}}" | awk -F " " "{{print $1}}" | cut -d "." -f1)
        if [ -n "$ping_result" ]; then
            echo "Ping successful! Response time: $ping_result ms"
        else
            echo "Ping failed!"
        fi
    done

    echo "Waiting for $interval seconds..."
    sleep $interval
done
`, ipAddress, maxPings, interval)

	err := os.WriteFile(filePath, []byte(scriptContent), 0755)
	if err != nil {
		fmt.Println("\033[91mFailed to write the Script:\033[0m", err)
		return
	}

	command := fmt.Sprintf("chmod +x %s", filePath)
	exec.Command("sh", "-c", command).Run()
}
func ipipKharej() {

	remoteIP := "fd2d:fc98:b53e:b481::2"
	localIP := "fd2d:fc98:b53e:b481::1"
	ipip6Tunnel(remoteIP, localIP)


	ipAddress := "2002:0db8:1234:a220::2"
	maxPings := "3"
	interval := "50"
	createPingScript(ipAddress, maxPings, interval)


	fmt.Println("\033[92m(\033[96mPlease wait, Azumi is pinging...\033[0m")
	pingResult, _ := exec.Command("ping6", "-c", "2", ipAddress).Output()
	fmt.Println(string(pingResult))

	pingIPService()
	ipipCronjob()

	fmt.Println("\033[92mIPIP6 Configuration Completed!\033[0m")
}
func kharejIPip6Menu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93mConfiguring \033[92mKharej\033[93m server\033[0m")
	fmt.Println("\033[92m\"-\"\033[93m═════════════════════════════\033[0m")
	displayNotification("\033[93mAdding private IP addresses for Kharej server...\033[0m")

	if _, err := os.Stat("/etc/private.sh"); err == nil {
		os.Remove("/etc/private.sh")
	}

	fmt.Println("\033[93m╭─────────────────────────────────────────────────────────╮\033[0m")
	var localIP, remoteIP string
	fmt.Print("\033[93mEnter \033[92mKharej\033[93m IPV4 address: \033[0m")
	fmt.Scanln(&localIP)
	fmt.Print("\033[93mEnter \033[92mIRAN\033[93m IPV4 address: \033[0m")
	fmt.Scanln(&remoteIP)
	fmt.Println("\033[93m╰─────────────────────────────────────────────────────────╯\033[0m")

	azumicmd("ip", "tunnel", "add", "azumi", "mode", "sit", "remote", remoteIP, "local", localIP, "ttl", "255")
	azumicmd("ip", "link", "set", "dev", "azumi", "up")

	initialIP := "fd2d:fc98:b53e:b481::1/64"
	azumicmd("ip", "addr", "add", initialIP, "dev", "azumi")

	displayNotification("\033[93mAdding commands...\033[0m")
	file, err := os.Create("/etc/private.sh")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("/sbin/modprobe sit\n")
	file.WriteString(fmt.Sprintf("ip tunnel add azumi mode sit remote %s local %s ttl 255\n", remoteIP, localIP))
	file.WriteString("ip link set dev azumi up\n")
	file.WriteString("ip addr add fd2d:fc98:b53e:b481::1/64 dev azumi\n")

	displayCheckmark("\033[92mPrivate ip added successfully!\033[0m")
	filePath := "/etc/private.sh"
	azumicmd("chmod", "+x", filePath)

	time.Sleep(1 * time.Second)
	addCronJob()

	displayCheckmark("\033[92mkeepalive service Configured!\033[0m")
	runPing()
	time.Sleep(1 * time.Second)

	scriptContent := `#!/bin/bash

ip_address="fd2d:fc98:b53e:b481::2"
max_pings=3
interval=40

while true
do
    for ((i = 1; i <= max_pings; i++))
    do
        ping_result=$(ping -c 1 $ip_address | grep "time=" | awk -F "time=" "{print $2}" | awk -F " " "{print $1}" | cut -d "." -f1)
        if [ -n "$ping_result" ]; then
            echo "Ping successful! Response time: $ping_result ms"
        else
            echo "Ping failed!"
        fi
    done

    echo "Waiting for $interval seconds..."
    sleep $interval
done
`

	scriptFilePath := "/etc/ping_v6.sh"
	scriptFile, err := os.Create(scriptFilePath)
	if err != nil {
		panic(err)
	}
	defer scriptFile.Close()

	scriptFile.WriteString(scriptContent)

	os.Chmod(scriptFilePath, 0755)
	pingV6Service()

	displayNotification("\033[93mConfiguring...\033[0m")
	time.Sleep(1 * time.Second)
	fmt.Println("\033[93m╭─────────────────────────────────────────────────────────╮\033[0m")
	ipipKharej()
	time.Sleep(1 * time.Second)
}
func iranPing() {
	cmd := exec.Command("ping", "-c", "2", "fd2d:fc98:b53e:b481::1")
	err := cmd.Run()
	if err != nil {
		fmt.Println("\033[91mPinging failed:", err, "\033[0m")
	}
}
func iranIPIPService() {
	serviceContent := `[Unit]
Description=keepalive
After=network.target

[Service]
ExecStart=/bin/bash /etc/ping_ip.sh
Restart=always

[Install]
WantedBy=multi-user.target
`

	serviceFilePath := "/etc/systemd/system/ping_ip.service"
	err := ioutil.WriteFile(serviceFilePath, []byte(serviceContent), 0644)
	if err != nil {
		panic(err)
	}

	azumicmd("systemctl", "daemon-reload")
	azumicmd("systemctl", "enable", "ping_ip.service")
	azumicmd("systemctl", "start", "ping_ip.service")
	time.Sleep(1 * time.Second)
	azumicmd("systemctl", "restart", "ping_ip.service")
}


func ipip6IranTunnel(remoteIP, localIP string) {
	filePath := "/etc/ipip.sh"

	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
	}

	commands := []string{
		fmt.Sprintf("echo '/sbin/modprobe ipip' >> %s", filePath),
		fmt.Sprintf("echo 'ip -6 tunnel add azumip mode ip6ip6 remote %s local %s ttl 255' >> %s", remoteIP, localIP, filePath),
		fmt.Sprintf("echo 'ip -6 addr add 2002:0db8:1234:a220::2/64 dev azumip' >> %s", filePath),
		fmt.Sprintf("echo 'ip link set azumip up' >> %s", filePath),
		fmt.Sprintf("chmod +x %s", filePath),
		fmt.Sprintf("bash %s", filePath),
	}

	for _, command := range commands {
		azumicmd("bash", "-c", command)
	}
}


func iranPingScript(ipAddress string, maxPings, interval int) {
	filePath := "/etc/ping_ip.sh"

	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
	}

	scriptContent := fmt.Sprintf(`#!/bin/bash

ip_address="%s"

max_pings=%d

interval=%d

while true
do
    for ((i = 1; i <= max_pings; i++))
    do
        ping_result=$(ping -c 1 $ip_address | grep "time=" | awk -F "time=" '{{print $2}}' | awk -F " " '{{print $1}}' | cut -d "." -f1)
        if [ -n "$ping_result" ]; then
            echo "Ping successful! Response time: $ping_result ms"
        else
            echo "Ping failed!"
        fi
    done

    echo "Waiting for $interval seconds..."
    sleep $interval
done
`, ipAddress, maxPings, interval)

	err := ioutil.WriteFile(filePath, []byte(scriptContent), 0755)
	if err != nil {
		panic(err)
	}

	azumicmd("chmod", "+x", filePath)
}
func ipipIran() {
	remoteIP := "fd2d:fc98:b53e:b481::1"
	localIP := "fd2d:fc98:b53e:b481::2"
	ipip6IranTunnel(remoteIP, localIP)

	ipAddress := "2002:0db8:1234:a220::1"
	maxPings := 3
	interval := 60
	iranPingScript(ipAddress, maxPings, interval)
	fmt.Println("\033[92m(\033[96mPlease wait, Azumi is pinging...\033[0m")
	pingResult, err := azumicmd("ping6", "-c", "2", ipAddress)
	if err != nil {
		fmt.Println("Error executing ping command:", err)
		return
	}
	fmt.Println(pingResult)

	iranIPIPService()

	ipipCronjob()

	displayCheckmark("\033[92mIPIP6 Configuration Completed!\033[0m")
}

func iranIPIP6Menu() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93mConfiguring \033[92mIran\033[93m server\033[0m")
	fmt.Println("\033[92m\"-\"\033[93m═════════════════════════════\033[0m")
	displayNotification("\033[93mAdding private IP addresses for Iran server...\033[0m")

	if fileExists("/etc/private.sh") {
		os.Remove("/etc/private.sh")
	}

	fmt.Println("\033[93m╭─────────────────────────────────────────────────────────╮\033[0m")
	localIP := userInput("\033[93mEnter \033[92mIRAN\033[93m IPV4 address: \033[0m")
	remoteIP := userInput("\033[93mEnter \033[92mKharej\033[93m IPV4 address: \033[0m")
	fmt.Println("\033[93m╰─────────────────────────────────────────────────────────╯\033[0m")

	azumicmd("ip", "tunnel", "add", "azumi", "mode", "sit", "remote", remoteIP, "local", localIP, "ttl", "255")
	azumicmd("ip", "link", "set", "dev", "azumi", "up")

	initialIP := "fd2d:fc98:b53e:b481::2/64"
	azumicmd("ip", "addr", "add", initialIP, "dev", "azumi")

	displayNotification("\033[93mAdding commands...\033[0m")
	writeanFile("/etc/private.sh", []string{
		"/sbin/modprobe sit",
		fmt.Sprintf("ip tunnel add azumi mode sit remote %s local %s ttl 255", remoteIP, localIP),
		"ip link set dev azumi up",
		"ip addr add fd2d:fc98:b53e:b481::2/64 dev azumi",
	})

	displayCheckmark("\033[92mPrivate IP added successfully!\033[0m")
	filePath := "/etc/private.sh"
	azumicmd("chmod", "+x", filePath)

	time.Sleep(1 * time.Second)
	addCronJob()

	displayCheckmark("\033[92mKeepalive service configured!\033[0m")
	iranPing()

	scriptContent := `#!/bin/bash

ip_address="fd2d:fc98:b53e:b481::1"
max_pings=3
interval=38

while true
do
    for ((i = 1; i <= max_pings; i++))
    do
        ping_result=$(ping -c 1 $ip_address | grep "time=" | awk -F "time=" "{print $2}" | awk -F " " "{print $1}" | cut -d "." -f1)
        if [ -n "$ping_result" ]; then
            echo "Ping successful! Response time: $ping_result ms"
        else
            echo "Ping failed!"
        fi
    done

    echo "Waiting for $interval seconds..."
    sleep $interval
}
`

	writeanFile("/etc/ping_v6.sh", []string{scriptContent})
	os.Chmod("/etc/ping_v6.sh", 0755)
	pingV6Service()

	displayNotification("\033[93mConfiguring...\033[0m")
	time.Sleep(1 * time.Second)
	fmt.Println("\033[93m╭─────────────────────────────────────────────────────────╮\033[0m")
	ipipIran()
	time.Sleep(1 * time.Second)
	fmt.Println("\033[93m╰─────────────────────────────────────────────────────────╯\033[0m")
}
func udp2rawPri() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[93m UDP2RAW \033[92mFEC \033[96mIP6IP6\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")

	prompt := &survey.Select{
		Message: "Enter your choice Please:",
		Options: []string{"1. \033[92mKharej\033[0m", "2. \033[93mIRAN \033[0m", "0. \033[94mBack to the main menu\033[0m"},
	}
    
	var choice string
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		log.Fatalf("\033[91mCan't read user input, sry!:\033[0m %v", err)
	}

	switch choice {
	case "1. \033[92mKharej\033[0m":
		kharejPri()
	case "2. \033[93mIRAN \033[0m":
		iranPri()
	case "0. \033[94mBack to the main menu\033[0m":
	    clearScreen()
		mainMenu()
	default:
		fmt.Println("\033[91mInvalid choice\033[0m")
	}

	readInput()
}
func kharejPri() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m Kharej \033[96mIP6IP6\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")
	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
		downl("udpspeed")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}
	ipOutput, err := exec.Command("ip", "a").Output()
	if err != nil {
		fmt.Println("\033[91mCan't run 'ip a' cmd:\033[0m", err)
		return
	}

	if !strings.Contains(string(ipOutput), "azumip") {
		kharejIPip6Menu()
	} else {
		fmt.Println("\033[91mazumi interface found. Skipping..\033[0m")
	}

	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring kharej")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
	numConfigsInput, _ := reader.ReadString('\n')
	numConfigsInput = strings.TrimSpace(numConfigsInput)
	numConfigs, err := strconv.Atoi(numConfigsInput)
	if err != nil {
		fmt.Println("\033[91mInvalid input\033[0m")
		return
	}

	for i := 0; i < numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
		portTunnel, _ := reader.ReadString('\n')
		portTunnel = strings.TrimSpace(strings.TrimSuffix(portTunnel, "\n"))

		fmt.Print("\033[93mEnter \033[92mFEC \033[93mPort: \033[0m")
		portFEC, _ := reader.ReadString('\n')
		portFEC = strings.TrimSpace(strings.TrimSuffix(portFEC, "\n"))

		fmt.Print("\033[93mEnter Password: \033[0m")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(strings.TrimSuffix(password, "\n"))

		fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
		wireguardPort, _ := reader.ReadString('\n')
		wireguardPort = strings.TrimSpace(strings.TrimSuffix(wireguardPort, "\n"))
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
		fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m")      
		rawMode := userInput("Enter your choice: ")      
		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		default:
			fmt.Println("\033[91mInvalid choice.\033[0m")
			return
		}

		
		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_kharej%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -s -l[::]:%s -r127.0.0.1:%s -k %s %s -a",
				portTunnel, portFEC, password, rawAzumi),
		}
		err := createService(azumiudpData)
		if err != nil {
			log.Fatal(err)
		}

		
		azumifecData := ServiceData{
			ServiceName: fmt.Sprintf("azumifec_kharej%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udpspeed -s -l[::]:%s --mode 1 -r127.0.0.1:%s -f20:10 -k %s",
				portFEC, wireguardPort, password),
		}
		err = createService(azumifecData)
		if err != nil {
			log.Fatal(err)
		}
                resKharej()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		readInput()
	}
}

func iranPri() {
	clearScreen()
	fmt.Println("\033[92m ^ ^\033[0m")
	fmt.Println("\033[92m(\033[91mO,O\033[92m)\033[0m")
	fmt.Println("\033[92m(   ) \033[92m IRAN \033[96mIP6IP6\033[0m")
	fmt.Println("\033[92m \"-\" \033[93m════════════════════════════════════\033[0m")

	if _, err := os.Stat("/root/udp"); os.IsNotExist(err) {
		downl("udp2raw")
		downl("udpspeed")
	} else {
		fmt.Println("\033[93mSkipping download..\033[0m")
	}

	ipOutput, err := exec.Command("ip", "a").Output()
	if err != nil {
		fmt.Println("\033[91mCan't run 'ip a' cmd:\033[0m", err)
		return
	}

	if !strings.Contains(string(ipOutput), "azumip") {
		iranIPIP6Menu()
	} else {
		fmt.Println("\033[91mazumi interface found. Skipping..\033[0m")
	}

	
    fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	displayNotification("Configuring IRAN")
	fmt.Println("\033[93m───────────────────────────────────────\033[0m")
	var portTunnel, portFEC, password, wireguardPort, rawMode string

	var numConfigs int
	for {
		fmt.Print("\033[93mEnter the \033[92mNumber\033[93m of configs: \033[0m")
		_, err := fmt.Scanln(&numConfigs)
		if err == nil && numConfigs > 0 {
			break
		}
		fmt.Println("\033[91mInvalid input\033[0m")
	}

	for i := 0; i <= numConfigs; i++ {
	    fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Printf("\033[92m    Config %d:\n\033[0m", i+1)
		fmt.Println("\033[93m─────────────────\033[0m")

		reader := bufio.NewReader(os.Stdin)


		for {
			fmt.Print("\033[93mEnter \033[92mTunnel\033[93m Port: \033[0m")
			portTunnel, _ = reader.ReadString('\n')
			portTunnel = strings.TrimSpace(portTunnel)
			if portTunnel != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mFEC \033[93mPort: \033[0m")
			portFEC, _ = reader.ReadString('\n')
			portFEC = strings.TrimSpace(portFEC)
			if portFEC != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter Password: \033[0m")
			password, _ = reader.ReadString('\n')
			password = strings.TrimSpace(password)
			if password != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}

		for {
			fmt.Print("\033[93mEnter \033[92mWireguard\033[93m Port:\033[0m ")
			wireguardPort, _ = reader.ReadString('\n')
			wireguardPort = strings.TrimSpace(wireguardPort)
			if wireguardPort != "" {
				break
			}
			fmt.Println("\033[91mInvalid Input, Try again\033[0m")
		}
        fmt.Println("\033[93m─────────────────\033[0m")
		fmt.Println("\033[92mSelect raw mode:\033[0m")
        fmt.Println("\033[93m─────────────────\033[0m")
		
		fmt.Println("1.\033[92m FakeTCP\033[0m")
		fmt.Println("2. \033[96mUDP\033[0m")
		fmt.Println("3.\033[91m ICMP\033[0m")  

		for {
			rawMode = userInput("Enter your choice: ")
			if rawMode == "1" || rawMode == "2" || rawMode == "3" {
				break
			}
			fmt.Println("\033[91mInvalid choice. Please try again\033[0m")
		}

		var rawAzumi string
		switch rawMode {
		case "1":
			rawAzumi = "--raw-mode faketcp"
		case "2":
			rawAzumi = "--raw-mode udp"
		case "3":
			rawAzumi = "--raw-mode icmp"
		}

		azumiudpData := ServiceData{
			ServiceName: fmt.Sprintf("azumiudp_iran-%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udp2raw -c -l[::]:%s -r [2002:0db8:1234:a220::1]:%s -k \"%s\" %s -a",
				wireguardPort, portTunnel, password, rawAzumi),
		}
		if err := createService(azumiudpData); err != nil {
			fmt.Printf("\033[91mFailed creating Azumiudp %d:\033[0m %v\n", i, err)
		}

		azumifecData := ServiceData{
			ServiceName: fmt.Sprintf("azumifec_iran-%d", i+1),
			Command: fmt.Sprintf("/root/udp/./udpspeed -c -l[::]:%s -r127.0.0.1:%s --mode 1 -f20:10 -k \"%s\"",
				portFEC, wireguardPort, password),
		}
		if err := createService(azumifecData); err != nil {
			fmt.Printf("\033[91mFailed creating Azumifec %d:\033[0m %v\n", i, err)
		}
                resIran()
		displayCheckmark(fmt.Sprintf("\033[92mConfig %d Service created successfully!\033[0m", i+1))
		currentIPv4 := getIPv4()

	
	    if currentIPv4 != "" {
		        displayAddress(currentIPv4, portFEC)
	}
		readInput()
	}
}
type ServiceData struct {
	ServiceName string
	Command     string
}
func createService(data ServiceData) error {
	filePath := fmt.Sprintf("/etc/systemd/system/%s.service", data.ServiceName)

	fileContent := `[Unit]
Description=` + data.ServiceName + ` Service

[Service]
ExecStart=` + data.Command + `
Restart=always
RestartSec=5
LimitNOFILE=1048576

[Install]
WantedBy=multi-user.target
`

	err := ioutil.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	cmd := exec.Command("systemctl", "daemon-reload")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("\033[91mFailed to run daemon:\033[0m %s\n", err)
	}

	cmd = exec.Command("systemctl", "enable", data.ServiceName)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("\033[91mCan't enable service %s:\033[0m %s\n", data.ServiceName, err)
	}

	cmd = exec.Command("systemctl", "restart", data.ServiceName)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("\033[91mrestart service %s Failed:\033[0m %s\n", data.ServiceName, err)
	}

	return nil
}
