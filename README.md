# movemouse

A lightweight utility that continuously moves your mouse cursor in random patterns. Useful for preventing your computer from entering sleep mode, screensaver activation, or idle timeouts.

## What It Does

The app runs an infinite loop that:
1. Moves the mouse cursor randomly (0-300 pixels in both X and Y directions)
2. Waits for 1-10 seconds randomly
3. Moves the mouse back to its original position
4. Waits again for 1-10 seconds randomly
5. Repeats indefinitely

This simulates light mouse activity without interfering with normal usage.

## Usage

### Using the Pre-built Binary

A pre-compiled binary for macOS (ARM64) is included:

```bash
./bin/movemouse-macos-arm64
```

The app will start moving your mouse and print movement details to the console. Press `Ctrl+C` to stop it.

### Building from Source

**Requirements:**
- Go 1.26.2 or later
- macOS, Linux, or Windows

**Build:**

```bash
make build
```

Or manually with Go:

```bash
go build -o movemouse main.go
./movemouse
```

## Installation

For convenience, you can copy the binary to a system location:

```bash
sudo cp bin/movemouse-macos-arm64 /usr/local/bin/movemouse
movemouse
```

## Use Cases

- Prevent screensaver activation during presentations
- Keep system awake during downloads or long-running operations
- Prevent automatic idle timeout on work machines
- Testing mouse input behavior in applications
