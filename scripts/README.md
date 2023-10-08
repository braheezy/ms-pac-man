These scripts help split a PNG sprite sheet. Here's what I did:

- Find the PNG sprite sheet online and download it.
- Open the PNG sheet in Krita and scale it up so each character asset is 32x32 pixel
- Select and copy that to a new PNG and save it off as `all-characters.sh`

For levels,
- In Krita, move the PNG sheet layer until the grid lines up appropriately for 16x16 tiles for the levels
- Select and copy that layer to a new PNG
- Mess with the Canvas until there's the right aspect ratio. I needed to added 8 pixels vertically to ensure when split, every tile is 16px and the tiles are split the right way to build a level.
- Save off `level1.png`

Put the PNGs in the right places in `assets/` then run the scripts in this directory.
