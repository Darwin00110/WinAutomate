package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ErrorPainel() {
	fmt.Println(`
	---------------------------------------------------------------
	---------------------------------------------------------------
             ________  _______   _______    ______   _______  
            |        \|       \ |       \  /      \ |       \ 
            | $$$$$$$$| $$$$$$$\| $$$$$$$\|  $$$$$$\| $$$$$$$\
            | $$__    | $$__| $$| $$__| $$| $$  | $$| $$__| $$
            | $$  \   | $$    $$| $$    $$| $$  | $$| $$    $$
            | $$$$$   | $$$$$$$\| $$$$$$$\| $$  | $$| $$$$$$$\
            | $$_____ | $$  | $$| $$  | $$| $$__/ $$| $$  | $$
            | $$     \| $$  | $$| $$  | $$ \$$    $$| $$  | $$
             \$$$$$$$$ \$$   \$$ \$$   \$$  \$$$$$$  \$$   \$$
	---------------------------------------------------------------
	---------------------------------------------------------------
	
	======================================================
	Please verify your param in the application, try again.
	======================================================
	`)
}

func InfoPainel() {
	fmt.Println(`
		-------------------------------------------------------------
		-------------------------------------------------------------
			     ______  __    __  ________   ______  
			    |      \|  \  |  \|        \ /      \ 
			     \$$$$$$| $$\ | $$| $$$$$$$$|  $$$$$$\
			      | $$  | $$$\| $$| $$__    | $$  | $$
			      | $$  | $$$$\ $$| $$  \   | $$  | $$
			      | $$  | $$\$$ $$| $$$$$   | $$  | $$
			     _| $$_ | $$ \$$$$| $$      | $$__/ $$
			    |   $$ \| $$  \$$$| $$       \$$    $$
			     \$$$$$$ \$$   \$$ \$$        \$$$$$$ 
                                      
		-------------------------------------------------------------
		-------------------------------------------------------------

		==Commands==
		
		--matar processo: "--killProcess --name <name>" 
		--abrir programas/Pastas: "--openProgram --path <path>"
		--limpar arquivos temporarios ou cache: "--cleanCacheFiles -confirm"

		======================================================
		`)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func SucessPainel() {
	fmt.Println(`
-------------------------------------------------------------------------
-------------------------------------------------------------------------
      ______   __    __   ______   ________   ______    ______  
     /      \ |  \  |  \ /      \ |        \ /      \  /      \ 
    |  $$$$$$\| $$  | $$|  $$$$$$\| $$$$$$$$|  $$$$$$\|  $$$$$$\
    | $$___\$$| $$  | $$| $$   \$$| $$__    | $$___\$$| $$___\$$
     \$$    \ | $$  | $$| $$      | $$  \    \$$    \  \$$    \ 
     _\$$$$$$\| $$  | $$| $$   __ | $$$$$    _\$$$$$$\ _\$$$$$$\
    |  \__| $$| $$__/ $$| $$__/  \| $$_____ |  \__| $$|  \__| $$
     \$$    $$ \$$    $$ \$$    $$| $$     \ \$$    $$ \$$    $$
      \$$$$$$   \$$$$$$   \$$$$$$  \$$$$$$$$  \$$$$$$   \$$$$$$ 
	  
-------------------------------------------------------------------------
-------------------------------------------------------------------------

   ========================
   Press any key for quit!
   ========================	
   `)
}

func CoreApp(args []string) {
	if len(args) < 2 {
		ErrorPainel()
		return
	}

	CleanTemporatyFiles := func(pathFiles string) {
		cmdStr := fmt.Sprintf("Remove-Item -Path '%s' -Force -Recurse", pathFiles)
		cmd := exec.Command("powershell", "-Command", cmdStr)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Erro:", err)
			ErrorPainel()
			return
		}
		fmt.Println(string(out))
		SucessPainel()
	}

	Command := args[0]
	argFlag := args[1]
	target := ""
	if len(args) > 2 {
		target = args[2]
	}

	// Path temporário do Windows
	tempPathUsers := filepath.Join(os.Getenv("TEMP"), "*")
	tempPathWindows := filepath.Join(os.Getenv("SystemRoot"), "Temp", "*")
	tempPathPrefetch := filepath.Join(os.Getenv("SystemDrive"), "Prefetch", "*")

	switch Command {
	case "--killProcess":
		if argFlag == "--name" && target != "" {
			cmd := exec.Command("cmd", "/c", "Taskkill", "/IM", target)
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Erro:", err)
				ErrorPainel()
			}
			fmt.Println(string(out))
			SucessPainel()
		} else {
			ErrorPainel()
		}

	case "--openProgram":
		if argFlag == "--path" && target != "" {
			if PathExists(target) {
				cmd := exec.Command("cmd", "/c", "start", "", target)
				out, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Println("Erro:", err)
					ErrorPainel()
					return
				}
				fmt.Println(string(out))
				SucessPainel()

			} else {
				ErrorPainel()
			}
		} else {
			ErrorPainel()
		}

	case "--cleanCacheFiles":
		if argFlag == "-confirm" && target == "" {
			CleanTemporatyFiles(tempPathUsers)
			CleanTemporatyFiles(tempPathWindows)
			CleanTemporatyFiles(tempPathPrefetch)
			SucessPainel()
		} else {
			ErrorPainel()
		}

	default:
		ErrorPainel()
	}
}

func main() {

	Args := os.Args[1:] // ignora o executável
	if len(Args) == 0 {
		InfoPainel()
		return
	}

	CoreApp(Args)
	fmt.Scanln()
}
