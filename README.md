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

| Flag / Environment |  Description   |  Type    | Required | Default |
|---------------- | --------------- | --------------- |  --------------- |  --------------- |
| `$VIZIER_CONFIG` | Steps to run for the application, will be ignored when configuration file is read. | `string`<br/>`json(https://raw.githubusercontent.com/cenk1cenk2/docker-vizier/main/schema.json)` | `false` | <code></code> |

**CLI**

| Flag / Environment |  Description   |  Type    | Required | Default |
|---------------- | --------------- | --------------- |  --------------- |  --------------- |
| `$LOG_LEVEL` | Define the log level for the application. | `string`<br/>`enum("panic", "fatal", "warn", "info", "debug", "trace")` | `false` | <code>"info"</code> |
| `$ENV_FILE` | Environment files to inject. | `string[]` | `false` | <code></code> |

**Config**

| Flag / Environment |  Description   |  Type    | Required | Default |
|---------------- | --------------- | --------------- |  --------------- |  --------------- |
| `$VIZIER_CONFIG_FILE` | Configuration file to read from. | `string`<br/>`json(https://raw.githubusercontent.com/cenk1cenk2/docker-vizier/main/schema.json)` | `false` | <code></code> |

<!-- clidocsstop -->
