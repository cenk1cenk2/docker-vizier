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
| `$VIZIER_STEPS` | Steps to run for the application, will be ignored when configuration file is read. json([]struct {<br />  name?: string<br />  commands?: []struct {<br />    cwd?: string<br />    command: string<br />    retry?: struct {<br />      retries?: number<br />      always?: boolean<br />      delay?: string<br />    }<br />    ignore_error?: boolean<br />    log?: struct {<br />      stdout?: VizierLogLevels<br />      stderr?: VizierLogLevels<br />      lifetime?: VizierLogLevels<br />    }<br />    environment?: map[string]string<br />    run_as?: struct {<br />      user?: string<br />      group?: string<br />    }<br />  }<br />  permissions?: []struct {<br />    path: string<br />    chown?: struct {<br />      user?: string<br />      group?: string<br />    }<br />    chmod?: struct {<br />      file?: string<br />      dir?: string<br />    }<br />    recursive?: boolean<br />  }<br />  delay?: string<br />  background?: boolean<br />  parallel?: boolean<br />}) | `String` | `false` |  |

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
