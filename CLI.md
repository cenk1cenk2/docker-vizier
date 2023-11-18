# vizier

Supervisor for running multiple tasks in a Docker container.

`vizier [FLAGS]`

## Flags

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
