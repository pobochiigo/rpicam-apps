package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("rpicam-vid",
		"-t", "0",
		"-v", "2","-n",
		"--metadata", "-",
		"--post-process-file", "/usr/share/rpi-camera-assets/hailo_yolov8_inference.json",
	)

	// Get stdout pipe to read inference logs
	stdout, err := cmd.StderrPipe()
	if err != nil {
		panic(fmt.Sprintf("Failed to get stdout: %v", err))
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		panic(fmt.Sprintf("Failed to start rpicam-hello: %v", err))
	}

	// Read line by line
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
	    line := scanner.Text()
	    fmt.Println(line)

		// Optional: parse JSON lines here
		//if strings.HasPrefix(line, "{") {}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scan error:", err)
	}

	cmd.Wait()
}
