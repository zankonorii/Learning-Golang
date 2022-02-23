package main

/*
 *	go.mod is something like package.json in js or composer.json in php
 *	go.sum is locked of that like package.json.lock or composer.json.lock
 *
 * === === === === === === === === === === === === === === === === ===
 *
 *	module GolandProjects	=> this is project name
 *
 *	go 1.17				=>	this is version of Go
 *		we can add more quality like 1.17.3 for more حساسیت
 *
 *	go get == yarn add == npm install == composer install
 *
 *	go get add automatically packages to this section
 *
 *	with go run command can run install this package file => go run go.mod
 *
 *	go mod == module maintenance
 *
 *	go mod init AName.mod == create a config file like go.mod
 *
 *	for more information => go help mode
 *
 *
 *	go mod tidy == for clearing go.mode file => remove spaces & comments
 *	go mod tiny is for upgrading too. like composer update Note: from go 1.17 <
 *
 *	go build => for compile source and return an execute file
 *
 *	go.sum => this is lock file just like composer.json.lock in laravel or php projects
 *
 *	go doc => go Documentation
 *
 *	go fmt => format
 *
 * go generate => generate Go file by processing source
 *
 *	go get => get and install new package
 *
 *	go compile => compile and install packages and dependencies	=> go help install
 *
 *	go list => list packages or modules	=> go list -m all => this show all installed packages
 *										=> go list -m     => this show over installed packages
 *
 *	go test => for testings
 *
 *	go version => print Go version
 *
 *	go vet	=> report likely mistakes in packages
 *
 *	go tool  ???! =>I will read that
 *
 *	require (
 *		github.com/pelletier/go-toml v1.9.4 // indirect
 *		github.com/pkg/errors v0.9.1 // indirect
 *	)
 *
 *	thanks from amir hasan azimi ,
 *
 *
 */
