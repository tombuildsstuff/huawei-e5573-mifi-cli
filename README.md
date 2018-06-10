## CLI for the Huawei E5573 Mifi

This is a CLI which currently displays information for [the Huawei E5573 Mifi](https://consumer.huawei.com/in/support/mobile-broadband/e5573/), since I find the official UI lacking.

## Requirements

- [A Huawei E5573 Mifi](https://consumer.huawei.com/in/support/mobile-broadband/e5573/)
- Go 1.10+ (it _may_ work in an earlier version, but I've not tested it)

## Example Usage

```
$ ./huawei-e5573-mifi-cli

Mifi Status:

  Network:
    Signal Strength: 5/5 bars
    Network:         "I WIND" (ID: 22288 | Mode: 4G/LTE Enabled)
    Bandwidth used:  82.37MB down / 12.13MB up
    Connected for:   2 hours (175 minutes)

  Information:
    Battery: 100%
    Wifi: "Example SSID" (Country: IT | 1 devices connected)
```

## Command Line Arguments

- `-dashboard` - opens the Web Portal for the Mifi in your default web browser.
- `-endpoint` - specifies the endpoint to the Web Portal. Defaults to `http://192.168.1.1`.
- `-help` - displays all possible command line arguments.
- `-version` - displays the version of the CLI tool.


## Possible Future Extensions

- Toggling 4G on and off
- Toggling Automatic/Manual Network Selection
- Viewing the SMS Inbox
