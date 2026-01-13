# vizier

Supervisor for running multiple tasks in a Docker container.

`vizier [FLAGS]`

## Flags

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
