# Gator - RSS Feed Aggregator

A command-line RSS feed aggregator built with Go and PostgreSQL.

## Prerequisites

- **Go** (1.19 or higher) - [Install Go](https://golang.org/doc/install)
- **PostgreSQL** - [Install PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install the `gator` CLI tool using `go install`:

```bash
go install github.com/ianivr/gator@latest
```

This will compile and install the `gator` executable to your `$GOPATH/bin` directory.

## Configuration

Create a config file at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator",
  "current_user": "your_username"
}
```

## Running Gator

Start the program:

```bash
gator
```

### Available Commands

- `register <username>` - Create a new user
- `login <username>` - Switch to a different user
- `addfeed <name> <url>` - Subscribe to an RSS feed
- `feeds` - List all available feeds
- `follow <feed_url>` - Follow a specific feed
- `following` - Show feeds you're following
- `browse [limit]` - View recent posts (default: 2 posts)
- `agg <duration>` - Start aggregating feeds (e.g., `agg 1m` for every minute)
