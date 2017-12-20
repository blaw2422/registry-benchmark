package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	log.SetOutput(os.Stderr)
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <image[:tag]> <iterations>", os.Args[0])
	}

	image := os.Args[1]
	iter, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Error parsing iterations count: %v", err)
	}

	// Clean after we finish the run
	defer rmi(image) // ignore error

	tryLogin()

	log.Printf("Image: %s", image)
	for i := 0; i < iter; i++ {
		log.Printf("Run %d:", i)
		if err := rmi(image); err != nil {
			log.Fatal(err)
		}

		start := time.Now()
		if err := pull(image); err != nil {
			log.Fatal(err)
		}
		d := time.Since(start)
		fmt.Printf("%d\n", int64(d.Seconds()))
	}

}

func tryLogin() {
  var flags []string
  flags = append(flags, "login")

  if os.Getenv("DOCKER_USERNAME") != "" {
    flags = append(flags, "-u", os.Getenv("DOCKER_USERNAME"))
  }

  if os.Getenv("DOCKER_PASSWORD") != "" {
    flags = append(flags, "-p", os.Getenv("DOCKER_PASSWORD"))
  }

  if os.Getenv("DOCKER_EMAIL") != "" {
    flags = append(flags, "-e", os.Getenv("DOCKER_EMAIL"))
  }

  if os.Getenv("DOCKER_SERVER") != "" {
    flags = append(flags, os.Getenv("DOCKER_SERVER"))
  }

  if res, err := dockerCmd(flags...); err != nil {
    log.Fatalf("login failed: %s: %s", res, err)
  }
}

func rmi(image string) error {
	// See if image exists
	if _, err := dockerCmd("inspect", image); err != nil {
		// image doesn't exist (or some other failure which will come up below)
		return nil
	}

	if b, err := dockerCmd("rmi", "-f", image); err != nil {
		return fmt.Errorf("rmi failed: %v: %q", err, string(b))
	}
	return nil
}

func pull(image string) error {
	if b, err := dockerCmd("pull", image); err != nil {
		return fmt.Errorf("pull failed: %v: %q", err, string(b))
	}
	return nil
}

func dockerCmd(args ...string) ([]byte, error) {
	if os.Getenv("DOCKER_CERT_PATH") != "" {
		args = append([]string{"--tls"}, args...)
	}
	cmd := exec.Command("docker", args...)
	return cmd.CombinedOutput()
}
