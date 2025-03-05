# Enigma WASM

A WebAssembly implementation of the Enigma encryption machine compiled from Go, allowing secure encryption directly in the browser.

## Overview

This project brings the historical Enigma encryption machine to the modern web through WebAssembly. It provides a fully functional Enigma machine implementation that runs entirely in the browser, with no server-side processing required. The encryption and decryption processes happen locally, ensuring privacy and security.

The implementation features configurable rotors, cross-connections, and reflectors, allowing for millions of possible encryption combinations.

## Features

- **Complete Enigma Implementation**: Accurate simulation of the Enigma machine mechanics
- **WebAssembly Performance**: Compiled from Go to WebAssembly for near-native performance
- **Client-Side Processing**: All encryption happens in the browser
- **Configurable Components**:
  - Multiple rotors with customizable sequences
  - Custom cross-connections
  - Reflector configuration
- **Extended Character Set**: Supports uppercase letters, numbers, and special characters
- **Two-Way Encryption**: The same process encrypts and decrypts messages
- **Configuration Export/Import**: Save and load custom Enigma configurations

## How to Use

### Running the Application

1. Clone the repository:
   ```bash
   git clone https://github.com/rohandhamapurkar/enigma-wasm.git
   cd enigma-wasm
   ```

2. Serve the files using a basic HTTP server (required for WebAssembly):
   ```bash
   # Using Python
   python -m http.server

   # Using Node.js (with http-server package)
   npx http-server
   ```

3. Open http://localhost:8000 (or http://localhost:8080 for Go) in your browser

### Building from Source

To rebuild the WebAssembly binary:

```bash
GOOS=js GOARCH=wasm go build -o main.wasm ./cmd
```

### Using the Enigma Machine

1. Enter text in the input field
2. The text is automatically encrypted/decrypted using the current Enigma configuration
3. Use the "Get Default Config" button to download a random Enigma configuration
4. Use "Import Config" to load a previously saved configuration

## How It Works

The Enigma machine operates on a principle of substitution cipher with a twist:

1. **Rotor System**: When a key is pressed, an electrical signal passes through a series of rotors
2. **Reflector**: After passing through the rotors, the signal hits a reflector and returns through the rotors in reverse
3. **Rotation**: After each keystroke, one or more rotors rotate, changing the encryption pattern
4. **Symmetry**: The encryption is symmetrical - using the same settings to encrypt a message will decrypt it when the original message is fed back through

Our implementation includes:
- **Rotors**: Each with a unique scrambled sequence of characters and cross-connections
- **Reflector**: Maps each character to another, ensuring the encryption is reversible
- **Rotation Logic**: Simulates the mechanical rotation of rotors, including cascade rotation

## Technical Details

### Architecture

The project is organized into several packages:

- **cmd**: Contains the main.go entry point and JavaScript interface registration
- **config**: Handles Enigma configuration management and generation
- **machine**: Implements the core Enigma encryption logic
- **helpers**: Provides utility functions for character mapping, rotation, etc.
- **constants**: Defines the character set used for encryption

The root directory contains:
- **index.html**: The web interface for the application
- **wasm_exec.js**: Standard Go WebAssembly runtime
- **wasm_exec_tiny.js**: Lightweight alternative WebAssembly runtime

### WebAssembly Integration

The Go code is compiled to WebAssembly and integrated with JavaScript through the following:

- **wasm_exec.js**: The Go WebAssembly runtime
- **JavaScript Bindings**: Functions exported from Go and available in JavaScript:
  - `getDefaultEnigmaConfigStream`: Returns a random Enigma configuration
  - `getCurrentConfigStream`: Returns the current machine configuration
  - `setEnigmaConfig`: Loads a configuration into the machine
  - `ScrambleText`: Encrypts/decrypts text using the current configuration

## License

```
MIT License

Copyright (c) 2025 Rohan Dhamapurkar

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## Historical Context

The original Enigma machine was used for encrypting military communications during World War II, particularly by Nazi Germany. It was considered unbreakable until British mathematicians, including Alan Turing, developed techniques to decrypt messages at Bletchley Park. This work significantly contributed to the Allied victory and laid foundations for modern computing.

This project reimagines the Enigma machine with modern technology while maintaining the core principles that made it such a fascinating encryption device.

## Browser Compatibility

This application works in all modern browsers that support WebAssembly:

- Chrome 57+
- Firefox 53+
- Safari 11+
- Edge 16+

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
