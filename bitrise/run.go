package bitrise

import (
	"os"
	"os/exec"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// ------------------
// --- Stepman

// RunStepmanSetup ...
func RunStepmanSetup(collection string) error {
	logLevel := log.GetLevel().String()
	args := []string{"--debug", "--loglevel", logLevel, "setup", "--collection", collection}
	return RunCommand("stepman", args...)
}

// RunStepmanActivate ...
func RunStepmanActivate(collection, stepID, stepVersion, dir, ymlPth string) error {
	logLevel := log.GetLevel().String()
	args := []string{"--debug", "--loglevel", logLevel, "activate", "--collection", collection, "--id", stepID, "--version", stepVersion, "--path", dir, "--copyyml", ymlPth}
	return RunCommand("stepman", args...)
}

// ------------------
// --- Envman

// RunEnvmanInit ...
func RunEnvmanInit() error {
	logLevel := log.GetLevel().String()
	args := []string{"--loglevel", logLevel, "init"}
	return RunCommand("envman", args...)
}

// RunEnvmanAdd ...
func RunEnvmanAdd(key, value string, expand bool) error {
	logLevel := log.GetLevel().String()
	args := []string{"--loglevel", logLevel, "add", "--key", key}
	if !expand {
		args = []string{"--loglevel", logLevel, "add", "--key", key, "--no-expand"}
	}

	envman := exec.Command("envman", args...)
	envman.Stdin = strings.NewReader(value)
	envman.Stdout = os.Stdout
	envman.Stderr = os.Stderr
	return envman.Run()
}

// RunEnvmanRunInDir ...
func RunEnvmanRunInDir(dir string, cmd []string) error {
	logLevel := log.GetLevel().String()
	args := []string{"--loglevel", logLevel, "run"}
	args = append(args, cmd...)
	return RunCommandInDir(dir, "envman", args...)
}

// RunEnvmanRun ...
func RunEnvmanRun(cmd []string) error {
	return RunEnvmanRunInDir("", cmd)
}

// RunEnvmanEnvstoreTest ...
func RunEnvmanEnvstoreTest(pth string) error {
	logLevel := log.GetLevel().String()
	args := []string{"--loglevel", logLevel, "--path", pth, "print"}
	cmd := exec.Command("envman", args...)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ------------------
// --- Common

// RunCopy ...
func RunCopy(src, dst string) error {
	args := []string{src, dst}
	return RunCommand("cp", args...)
}

// RunCommand ...
func RunCommand(name string, args ...string) error {
	return RunCommandInDir("", name, args...)
}

// RunCommandInDir ...
func RunCommandInDir(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if dir != "" {
		cmd.Dir = dir
	}
	log.Debugln("Run command: $", cmd)
	return cmd.Run()
}