## Chip-8

This is a Chip-8 emulator.   

Chip-8 is an interpreted programming language, developed by Joseph Weisbecker. It was initially used on the COSMAC VIP and Telmac 1800 8-bit microcomputers.  

[Here's](https://en.wikipedia.org/wiki/CHIP-8) there is link to its wikipedia page, and [here's](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM) a link to Cowgod's Chip-8 Technical Reference.

This implementation uses faiface [pixel](https://github.com/faiface/pixel) for graphics and [beep](https://github.com/faiface/beep) for sound. 



##Additional Memory and Instructions
This emulator supports more memory than the original Chip 8 system, allowing you to run games and programs that require more memory. In addition, we have added two more instructions to the original set to enhance the compatibility with ROMs files generated with the "c8-compiler".

The two new instructions are:

- `9XY1`: Save the value of register `VX` in the first 8 bits of the index register `I`, and the value of register `VY` in the last 8 bits of I.
- `9XY2`: Save the first 8 bits of the value of index register `I` in register `VX`, and the last 8 bits of `I` in register `VY`.
  These instructions are not present in the original Chip 8 instruction set but are required to run ROMs files generated with the "c8-compiler".

To learn more about the "c8-compiler", please visit the [c8-compiler Git repository](https://github.com/NoetherianRing/c8-compiler).


## Requirements

This Chip-8 emulator uses [PixelGL](https://github.com/faiface/pixel/blob/master/README.md) and PixelGL uses OpenGL to render graphics. Because of that, OpenGL development libraries are needed for compilation. The dependencies are same as for [GLFW](https://github.com/go-gl/glfw).

The OpenGL version used is **OpenGL 3.3**.

- On macOS, you need Xcode or Command Line Tools for Xcode (`xcode-select --install`) for required headers and libraries.

- On Ubuntu/Debian-like Linux distributions, you need `libgl1-mesa-dev` and `xorg-dev` packages.

- On CentOS/Fedora-like Linux distributions, you need `libX11-dev9el libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel libXxf86vm-devel` packages.

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

By default it's going to execute a Pong game.  To change it to another you can modify the line

```yml
 rom: "../Chip-8/assets/PONG.ch8"
```

to a different relative root.

#### Debug mode

The debug mode runs the chip8 taking the state of the chip in every cycle and saving it into a json file with the name specified in the field "file".
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
It also can be tested by running the chip8-test-suite.ch8 ROM in the assets folder, it was taken from [Timendus](https://github.com/Timendus/chip8-test-suite) github. 

## Keys

To quit the app you need to press the key Esc. 

Here is the list of the keys mapping for the Chip 8:


|        KEYPAD          |       KEYBOARD      |
| :--------------------: | :-----------------: |
|        1 2 3 C         |       1 2 3 4       |
|        4 5 6 D         |       Q W E R       |
|        7 8 9 E         |       A S D F       |
|        A 0 B F         |       Z X C V       |

