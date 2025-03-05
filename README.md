# Bitcoin Handshake Project

## Overview

This project is a Bitcoin handshake implementation written in Go. It establishes a connection with Bitcoin nodes using TCP transport and performs a handshake to exchange version information.

### Prerequisites

- Go programming language installed on your system.
- Access to the internet to connect to Bitcoin nodes.
- To run BTCD node:
    ```bash
    git clone https://github.com/btcsuite/btcd.git
    ```

    ```bash
    go install -v . ./cmd/...
    ```
    - btcd (and utilities) will now be installed in $GOPATH/bin

    ```bash
    ./btcd
    ```

### Installation

1. Clone the repository to your local machine:

    ```bash
    git clone <repository_url>
    ```

2. Change to the project directory:

    ```bash
    cd bitcoin-handshake
    ```

### Running the Project

1. Ensure you have the necessary dependencies installed:

    ```bash
    go mod download
    ```

2. Run the project:

    ```bash
    go run main.go
    ```

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please submit an issue or a pull request.
