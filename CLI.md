# vizier

Supervisor for running multiple tasks in a Docker container.

`vizier [GLOBAL FLAGS] [COMMAND] [FLAGS]`

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
