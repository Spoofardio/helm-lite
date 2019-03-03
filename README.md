# Helm-Lite 

An utility tool for kubernetes. Think of it as a lightweight and minimal Helm.

## How to use it

build the program and rename it
```
go build tuner-release-tool/main.go && mv main hl
```

### Generate a release for a given environment

Build a version (folder name) using values from `(env target).conf` This config file can either be in the root of the version directory or in the `config/` folder which is a subdirectory of the version directory. 

In this case we will build the dev version of v0.1 (fills in all templates with the values from dev.conf)
```
./hl gen v0.1 dev
```

The generated files go under target/dev/v0.1/

### Generates a new release version by copying the specified release

Makes an exact copy of an old version and uses that to create a new version. This feature is not complete yet.

In this case it copies the version v0.1 to make the new version v0.2
```
./hl new v0.1 v0.2
```

## Authors

* **Zachary Spofford** - *Initial work* - [Spoofardio](https://github.com/Spoofardio)