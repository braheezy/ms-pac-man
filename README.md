# Ms. Pac-Man
:construction: WIP :construction:

This repo is an attempt to write the original Ms. Pac-Man in Go. By original, I mean original assets and as high fidelity game logic as possible.

- [Sprite sheet](https://www.spriters-resource.com/arcade/mspacman/sheet/21043/)
- [Sound](https://www.sounds-resource.com/arcade/mspacman/sound/16742/)
- Logic:
    - [Pac-Man Dossier](https://www.gamedeveloper.com/design/the-pac-man-dossier)
    - [Pac-Man Ghost Behavior](https://gameinternals.com/understanding-pac-man-ghost-behavior)

For now I am assuming, Ms. Pac-Man game logic is the same as the original.

## Development
Install Go and required system packages. On Fedora-based distros:

    yum install go libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel libglvnd-devel libXxf86vm-devel

Then see `make help` for more workflows.