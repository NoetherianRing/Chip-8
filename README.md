## Chip-8

This is a Chip-8 emulator.   

Chip-8 is an interpreted programming language, developed by Joseph Weisbecker. It was initially used on the COSMAC VIP and Telmac 1800 8-bit microcomputers.  

[Here's](https://en.wikipedia.org/wiki/CHIP-8) there is link to its wikipedia page, and [here's](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM) a link to Cowgod's Chip-8 Technical Reference.

## Requirements

This Chip-8 emulator uses [PixelGL](https://github.com/faiface/pixel/blob/master/README.md) and PixelGL uses OpenGL to render graphics. Because of that, OpenGL development libraries are needed for compilation. The dependencies are same as for [GLFW](https://github.com/go-gl/glfw).

The OpenGL version used is **OpenGL 3.3**.

- On macOS, you need Xcode or Command Line Tools for Xcode (`xcode-select --install`) for required headers and libraries.

- On Ubuntu/Debian-like Linux distributions, you need `libgl1-mesa-dev` and `xorg-dev` packages.

- On CentOS/Fedora-like Linux distributions, you need `libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel libXxf86vm-devel` packages.

- See [here](http://www.glfw.org/docs/latest/compile.html#compile_deps) for full details.

  

## Installation

Chip8 executable for linux is provided in the [release](https://github.com/NoetherianRing/Chip-8/releases/tag/release) section.

## Run

By default, it's going to run with the ROM file PONG.ch8. If you want to change it, you need to change the configuration.

## Configuration

This Chip-8 emulator has a config.yml file that looks like this:

```yml
paths:
  beep: "../Chip-8/assets/beep.mp3"
  rom: "../Chip-8/assets/PONG.ch8"
  fonts: "../Chip-8/assets/chip8.font"

debug:
  on: "false"
  file: "DEBUG.json"

test:
  expectedStateROM1: "../fixtures/PONG.json"
  expectedStateROM2: "../fixtures/IBM_Logo.json"
  expectedStateROM3: "../fixtures/chip8_logo.json"
  ROM1: "../assets/PONG.ch8"
  ROM2: "../assets/IBM_Logo.ch8"
  ROM3: "../assets/chip8_logo.ch8"
  FONT: "../assets/chip8.font"




```

#### ROM Files

By default it's going to execute a Pong game.  To change to another you can modify the line

```yml
 rom: "../Chip-8/assets/PONG.ch8"
```

to a different relative root.

#### Debug mode

The debug mode runs the chip8 taking the state of the chip in every cycle and saving it into a json file whith the name specify in the field "file".

It can be activated modifying the config.yml file this way:

```yml
debug:
  on: "true"
  file: "DEBUG.json"
   
```

#### Test 

The tests of this Chip-8 emulator compares the state of the chip with a desired state for certain ROM files specified in the "test" section.

It should not to be modified.
```yml
test:
  expectedStateROM1: "../fixtures/PONG.json"
  expectedStateROM2: "../fixtures/IBM_Logo.json"
  expectedStateROM3: "../fixtures/chip8_logo.json"
  ROM1: "../assets/PONG.ch8"
  ROM2: "../assets/IBM_Logo.ch8"
  ROM3: "../assets/chip8_logo.ch8"
  FONT: "../assets/chip8.font"

```

 