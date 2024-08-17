package main

import (
	"errors"
	"fmt"
)

// ax + by = gcd(a, b)
func egcd(a, b int) (x, y, gcd int) {
	if b == 0 {
		return 1, 0, a
	}

	x1, y1, gcd := egcd(b, a%b)

	return y1, x1 - y1*(a/b), gcd
}

func lcm(a, b int) int {
	_, _, gcd := egcd(a, b)

	return a * b / gcd
}

// x = val (modulo mod)
type ModularEqn struct {
	val int
	mod int
}

func (eqn ModularEqn) String() string {
	return fmt.Sprintf("x ≡ %d (mod %d)", eqn.val, eqn.mod)
}

func solveModularEqn(eqn1, eqn2 ModularEqn) (ModularEqn, error) {
	x, _, gcd := egcd(eqn1.mod, eqn2.mod)

	if (eqn1.val-eqn2.val)%gcd != 0 {
		return ModularEqn{}, errors.New(fmt.Sprintf("Cannot solve congruences [%v] and [%v] since [%d %% %d != 0]", eqn1, eqn2, (eqn1.val - eqn2.val), gcd))
	}

	lmbda := (eqn1.val - eqn2.val) / gcd

	return ModularEqn{
		val: eqn1.val - eqn1.mod*x*lmbda,
		mod: lcm(eqn1.mod, eqn2.mod),
	}, nil
}

func solveModularEqns(eqns []ModularEqn) (ModularEqn, error) {
	res := eqns[0]

	for _, eqn2 := range eqns[1:] {
		var err error = nil

		res, err = solveModularEqn(res, eqn2)
		if err != nil {
			return ModularEqn{}, err
		}
	}

	return res, nil
}
