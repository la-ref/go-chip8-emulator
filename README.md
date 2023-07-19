
# **CHIP 8 EMULATOR**

![Chip8 Emulator](chip8_emulator_screenshot.png)

A Chip8 emulator written in Golang, using the SDL library with Go SDL for the graphical interface.


> Author : CZARKOWSKI Matthieu


## **Intoduction**

Chip-8 is an interpreted programming language, developed by Joseph Weisbecker. It was initially used on the COSMAC VIP and Telmac 1800 8-bit microcomputers. 

The Chip8 platform featured a 4KB memory and was designed to create simple games and applications. While Chip8 programs were originally intended for microcomputers, enthusiasts later developed emulators to run these vintage programs on modern systems (like me).

## **The project 📢**

The project is to practice a discover the world of emulation and the SDL library, so my first approach was to create a Chip8 emulator with Go and Go-SDL2 bindings

The primary goal is to provide a stable and feature-rich emulator that can faithfully recreate the behavior of original Chip8 systems. 

SDL provide a graphical user interface and handle user input.

### Features
- Accurate Chip8 emulation
- Graphical display rendering (64x32 for the orignal chip8) scaled
- Audio playback (if applicable to the ROM)
- Keyboard input handling (check Key Bindings)
- Save and load state functionality (if applicable in future updates)


## **Requirements 📍**

### Language 📙

This project is written in the Go programming language (Golang). Go is known for its simplicity, efficiency, and excellent concurrency support, making it a great choice for developing high-performance applications and emulators like this Chip8 project.

Download : https://go.dev/doc/install

### Libraries 📚

#### 1. SDL2 (Go-SDL2)

**Go-SDL2** a library that wrap SDL2 for Go.
SDL is multimedia development library designed to provide low-level access to audio, keyboard, mouse, joystick, and graphics hardware via OpenGL and DirectX

```
go get github.com/veandco/go-sdl2/sdl
```

Note: SDL2 library is written in C. That means that you need the original SDL2 installated. Read the installation instructions for sdl2 library from [`go-sdl2`](github.com/veandco/go-sdl2) for your os.

## **How to use the application ❓**

### Launch the program ▶️

#### 1. Download the project:

Clone this repository to your local machine:

```bash
git clone https://github.com/la-ref/go-chip8-emulator
cd go-chip8-emulator
```
OR

Get code :
```
go get -u github.com/skatiyar/go-chip8
```


2   . Load a Chip8 ROM:
```
Usage: go run main.go <path_to_rom_file>
```
3   . Build and run the emulator:
```
go run main.go
```

### Key Bindings 🎲

```
    Chip8           Keyboard (AZERTY)       Keyboard (QWERTY)
1 | 2 | 3 | C        1 | 2 | 3 | 4           1 | 2 | 3 | 4
- - - - - - -        - - - - - - -           - - - - - - - 
4 | 5 | 6 | D        A | Z | E | R           Q | W | E | R
- - - - - - -        - - - - - - -           - - - - - - - 
7 | 8 | 9 | E        Q | S | D | F           A | S | D | F
- - - - - - -        - - - - - - -           - - - - - - - 
A | 0 | B | F        W | X | C | V           Z | X | C | V
```
The default keyboard mapping is assigned to AZERTY

## **Sources**

- [Guide to making a CHIP-8 emulator](https://tobiasvl.github.io/blog/write-a-chip-8-emulator/)
- [Chip-8 opcode table](https://en.wikipedia.org/wiki/CHIP-8)

## **Support Me**
Give me a ⭐ if this project was helpful in any way!

## **Licence**
This project is licensed under the MIT License. You are free to use, modify, and distribute this code following the terms of the license
