package config

import (
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

type Options struct {
	paths []string
	dst   any
}

type Option func(opts *Options)

func WithPath(path string) Option {
	return func(opts *Options) {
		opts.paths = append(opts.paths, path)
	}
}

func WithDestination(dst any) Option {
	return func(opts *Options) {
		opts.dst = dst
	}
}

func ReadInConfig(opts ...Option) error {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}

	if err := validate(options); nil != err {
		return err
	}

	return readYaml(options)
}

func validate(opts Options) error {
	if len(opts.paths) < 1 {
		return fmt.Errorf("path not found")
	}

	if reflect.TypeOf(opts.dst).Kind() != reflect.Ptr {
		return fmt.Errorf("destination not ptr")
	}

	return nil
}

func readYaml(opts Options) error {
	for _, path := range opts.paths {
		f, err := os.Open(path)
		if nil == err {
			err = yaml.NewDecoder(f).Decode(opts.dst)
		}

		if nil != err {
			return err
		}
	}

	return nil
}
