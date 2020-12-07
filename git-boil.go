package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type remotes struct {
	Remotes []struct {
		Repo struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		} `yaml:"repo"`
	} `yaml:"remotes"`
}

func getConf() (parsed *remotes) {

	buf, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		panic("conf.yaml - not in path")
	}

	conf := &remotes{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		panic("getCred()")
	}

	return conf
}

func genRemotes(project string) []string {
	repos := getConf()
	remotes := make([]string, len(repos.Remotes))

	for idx, repo := range repos.Remotes {
		remotes[idx] = fmt.Sprintf("%v/%v.git", repo.Repo.URL, project)
	}

	return remotes
}

func runcmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Fatal(err)
			panic("some error found")
		}
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func remoteInit(remotes []string) {
	cmd := fmt.Sprintf("git remote add origin %v", remotes[0])
	fmt.Println(cmd)
	runcmd(cmd, true)

	for i := 0; i < len(remotes); i++ {
		cmd = fmt.Sprintf("git remote set-url --add --push origin %v", remotes[i])
		runcmd(cmd, true)
	}
}

// TODO(#1): Interact with gitea/github api inorder to create repositories
// func createRepo()

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}

	remotes := genRemotes(os.Args[1])
	if len(remotes) < 1 {
		fmt.Println("Empty configuration")
		os.Exit(2)
	}

	remoteInit(remotes)
}
