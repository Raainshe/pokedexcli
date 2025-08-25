# Pokédex CLI

A command-line Pokédex application built in Go that allows you to explore the Pokémon world, catch Pokémon, and manage your collection. This project demonstrates various Go programming concepts and best practices.

## 🚀 Features

- **Interactive REPL**: Command-line interface with a Read-Eval-Print Loop
- **Map Exploration**: Navigate through different locations in the Pokémon world
- **Pokémon Discovery**: Explore areas to find wild Pokémon
- **Catch Mechanics**: Attempt to catch Pokémon with success rates based on base experience
- **Personal Pokédex**: Store and inspect your caught Pokémon
- **Smart Caching**: HTTP response caching with automatic cleanup to improve performance
- **Comprehensive Testing**: Unit tests for core functionality

## 🏗️ Architecture

```
pokedexcli/
├── main.go              # Entry point and REPL implementation
├── commands.go          # Command handlers and API interactions
├── internal/
│   └── pokecache.go     # HTTP response caching system
├── repl_test.go         # Unit tests
├── go.mod               # Go module definition
└── README.md           # This file
```

## 🛠️ Build and Run

### Prerequisites
- Go 1.24.4 or later

### Building the Application
```bash
# Clone the repository
git clone https://github.com/Raainshe/pokedexcli.git
cd pokedexcli

# Build the application
go build -o pokedexcli

# Run the application
./pokedexcli
```

### Alternative: Run without building
```bash
go run .
```

### Running Tests
```bash
go test -v
```

## 🎮 Usage

Once the application starts, you'll see the `Pokedex >` prompt. Available commands:

| Command | Description | Example |
|---------|-------------|---------|
| `help` | Display all available commands | `help` |
| `exit` | Exit the Pokedex | `exit` |
| `map` | Show the next 20 locations | `map` |
| `mapb` | Show the previous 20 locations | `mapb` |
| `explore <area>` | List Pokémon in a specific area | `explore canalave-city-area` |
| `catch <pokemon>` | Attempt to catch a Pokémon | `catch pikachu` |
| `inspect <pokemon>` | View details of a caught Pokémon | `inspect pikachu` |
| `pokedex` | List all your caught Pokémon | `pokedex` |

### Example Session
```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - staryu
 - magikarp

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  -hp: 20
  -attack: 10
  -defense: 55
  -special-attack: 15
  -special-defense: 20
  -speed: 80
Types:
  - water
```

## 📚 Go Concepts Learned

This project demonstrates numerous Go programming concepts:

### 1. **HTTP Client and API Integration**
- Using `net/http` package to make HTTP requests
- Handling HTTP responses and error codes
- Working with REST APIs (PokéAPI)
- Reading response bodies with `io.ReadAll`

### 2. **JSON Handling**
- Struct tags for JSON marshaling/unmarshaling
- Working with nested JSON structures
- Custom data transformation from API responses

### 3. **Caching System**
- Custom in-memory cache implementation
- Thread-safe operations using `sync.Mutex`
- Automatic cache cleanup with goroutines and tickers
- Time-based cache expiration

### 4. **Concurrency**
- Goroutines for background tasks (cache cleanup)
- Mutex for thread-safe access to shared data
- Channels and `select` statements for timer-based operations

### 5. **Package Organization**
- Internal packages for encapsulation
- Proper module structure with `go.mod`
- Package imports and dependency management

### 6. **Command Pattern**
- Function types and callbacks
- Map-based command dispatch
- Clean separation of command logic

### 7. **Error Handling**
- Idiomatic Go error handling with explicit error returns
- Error wrapping with `fmt.Errorf` and `%w` verb
- Graceful error recovery and user feedback

### 8. **Input Processing**
- String manipulation and cleaning
- Command-line argument parsing
- User input validation

### 9. **Testing**
- Unit tests with the `testing` package
- Table-driven tests for multiple test cases
- Testing concurrent code and time-based functionality

### 10. **REPL Implementation**
- Interactive command-line interface
- Input scanning with `bufio.Scanner`
- Command parsing and execution loop

### 11. **Struct Methods and Receivers**
- Methods on custom types
- Pointer vs value receivers
- Data encapsulation and behavior

### 12. **Random Number Generation**
- Using `math/rand` for game mechanics
- Probability-based catch system

### 13. **Memory Management**
- Efficient data structures
- Proper resource cleanup (defer statements)
- Cache size management

## 🔧 Technical Implementation Details

### Caching Strategy
The application implements a sophisticated caching system that:
- Stores HTTP responses in memory to reduce API calls
- Uses mutex locks for thread-safe operations
- Automatically cleans up expired entries every 10 minutes
- Significantly improves performance for repeated requests

### Catch Mechanics
Pokémon catch rates are determined by base experience:
- ≤ 80 base experience: 80% catch rate
- ≤ 150 base experience: 60% catch rate
- ≤ 250 base experience: 40% catch rate
- \> 250 base experience: 20% catch rate

### Data Structures
The application uses several custom structs to model Pokémon data:
- `Pokemon`: Complete Pokémon information with stats and types
- `mapsList`: API response structure for location data
- `LocationAreaDetail`: Pokémon encounters in specific areas
- `Cache`: Thread-safe caching system

## 🧪 Testing

The project includes comprehensive tests covering:
- Input cleaning and validation (`TestCleanInput`)
- Cache add/get operations (`TestAddGet`)
- Cache expiration and cleanup (`TestReapLoop`)

Run tests with verbose output:
```bash
go test -v
```

## 🤝 Contributing

This project is part of a Go learning journey. Feel free to:
- Report bugs or issues
- Suggest improvements
- Submit pull requests
- Share your own learning experiences


## 🙏 Acknowledgments

- [PokéAPI](https://pokeapi.co/) for providing the Pokémon data
- The Go community for excellent documentation and resources
- Boot.dev for the Go learning curriculum that inspired this project

## 📧 Author

**@Raainshe** - [GitHub Profile](https://github.com/Raainshe)

---

*Built with ❤️ while learning Go*
