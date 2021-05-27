# ff-pws

## Motivation

Dump passwords on a windows machine using the golang port of laZagne.
Especially when you want to migrate a windows machine to a sane OS (linux lel) this is
really handy.

## Installation (on a deban-ish host, rest is crappy anyways lel)

Install gcc-\* shizzle:

```
sudo apt update
sudo apt install gcc-multilib gcc-mingw-w64
make # binary is called ff-pws.exe
```

## Usage (on a Windows box, obviously)

```
./ff-pws.exe > pws.txt
```

## License

This project is licensed under the GPL-3 license.
