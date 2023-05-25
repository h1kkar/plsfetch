package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func echo(str string) string {
	cmd := exec.Command("bash", "-c", "echo $"+strings.ToUpper(str))
	out, err := cmd.Output()

	if err != nil {
		panic(err)
	} else {
		result := strings.Split(string(out), "\n")

		return result[0]
	}
}

func run(str string) string {
	cmd := exec.Command("bash", "-c", str)
	out, err := cmd.Output()

	if err != nil {
		panic(err)
	} else {
		result := strings.Split(string(out), "\n")

		return result[0]
	}
}

func user() string {
	return run("whoami")
}

func host() string {
	return run("uname -n")
}

func distro() string {
	distros := run("cat /etc/lsb-release")
	out := strings.Split(distros, "=")

	return out[1]
}

func arch() string {
	return run("uname -m")
}

func kernel() string {
	return run("uname -r")
}

func up() string {
	uptime := run("uptime -p")
	out := strings.Split(uptime, "up")
	out = strings.Split(out[1], " ")

	return out[1] + "h " + out[3] + "m"
}

func shell() string {
	return echo("shell")
}

func wm() string {
	return echo("desktop_session")
}

func term() string {
	terminal := run("w -h")
	arr := strings.Split(terminal, " ")
	arr = strings.Split(arr[len(arr)-1], "/")

	return arr[len(arr)-1]
}

var d string = "\033[0m"
var r string = string("\033[31m")
var g string = string("\033[32m")
var y string = string("\033[33m")
var b string = string("\033[34m")
var p string = string("\033[35m")
var c string = string("\033[36m")
var w string = string("\033[37m")

func color() string {
	out := r + "██ " + g + "██ " + y + "██ " + b + "██ " + p + "██ " + c + "██ " + w + "██ " + d + "██ "
	return out
}

func ver() string {
	return "plsfetch v0.5"
}

func out() {
	fmt.Println("   " + b + user() + d + " @ " + b + host())
	fmt.Println(b + "distro" + w + " · " + b + distro() + " " + arch())
	fmt.Println(g + "kernel" + w + " · " + g + kernel())
	fmt.Println(y + "uptime" + w + " · " + y + up())
	fmt.Println(p + " shell" + w + " · " + p + shell())
	fmt.Println(r + "    wm" + w + " · " + r + wm())
	//fmt.Println("\n" + color())
}

func main() {

	if len(os.Args[1:]) != 0 {
		switch os.Args[1] {
		case "ver":
			fmt.Println(ver())
		default:
			out()
		}
	} else {
		out()
	}

}
