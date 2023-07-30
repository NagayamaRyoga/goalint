package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"text/template"
)

func main() {
	ctx := context.Background()

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <package>\n", os.Args[0])
		os.Exit(1)
	}

	targetPackage := os.Args[1]

	exitCode, err := run(ctx, targetPackage)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(exitCode)
}

func run(ctx context.Context, targetPackage string) (int, error) {
	tmpDir, err := os.MkdirTemp(".", "goalint*")
	if err != nil {
		return -1, fmt.Errorf("os.MkdirTemp() failed: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	exeExt := ".out"
	if runtime.GOOS == "windows" {
		exeExt = ".exe"
	}

	mainPath := tmpDir + "/main.go"
	exePath := tmpDir + "/a" + exeExt
	out, err := os.OpenFile(mainPath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return -1, fmt.Errorf("os.OpenFile(%q) failed: %w", mainPath, err)
	}
	defer out.Close()

	data := map[string]any{
		"TargetPackage": targetPackage,
	}

	tpl := template.Must(template.New("main").Parse(tplMain))
	if err := tpl.Execute(out, data); err != nil {
		return -1, fmt.Errorf("tpl.Execute() failed: %w", err)
	}

	if _, err := runCommand(ctx, "go", "build", "-o", exePath, mainPath); err != nil {
		return -1, fmt.Errorf("cmd.Run() failed: %w", err)
	}

	exitCode, err := runCommand(ctx, exePath)
	if err != nil && exitCode < 0 {
		return -1, err
	}

	return exitCode, nil
}

func runCommand(ctx context.Context, name string, args ...string) (int, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return cmd.ProcessState.ExitCode(), fmt.Errorf("exec.Command(%q).Run() failed: %w", cmd, err)
	}

	return cmd.ProcessState.ExitCode(), nil
}

const tplMain string = `package main

import (
	"fmt"
	"os"

	_ {{ printf "%q" .TargetPackage }}
	"github.com/NagayamaRyoga/goalint"
	"github.com/NagayamaRyoga/goalint/inner/config"
	"github.com/NagayamaRyoga/goalint/inner/runner"
	"goa.design/goa/v3/eval"
)

func main() {
	if err := eval.RunDSL(); err != nil {
		panic(err)
	}

	roots, err := eval.Context.Roots()
	if err != nil {
		panic(err)
	}

	cfg := config.NewConfig()
	cfg.Disabled = false

	if lint.Configurator != nil {
		lint.Configurator(cfg)
	}

	if err := runner.Run(cfg, {{ printf "%q" .TargetPackage }}, roots); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`
