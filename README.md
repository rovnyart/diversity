# Diversity - Windows wallpaper changer written on Go

## Unsplash API token

First of all you need to get an API token from Unsplash. Go to `https://unsplash.com/developers`, create an account, create an app and get API Key.

## Installation

You can download pre-built Windows amd64 binary from the [releases](https://github.com/rovnyart/diversity/releases) section.

Alternatively, you can build an app on your own:

```powershell
go get github.com/rovnyart/diversity

$env: GOARCH='windows'
$env: GOTARGET='amd64'

go build github/rovnyart/diversity
```

This will produce `diversity.exe` executable. In order to run it you need to place `config.yaml` file in the same directory. See `Config` section for more details.

## Config

Config file has a very simple structure:

```yaml
---
apikey: your-api-token-here # Your Unsplash API token
scope: nature # Picture search query
schedule: 2h # Automatic wallpaper change schedule
```
