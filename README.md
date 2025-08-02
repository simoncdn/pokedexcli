# Pokedex CLI

A command-line Pokedex application built in Go that allows you to explore Pokemon locations, catch Pokemon, and manage your collection using the PokeAPI.

## Features

- **Explore Locations**: Browse through different Pokemon locations
- **Catch Pokemon**: Attempt to catch wild Pokemon found in locations
- **Pokedex Management**: View and inspect your caught Pokemon collection
- **Caching**: Built-in caching system for improved performance
- **Interactive REPL**: User-friendly command-line interface

## Installation

1. Clone the repository:
```bash
git clone https://github.com/simoncdn/pokedexcli.git
cd pokedexcli
```

2. Build the application:
```bash
go build
```

3. Run the application:
```bash
./pokedexcli
```

## Usage

Once started, the application enters an interactive mode with the prompt `Pokedex > `. You can use the following commands:

### Available Commands

- **`help`** - Displays this help message with all available commands
- **`map`** - Get the next page of locations to explore
- **`mapb`** - Get the previous page of locations
- **`explore <location_name>`** - Explore a specific location to find Pokemon
- **`catch <pokemon_name>`** - Attempt to catch a Pokemon (must be in current location)
- **`inspect <pokemon_name>`** - View detailed information about a caught Pokemon
- **`pokedex`** - Show all Pokemon in your collection
- **`exit`** - Exit the application

### Example Session

```
Pokedex > help
Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Get the next page of locations
mapb: Get the previous page of locations
explore <location_name>: Explore a location
catch <pokemon_name>: Try to catch a Pokemon
inspect <pokemon_name>: View details about a caught Pokemon
pokedex: Show all of your Pokemon in your Pokedex
exit: Exit the Pokedex

Pokedex > map
# Lists available locations

Pokedex > explore canalave-city-area
# Shows Pokemon available in that location

Pokedex > catch pikachu
# Attempts to catch Pikachu

Pokedex > pokedex
# Shows your caught Pokemon

Pokedex > inspect pikachu
# Shows detailed stats for Pikachu

Pokedex > exit
```

## Project Structure

```
pokedexcli/
├── main.go                 # Application entry point
├── repl.go                 # REPL implementation and command handling
├── command_*.go            # Individual command implementations
├── internal/
│   ├── pokeapi/           # PokeAPI client and data structures
│   │   ├── client.go      # HTTP client for PokeAPI
│   │   ├── location_*.go  # Location-related API calls
│   │   ├── pokemon_*.go   # Pokemon-related API calls
│   │   └── types_*.go     # Data type definitions
│   └── pokecache/         # Caching system
│       ├── cache.go       # Cache implementation
│       └── pokecache.go   # Cache interface
└── README.md
```

## Technical Details

- **Language**: Go 1.24.2
- **API**: Uses the [PokeAPI](https://pokeapi.co/) for Pokemon data
- **Caching**: Implements a time-based cache with 5-minute expiration
- **Client Timeout**: 5-second timeout for API requests

## Dependencies

This project uses only Go's standard library with no external dependencies.

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o pokedexcli
```
