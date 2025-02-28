# Runtime-Generator

[![Build](https://github.com/CraftUniverse/CraftEngine-Runtime-Generator/actions/workflows/build.yml/badge.svg)](https://github.com/CraftUniverse/CraftEngine-Runtime-Generator/actions/workflows/build.yml)

The Runtime Generator generates dummy Methods & Classes for any **Supported Language**, defined by any **Supported Language**.

### Table of Contents

- [Run](#run)
- [CLI](#cli)
- [License](#license)
- [Contributers](#contributers)

## Run

```bash
go run cmd/generator/main.go <arguments>
```

## CLI

```
generator
    --input <string>
        The language to convert from (server = false)

    --src <string>
        The input project path (server = false)

    --output <string>
        The output package path (server = false)

    --port <int>
        The port to start the server on (server = true)

    --server <boolean>
        If enabled a websocket will be started and files can be converted on the fly
```

## License

[CraftEngine](https://craftengine.dev) & [CraftEngine Runtime Generator](https://github.com/CraftUniverse/CraftEngine-Runtime-Generator) © 2025 by [CraftUniverse](https://craftuniverse.net) is licensed under [Creative Commons Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/?ref=chooser-v1)

## Contributers

- [@Turboman3000](https://github.com/Turboman3000)
