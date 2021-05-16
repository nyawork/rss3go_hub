/*********************************************************************

rss3go: An alternative version of RSSHub for RSS3 written in go

Copyright (C) 2021 Nyawork, Candinya

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

 ********************************************************************/

package storage

import (
	"log"
	"os"
	"rss3go/config"
)

type TypeOfStorageUndefinedError struct {
	sType string
}

func (e *TypeOfStorageUndefinedError) Error() string {
	return "Storage type undefined: sType"
}

func Write(name string, content []byte) error {
	if config.GlobalConfig.Storage.Type == "local" {
		err := os.WriteFile(config.GlobalConfig.Storage.Path + name, content, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		return err
	}
	return &TypeOfStorageUndefinedError{config.GlobalConfig.Storage.Type}
}

func Read(name string) ([]byte, error) {
	if config.GlobalConfig.Storage.Type == "local" {
		data, err := os.ReadFile(config.GlobalConfig.Storage.Path + name)
		if err != nil {
			log.Fatalln(err)
		}
		return data, err
	}
	return nil, &TypeOfStorageUndefinedError{config.GlobalConfig.Storage.Type}
}

func Exist(name string) (bool, error) {
	if config.GlobalConfig.Storage.Type == "local" {
		_, err := os.Stat(config.GlobalConfig.Storage.Path + name)
		fileExist := os.IsNotExist(err)
		if !fileExist && err != nil {
			log.Fatalln(err)
		}
		return fileExist, err
	}
	return false, &TypeOfStorageUndefinedError{config.GlobalConfig.Storage.Type}
}

func Rm(name string) error {
	if config.GlobalConfig.Storage.Type == "local" {
		err := os.Remove(config.GlobalConfig.Storage.Path + name)
		if err != nil {
			log.Fatalln(err)
		}
		return err
	}
	return &TypeOfStorageUndefinedError{config.GlobalConfig.Storage.Type}
}
