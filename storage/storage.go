package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	homedir "github.com/mitchellh/go-homedir"
)

var Home string

const adventurersDir = ".TheAdventurersToolbelt"

func Bootstrap() {
	home, err := homedir.Dir()
	if err != nil {
		logrus.WithField("error", err).Fatal("could not find user home")
	}
	logrus.WithField("user home", home).Debug("init storage")

	perms := os.FileMode(0755)

	Home = filepath.Join(home, adventurersDir)
	err = os.MkdirAll(Home, perms)
	if err != nil {
		logrus.WithField("error", err).WithField("path", Home).Fatal("could not create app home dir")
	}

	subfolders := []string{"characters", "settings", "rules"}

	for _, folder := range subfolders {
		fullPath := filepath.Join(Home, folder)
		err = os.MkdirAll(fullPath, perms)
		if err != nil {
			logrus.WithField("error", err).WithField("path", fullPath).Fatal("could not create app dir")
		} else {
			logrus.WithField("path", fullPath).Debug("Created app dir")
		}
	}
}

func List() ([]os.FileInfo, error) {
	return ioutil.ReadDir(Home)
}
