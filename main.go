package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
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

func symb(num string) string {
	return "\033[3" + num + "m"
}

var (
	d string = "\033[0m"
	r string = symb("1")
	g string = symb("2")
	y string = symb("3")
	b string = symb("4")
	p string = symb("5")
	c string = symb("6")
	w string = symb("7")
)

func color() string {
	out := r + "██ " + g + "██ " + y + "██ " + b + "██ " + p + "██ " + c + "██ " + w + "██ " + d + "██ "
	return out
}

func ver() string {
	return "plsfetch v0.5"
}

func out() {
	num := len(user())
	if num <= 6 {
		num = 6
	}
	number := strconv.Itoa(num)

	fmt.Printf(b+"%"+number+"s"+d+" @ "+b+"%s\n", user(), host())
	fmt.Printf(b+"%"+number+"s"+w+" · "+b+"%s %s\n", "distro", distro(), arch())
	fmt.Printf(g+"%"+number+"s"+w+" · "+g+"%s\n", "kernel", kernel())
	fmt.Printf(y+"%"+number+"s"+w+" · "+y+"%s\n", "uptime", up())
	fmt.Printf(p+"%"+number+"s"+w+" · "+p+"%s\n", "shell", shell())
	fmt.Printf(r+"%"+number+"s"+w+" · "+r+"%s\n", "wm", wm())
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
