/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package main

import (
	_ "image/png"
)

// This is a prototype version of the game for a proof-of-concept purpose.
// I will obviously refactor a huge part of the codebase later since this
// version is not stable.

// Build the server from the server module and run it. (Set up the address).
// Set the address also in the client package of the game module.
// Open the file user.json at the root of this project to set your username.
// Build the game from the game module and run it.

// The server generates random matches each x seconds. Just open your game.

func main() {
	Run()
}
