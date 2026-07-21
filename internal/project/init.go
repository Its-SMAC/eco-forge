package project

import (
	"errors"
	"fmt"
	"forge/internal/build"
	"forge/internal/manifest"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
)

type Project struct {
	workDir     string
	projectName string
	framework   string
	docker      bool
}

func Init() error {

	var (
		project   Project
		currentWD bool
		err       error
	)

	project.workDir, err = os.Getwd()
	if err != nil {
		return err
	}

	err = huh.NewConfirm().
		Title(fmt.Sprintf("Is this the correct directory?\n%s", project.workDir)).
		Value(&currentWD).
		Run()

	if err != nil {
		return err
	}

	if !currentWD {
		err = huh.NewInput().
			Title("Enter the project directory").
			Value(&project.workDir).
			Run()

		if err != nil {
			return err
		}
	}

	if err = verifyIfExistProject(project); err != nil {
		return err
	}

	err = huh.NewInput().
		Title("Project name").
		Value(&project.projectName).
		Run()

	if err != nil {
		return err
	}

	err = huh.NewSelect[string]().
		Title("Framework").
		Options(
			huh.NewOption("Gin", "gin"),
			huh.NewOption("Chi", "chi"),
			huh.NewOption("net/http", "net/http"),
		).
		Value(&project.framework).
		Run()

	if err != nil {
		return err
	}

	err = huh.NewConfirm().
		Title("Use Docker?").
		Value(&project.docker).
		Run()

	if err != nil {
		return err
	}

	createFiles(project)

	return nil
}

func createFiles(p Project) error {
	content := fmt.Sprintf(string(manifest.ForgeToml), p.projectName, p.framework, "Description", "Ingots", build.Version)

	err := os.WriteFile(filepath.Join(p.workDir, "forge.toml"), []byte(content), 0o644)
	if err != nil {
		return err
	}

	if p.docker {
		err = os.WriteFile(filepath.Join(p.workDir, "docker-compose.yml"), []byte("version: 0.0.1"), 0o644)
		if err != nil {
			return err
		}
	}

	return nil
}

func verifyIfExistProject(p Project) error {
	_, err := os.Stat(filepath.Join(p.workDir, "forge.toml"))
	if err == nil {
		return errors.New("This directory already contains a Forge project")
	}

	return nil
}
