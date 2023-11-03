# cenk1cenk2/grand-vizier

[![pipeline status](https://gitlab.kilic.dev/docker/softether-vpnsrv/badges/master/pipeline.svg)](https://gitlab.kilic.dev/docker/softether-vpnsrv/-/commits/master) [![Docker Pulls](https://img.shields.io/docker/pulls/cenk1cenk2/softether-vpnsrv)](https://hub.docker.com/repository/docker/cenk1cenk2/softether-vpnsrv) [![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/cenk1cenk2/softether-vpnsrv)](https://hub.docker.com/repository/docker/cenk1cenk2/softether-vpnsrv) [![Docker Image Version (latest by date)](https://img.shields.io/docker/v/cenk1cenk2/softether-vpnsrv)](https://hub.docker.com/repository/docker/cenk1cenk2/softether-vpnsrv) [![GitHub last commit](https://img.shields.io/github/last-commit/cenk1cenk2/softether-vpnsrv)](https://github.com/cenk1cenk2/softether-vpnsrv)

## Description

some description

---

- [CLI Documentation](./CLI.md)

<!-- toc -->

<!-- tocstop -->

---

<!-- clidocs -->

| Flag / Environment |  Description   |  Type    | Required | Default |
|---------------- | --------------- | --------------- |  --------------- |  --------------- |
| `$VIZIER_STEPS` | Steps to run for the application, will be ignored when configuration file is read. | `String`<br/>`json([]struct{})` | `false` |  |

### CLI

| Flag / Environment |  Description   |  Type    | Required | Default |
|---------------- | --------------- | --------------- |  --------------- |  --------------- |
| `$LOG_LEVEL` | Define the log level for the application. | `String`<br/>`enum("panic", "fatal", "warn", "info", "debug", "trace")` | `false` | info |
| `$ENV_FILE` | Environment files to inject. | `StringSlice` | `false` |  |

### Config

| Flag / Environment |  Description   |  Type    | Required | Default |
|---------------- | --------------- | --------------- |  --------------- |  --------------- |
| `$VIZIER_CONFIG` | Configuration file to read from. | `String` | `false` |  |

<!-- clidocsstop -->
