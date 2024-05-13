# Captain Hook

## Description

Captain Hook is a simple, local first webhook tester written in Go to help you test out your webhooks. while it is designed for use locally you can test it the usual way a la ngrok by deploying the [server component](https://github.com/nznznz42/Captain-Hook-Server) to a webhost of your choice.

## Prerequisites

To use Captain Hook you will need to have the go compiler and package manager installed on your machine. You can find both [here](https://go.dev/dl/).

once you've installed go you will need to install the following dependencies:

1. [Cobra](github.com/spf13/cobra)
2. [Toml](github.com/BurntSushi/toml)
3. [pflag](github.com/spf13/pflag)

you can install them by running:

``` shell
go get <link-to-package>
```

## Installation

To install and run Captain Hook simply do the following:

```shell
git clone https://github.com/nznznz42/Captain-Hook.git
cd Captain-Hook
go build
```

## Using Captain Hook

This project is designed to be used either locally or in conjunction with the server component which you can find [here](https://github.com/nznznz42/Captain-Hook-Server).

### Commands

Captain Hook provides you with the following commands:

#### 1. ltest

This command is used when you want to test your webhook locally. you will need to first write a config file and place it in the Config Folder and provide an example payload in the Payload folder, an example of both have been provided in the respective folders.

To use the command you must run the following command replacing the arguments with your own values:

```shell
hooktest <ConfigFileName> <LogFileName> [Flags]
```

Provide the **-r** flag if you wish to randomise your payload values.

#### 2. ctest

This command is used when you want to test your webhook using the server component as you would with using ngrok or other alternatives.

To use the command you must run the following command replacing the arguments with your own values:

```shell
hooktest ctest <domain: wss://example.com/ws> <webhook-port-number>
```

#### 3. run

This command simply runs the last executed command again.

To use it simply run the following command:

```shell
hooktest run
```