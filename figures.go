// mostly from https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
package main

const (
	block = `
		##
		##
	`

	beehive = `
		_##_
		#  #
		_##
	`

	loaf = `
		_##_
		#__#
		_#_#
		__#
	`

	boat = `
		##_
		#_#
		_#
	`

	tub = `
		_#_
		#_#
		_#
	`

	blinker = `
		###
	`

	toad = `
		_###
		###
	`

	beacon = `
		##__
		#
		___#
		__##
	`

	pulsar = `
		__###___###__
		_
		#____#_#____#
		#____#_#____#
		#____#_#____#
		__###___###
		_
		__###___###
		#____#_#____#
		#____#_#____#
		#____#_#____#
		_
		__###___###
	`

	pentadecathlon = `
		_###_
		#___#
		#___#
		_###
		_
		_
		_
		_
		_###
		#___#
		#___#
		_###
	`

	glider = `
		##_
		#_#
		#_
	`

	lwss = `
		_##__
		####
		##_##
		__##
	`

	mwss = `
		_###__
		#####
		###_##
		___##
	`

	hwss = `
		_####__
		######
		####_##
		____##
	`

	rPentomino = `
		_##
		##
		_#
	`

	dieHard = `
		______#_
		##
		_#___###
	`

	acorn = `
		_#_____
		___#
		##__###
	`

	gosperGG = `
		________________________#___________
		______________________#_#
		____________##______##____________##
		___________#___#____##____________##
		##________#_____#___##
		##________#___#_##____#_#
		__________#_____#_______#
		___________#___#
		____________##
	`

	gun = `
		##_____##________________________
		##_____##
		_
		____##
		____##
		_
		_
		_
		_
		____________________##_##
		___________________#_____#
		___________________#______#__##
		___________________###___#___##
		________________________#
		_
		_
		_
		__________________##
		__________________#
		___________________###
		_____________________#
	`

	grow1 = `
		______#_
		____#_##
		____#_#
		____#
		__#
		#_#
	`

	grow2 = `
		###_#
		#
		___##
		_##_#
		#_#_#
	`

	grow3 = `
		########_#####___###______#######_#####
	`
)
