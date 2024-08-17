package main

import (
	"errors"
	"fmt"
	"math/big"
)

// x = val (modulo mod)
type ModularEqnBig struct {
	val *big.Int
	mod *big.Int
}

func (eqn ModularEqnBig) String() string {
	return fmt.Sprintf("x ≡ %d (mod %d)", eqn.val, eqn.mod)
}

func solveModularEqnBig(eqn1, eqn2 ModularEqnBig) (ModularEqnBig, error) {
	var x = new(big.Int) 
	var y = new(big.Int)
	var gcd = new(big.Int)

	gcd = gcd.GCD(x, y, eqn1.mod, eqn2.mod)

	var diff *big.Int = new(big.Int)
	diff = diff.Sub(eqn1.val, eqn2.val)

	var diffmodgcd *big.Int = new(big.Int)
	diffmodgcd = diffmodgcd.Mod(diff, gcd)

	if diffmodgcd.Cmp(big.NewInt(0)) != 0 {
		return ModularEqnBig{}, errors.New(fmt.Sprintf("Cannot solve congruences [%v] and [%v] since [%d %% %d != 0 (= %d)]", eqn1, eqn2, diff, gcd, diffmodgcd))
	}

	lmbda := new(big.Int)
	lmbda = lmbda.Div(lmbda.Sub(eqn1.val, eqn2.val), gcd)

	val := new(big.Int)
	val = val.Sub(eqn1.val, val.Mul(eqn1.mod, val.Mul(x, lmbda)))

	mod := new(big.Int)
	prod := new(big.Int)
	mod = mod.Div(prod.Mul(eqn1.mod, eqn2.mod), mod.GCD(x, y, eqn1.mod, eqn2.mod))

	soln := ModularEqnBig{
		val: val.Mod(val, mod),
		mod: mod,
	}

	fmt.Printf("Soln to [%v] and [%v] is [%v]\n", eqn1, eqn2, soln)

	return soln, nil
}

func solveModularEqnsBig(eqns []ModularEqnBig) (ModularEqnBig, error) {
	for _, eqn := range eqns {
		fmt.Println(eqn)
	}

	res := eqns[0]

	for _, eqn2 := range eqns[1:] {
		var err error = nil

		res, err = solveModularEqnBig(res, eqn2)
		if err != nil {
			return ModularEqnBig{}, err
		}
	}

	return res, nil
}
