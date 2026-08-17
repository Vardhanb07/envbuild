/*
Copyright © 2026 Vardhan Battula <vardhanbattula7@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Env []Env `mapstructure:"env"`
}

type Env struct {
	Key string `mapstructure:"key"`
	Cmd string `mapstructure:"cmd"`
}

func Add(path, key, cmd string) error {
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	cfg.Env = append(cfg.Env, Env{Key: key, Cmd: cmd})
	viper.Set("env", cfg.Env)
	return viper.WriteConfig()
}

func Update(path, key, cmd string) error {
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	for _, env := range cfg.Env {
		if env.Key == key {
			env.Cmd = cmd
		}
	}
	viper.Set("env", cfg.Env)
	return viper.WriteConfig()
}

func Delete(path, key string) error {
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	var newCfg Config
	for _, env := range cfg.Env {
		if env.Key != key {
			newCfg.Env = append(newCfg.Env, env)
		}
	}
	viper.Set("env", newCfg.Env)
	return viper.WriteConfig()
}

func Read(path string) (Config, error) {
	viper.SetConfigType("toml")
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
