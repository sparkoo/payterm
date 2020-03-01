# payterm
payment terminal toy project

## Used Hardware and connection
Price of all hardware, except Raspberry Pi, is less than 10$.
### Controller - Raspberry Pi Zero W
 - Pinout https://pinout.xyz/
 - numbers in `( )` in *Connection* tables are physical pins.

### Display - 1602A v2.0
 - 16x2 characters
 - +some 2 random pots I've found at home to control contrast (CPot) and backlight (LPot)

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
 - support cards - S50, S70, UltraLight, Pro, Desfire
 - +some blank S50 cards

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
### Buzzer - Hot Passive Buzzer

## Demos
### 2020-02-29 evening
[![Demo 2 (2020-02-29 evening)](http://img.youtube.com/vi/qaHyFebox_I/0.jpg)](http://www.youtube.com/watch?v=qaHyFebox_I)

### 2020-02-29 morning
[![Demo 1 (2020-02-29 morning)](http://img.youtube.com/vi/ukjnKA4MB-E/0.jpg)](http://www.youtube.com/watch?v=ukjnKA4MB-E)
