package action

import (
	"fmt"
	"os"
	"strings"
)

func AddPath(path string) error {
	Context.Path = fmt.Sprintf("%s:%s", Context.Path, path)
	return os.Setenv("GITHUB_PATH", Context.Path)
}

func Debug(msg string) {
	cmd("debug", msg, map[string]string{})
}

func Notice(msg string, opts map[string]string) {
	cmd("notice", msg, opts)
}

func Warning(msg string, opts map[string]string) {
	cmd("warning", msg, opts)
}

func Error(err error, opts map[string]string) {
	cmd("error", err.Error(), opts)
}

func StartGroup(title string) {
	cmd("group", title, map[string]string{})
}

func EndGroup() {
	cmd("endgroup", "", map[string]string{})
}

func GetInput(name string) string {
	s := fmt.Sprintf("INPUT_%s", name)
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, " ", "_")

	return os.Getenv(s)
}

func GetState(name string) string {
	return os.Getenv(fmt.Sprintf("%s_%s", "STATE_", name))
}

func SaveState(name string, value string) {
	cmd("save-state", value, map[string]string{"name": name})
}

func IsDebug() bool {
	return Context.RunnerDebug
}

func SetCommandEcho(value string) {
	cmd("echo", value, map[string]string{})
}

func SetFailed(err error, opts map[string]string) {
	Error(err, opts)
	os.Exit(1)
}

func SetOutput(name, value string) {
	cmd("set-output", value, map[string]string{"name": name})
}

func SetSecret(secret string) {
	cmd("add-mask", secret, map[string]string{})
}

func cmd(name, value string, opts map[string]string) {
	fmtStr := "::%s %s::%s\n"

	var parsedOpts []string
	for k, v := range opts {
		parsedOpts = append(parsedOpts, fmt.Sprintf("%s=%s", k, v))
	}

	fmt.Printf(fmtStr, name, strings.Join(parsedOpts, ","), value)
}
