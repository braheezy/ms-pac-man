# Ms. Pac-Man
:construction: WIP :construction:

This repo is an attempt to write the original Ms. Pac-Man in Go. By original, I mean original assets and as high fidelity game logic as possible.

- [Sprite sheet](https://www.spriters-resource.com/arcade/mspacman/sheet/21043/)
- [Sound](https://www.sounds-resource.com/arcade/mspacman/sound/16742/)
- Logic:
    - [Pac-Man Dossier](https://www.gamedeveloper.com/design/the-pac-man-dossier)
    - [Pac-Man Ghost Behavior](https://gameinternals.com/understanding-pac-man-ghost-behavior)

For now I am assuming, Ms. Pac-Man game logic is the same as the original.
## Getting Started (for developers)

There are quite a few system libraries required to build the binary. Below are instructions for installing the necessary software.

### Fedora

The commands are kinda rough. There might be an easier way to do this.

```
sudo dnf install libXinerama-devel libXi-devel libglvnd-devel mesa-libGL-devel opengl-devel

sudo dnf groupinstall "X Software Development"
```

If a header file or something like that is missing, you can run `dnf provides */<filename>` to find out which package installs it.

I think for some of these header files, there may be multiple providers, so I picked the one that seemed the most relevant.


## Building & Testing

After a `git clone`, to get up and running with development, run `make build` to build the binary first. After that, you can run `make run` to test the binary.

