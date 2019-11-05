# Pin Usage

This document maps the pin usage of the Raspberry Pi.

See [hifiberry](https://www.hifiberry.com/docs/hardware/gpio-usage-of-hifiberry-boards/)
for more info on the HifiBerry connections.

## Pins

```
GPIO  1
GPIO  2   AMP2 config
GPIO  3   AMP2 config
GPIO  4   AMP2 mute (pull low to mute)
GPIO  5
GPIO  6
GPIO  7
GPIO  8   MCP3008 CS
GPIO  9   MCP3008 DOUT
GPIO 10   MCP3008 DIN
GPIO 11   MCP3008 CLK
GPIO 12
GPIO 13
GPIO 14
GPIO 15
GPIO 16
GPIO 17
GPIO 18   AMP2 sound interface
GPIO 19   AMP2 sound interface
GPIO 20   AMP2 sound interface
GPIO 21   AMP2 sound interface
GPIO 22   Coded Rotary Pin A
GPIO 23   Coded Rotary Pin B
GPIO 24   Coded Rotary Pin C
GPIO 25   Coded Rotary Pin D
```

## MCP3008

We're using an MCP3008 to convert the volume knob voltage 
to a digital reading.

Here's how it's wired:

```
3.3V      VDD
3.3V      VREF
GND       AGND
GPIO 11   CLK
GPIO 9    DOUT
GPIO 10   DIN
GPIO 8    CS
GND       DGND
```

See the [reference image](https://i1.wp.com/cdn-learn.adafruit.com/assets/assets/000/030/456/original/sensors_raspberry_pi_mcp3008pin.gif?resize=564%2C423&ssl=1)
for more info.
