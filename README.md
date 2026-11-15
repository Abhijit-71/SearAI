# SearAI

Open metasearch engine for AI agents to retrive web results , written in Go which Aggregates results from multiple providers.


## Quick start

```bash
# Run
go run .
```

## Build

```bash
# build binary
go build -o searAI
./searAI
```

## Configure
Edit `config.toml`.

Key options:
- `server.default_size`: results per page
- `engines.*.enabled`: toggle engines
- `engines.*.priority`: ranking order (lower = higher priority)


## Requirements
- Go 1.20+ (recommended)


## Acknowledgements
- Its a fork of [SearxGo](https://github.com/Openmentia/SearxGo)

## Contributions
- Contributions are welcomed ... fork the repo run experiments , make a pr 
- Have good code and documentation practices
- All PRs should clearly state what it is meant for.
- Commits with code removal ; it is expected to have a more explicit documentation
- For contributions checkout [Roadmap](/ROADMAP.md) and help us with our future goals
