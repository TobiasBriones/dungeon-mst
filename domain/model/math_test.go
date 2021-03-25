/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 */

package model

import (
	"testing"
)

func TestRectIntersect(t *testing.T) {
	r1 := NewRect(
		0,
		0,
		50,
		40,
	)
	r2 := NewRect(
		10,
		20,
		30,
		30,
	)
	r3 := NewRect(
		24,
		30,
		30,
		70,
	)
	r4 := NewRect(
		30,
		20,
		330,
		300,
	)
	r5 := NewRect(
		100,
		20,
		300,
		30,
	)
	r6 := NewRect(
		10,
		200,
		80,
		230,
	)

	if !r1.Intersects(&r2) {
		t.Fatal("FAILED R1-R2")
	}
	if !r2.Intersects(&r1) {
		t.Fatal("FAILED R1-R2")
	}
	if !r1.Intersects(&r3) {
		t.Fatal("FAILED R1-R3")
	}
	if !r3.Intersects(&r1) {
		t.Fatal("FAILED R1-R3")
	}
	if !r1.Intersects(&r4) {
		t.Fatal("FAILED R1-R4")
	}
	if !r4.Intersects(&r1) {
		t.Fatal("FAILED R1-R4")
	}
	if r1.Intersects(&r5) {
		t.Fatal("FAILED R1-R5")
	}
	if r5.Intersects(&r1) {
		t.Fatal("FAILED R1-R5")
	}
	if r1.Intersects(&r6) {
		t.Fatal("FAILED R1-R6")
	}
	if r6.Intersects(&r1) {
		t.Fatal("FAILED R1-R6")
	}
}
