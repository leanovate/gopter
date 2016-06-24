package gopter

type derivedGen struct {
	biMapper *BiMapper
	upGens   []Gen
}

func (d *derivedGen) Generate(params *GenParameters) *GenResult {
	return nil
}

// DeriveGen derives a generator with shrinkers from a sequence of other
// generators mapped by a bijective function (BiMapper)
func DeriveGen(biMapper *BiMapper, gens ...Gen) Gen {
	derived := &derivedGen{
		biMapper: biMapper,
		upGens:   gens,
	}
	return derived.Generate
}
