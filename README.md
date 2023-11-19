# cenk1cenk2/grand-vizier

[![pipeline status](https://gitlab.kilic.dev/docker/vizier/badges/master/pipeline.svg)](https://gitlab.kilic.dev/docker/vizier/-/commits/master) [![Docker Pulls](https://img.shields.io/docker/pulls/cenk1cenk2/vizier)](https://hub.docker.com/repository/docker/cenk1cenk2/vizier) [![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/cenk1cenk2/vizier)](https://hub.docker.com/repository/docker/cenk1cenk2/vizier) [![Docker Image Version (latest by date)](https://img.shields.io/docker/v/cenk1cenk2/vizier)](https://hub.docker.com/repository/docker/cenk1cenk2/vizier) [![GitHub last commit](https://img.shields.io/github/last-commit/cenk1cenk2/vizier)](https://github.com/cenk1cenk2/vizier)

## Description

some description

---

- [CLI Documentation](./CLI.md)

<!-- toc -->

<!-- tocstop -->

---

<!-- clidocs -->

## Global Flags

| Flag / Environment |  Description   |  Type    | Required | Default |
|---------------- | --------------- | --------------- |  --------------- |  --------------- |
| `$VIZIER_STEPS` | Steps to run for the application, will be ignored when configuration file is read. | `String`<br/>`json(https://raw.githubusercontent.com/cenk1cenk2/docker-vizier/main/schema.json)` | `false` |  |

### CLI

| Flag / Environment |  Description   |  Type    | Required | Default |
|---------------- | --------------- | --------------- |  --------------- |  --------------- |
| `$LOG_LEVEL` | Define the log level for the application. | `String`<br/>`enum("panic", "fatal", "warn", "info", "debug", "trace")` | `false` | info |
| `$ENV_FILE` | Environment files to inject. | `StringSlice` | `false` |  |

### Config

| Flag / Environment |  Description   |  Type    | Required | Default |
|---------------- | --------------- | --------------- |  --------------- |  --------------- |
| `$VIZIER_CONFIG` | Configuration file to read from. | `String`<br/>`json(https://raw.githubusercontent.com/cenk1cenk2/docker-vizier/main/schema.json)` | `false` |  |

## Commands

### `generate`

Generate json schema

`vizier generate [GLOBAL FLAGS] [FLAGS]`

#### Flags

| Flag / Environment |  Description   |  Type    | Required | Default |
|---------------- | --------------- | --------------- |  --------------- |  --------------- |
| `$VIZIER_SCHEMA_OUTPUT` | Schema file to write to. | `String` | `false` | schema.json |

<!-- clidocsstop -->
