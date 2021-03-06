# payterm
payment terminal toy project

## Used Hardware and connection
Price of all hardware is less than 10$ + Raspberry Pi Zero W.
### Controller - Raspberry Pi Zero W

![Raspberry Pi Zero W](docs/rpi0w.png)

 - Pinout https://pinout.xyz/
 - numbers in `( )` in *Connection* tables are physical pins.

### Display - 1602A v2.0

![1602A Display](docs/display.png)

 - 16x2 characters
 - +some 2 random pots I've found at home to control contrast (CPot) and backlight (LPot)
 - useful link: https://pimylifeup.com/raspberry-pi-lcd-16x2/

#### Connection

| No   | Name  | Desc                   | To          |
| -    | -     | -                      | -           |
| 1    | VSS   | -Power                 | gnd (39)    |
| 2    | VDD   | +Power                 | 5V (2)      |
| 3    | V0    | Contrast               | CPot        |
| 4    | RS    | Register select        | GPIO26 (37) |
| 5    | R/W   | R/W switch (W=0; R=1)  | gnd (39)    |
| 6    | E     | GND                    | GPIO19 (35) |
| 7-10 | DB0-3 | Data                   | -           |
| 11   | DB4   | Data                   | GPIO13 (33) |
| 12   | DB5   | Data                   | GPIO21 (40) |
| 13   | DB6   | Data                   | GPIO20 (38) |
| 14   | DB7   | Data                   | GPIO16 (36) |
| 15   | A     | +Light                 | LPot        |
| 16   | K     | -Light                 | gnd (39)    |

### Card Reader - RFID-RC522

![Card Reader](docs/cardreader.png)

 - support cards - S50, S70, UltraLight, Pro, Desfire
 - +some blank S50 cards
 - useful link: https://pimylifeup.com/raspberry-pi-rfid-rc522/

#### Connection

| No   | Name | Desc                | To          |
| -    | -    | -                   | -           |
| 1    | SDA  | Serial Data Signal  | GPIO8 (24)  |
| 2    | SCK  | Serial Clock        | GPIO11 (23) |
| 3    | MOSI | Master Out Slave In | GPIO10 (19) |
| 4    | MISO | Master In Slave Out | GPIO9 (21)  |
| 5    | GND  | -Power              | gnd (39)    |
| 6    | RST  | Reset-Circuit       | GPIO25 (22) |
| 7    | 3.3v | +Power              | 3.3V (1)    |

### Keyboard - 3x4 Membrane Matrix Keyboard

![3x4 Keyboard](docs/keyboard.png)

#### Connection

| No   | Name | Desc      | To          |
| -    | -    | -         | -           |
| 1    | R1   | Row 1     | GPIO22 (15) |
| 2    | R2   | Row 2     | GPIO27 (13) |
| 3    | R3   | Row 3     | GPIO17 (11) |
| 4    | R4   | Row 4     | GPIO4 (7)   |
| 5    | C1   | Column 1  | GPIO18 (12) |
| 6    | C2   | Column 2  | GPIO15 (10) |
| 7    | C3   | Column 3  | GPIO14 (8)  |

### Buzzer - Hot Passive Buzzer

![Passive Buzzer](docs/buzzer.png)

#### Connection

| No   | Name | Desc   | To         |
| -    | -    | -      | -          |
| 1    | VCC  | +Power | 5V (2)     |
| 2    | I/O  | Data   | GPIO1 (28) |
| 3    | GND  | -Power | gnd (39)   |

### Case - Nexus 4 box

 - this is a must
 - just buy a phone, throw it away and keep the box
 - most expensive piece of hardware here

## Demos
### 2020-02-29 evening
![Case](docs/2020-02-29-case.jpg)
![Internals](docs/2020-02-29-internals.jpg)
#### Video
[![Demo 2 (2020-02-29 evening)](http://img.youtube.com/vi/qaHyFebox_I/0.jpg)](http://www.youtube.com/watch?v=qaHyFebox_I)

### 2020-02-29 morning (video)
[![Demo 1 (2020-02-29 morning)](http://img.youtube.com/vi/ukjnKA4MB-E/0.jpg)](http://www.youtube.com/watch?v=ukjnKA4MB-E)
