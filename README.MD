# Golang Pokemon API Concurrency Fetcher

This Golang application demonstrates the use of Go's concurrency features, including goroutines, channels, and synchronization, to concurrently fetch and print Pokemon data from the [PokeAPI](https://pokeapi.co/api/v2/pokemon).

## Overview

The application retrieves a list of Pokemon names and URLs from the PokeAPI, then concurrently fetches additional details (weight, base experience) for each Pokemon using goroutines. The fetched data is printed to the console as it becomes available.

## Dependencies

The application relies on the standard library of Go and has no external dependencies.

## How to Use

1. Clone the repository:

    ```bash
    git clone https://github.com/iliatalebzade/concurrency-in-go.git
    cd concurrency-in-go
    ```

2. Run the application:

    ```bash
    go run main.go
    ```

## Code Structure

- `main.go`: Contains the main application logic.
- `README.md`: Documentation for the application.

## Concurrency Features

The application leverages the following concurrency features:

- **Goroutines**: Utilized to concurrently fetch Pokemon details.
- **Channels**: Used for communication between goroutines, passing Pokemon names.
- **WaitGroup**: Ensures that all goroutines finish before closing the source channel.

## Code Explanation

- `fetchAndPrintPokemonData`: A function that fetches additional Pokemon details and sends the name to a channel.
- `main`: The main function where the PokeAPI data is retrieved, goroutines are spawned, and the results are printed.

## License

This project is licensed under the [MIT License](LICENSE).

## Contribution

Feel free to contribute to this project by opening issues or pull requests.
