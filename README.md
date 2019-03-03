# Helm-Lite 

An utility tool for kubernetes. Think of it as a lightweight and minimal Helm. Allows for templating of YAML k8 objects or docker-compose files. Multiple config files and templates allow you to only need one copy of each file for all your different environments.

## How to use it

build the program and rename it
```
go build main.go
```

### Generate a release for a given environment

Build a version (folder name) using values from `(env target).conf` This config file can either be in the root of the version directory or in the `config/` folder which is a subdirectory of the version directory. 

In this case we will build the dev version of v0.1 (fills in all templates with the values from dev.conf)
```
./helm-lite gen v0.1 dev
```

The generated files go under target/dev/v0.1/

### Generates a new release version by copying the specified release

## Authors

* **Zachary Spofford** - *Initial work* - [Spoofardio](https://github.com/Spoofardio)