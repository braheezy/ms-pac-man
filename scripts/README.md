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


To calculate whether Pac-Man is in a pre-turn or post-turn state based on his position relative to the centerline of the turn, you can use a simple comparison of his position within his current tile. Here's how you can do it:

Assuming you have Pac-Man's current position in pixel coordinates (pacmanX and pacmanY) and you want to determine the state while he is making a right turn:

go

// Assuming TileWidth is the width of a single tile
centerlineX := pacmanTileX * TileWidth + TileWidth/2

// Calculate the distance of Pac-Man's current position to the centerline
distanceToCenterline := pacmanX - centerlineX

// Define a threshold value to determine when Pac-Man is in a pre-turn or post-turn state
threshold := TileWidth / 4  // Adjust this value as needed

if distanceToCenterline < -threshold {
    // Pac-Man is in a pre-turn state, meaning he's a few pixels away from the centerline on the left side
    // Handle pre-turn logic here
} else if distanceToCenterline > threshold {
    // Pac-Man is in a post-turn state, meaning he's a few pixels away from the centerline on the right side
    // Handle post-turn logic here
} else {
    // Pac-Man is in neither a pre-turn nor a post-turn state, indicating normal movement
    // Handle normal movement logic here
}