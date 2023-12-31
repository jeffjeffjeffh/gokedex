package main

import (
	"errors"
	"internal/pokeapi"
)

type Pokedex struct {
	data map[string]Pokemon
}

type Pokemon struct {
	name           string
	height, weight int
	stats          map[string]int
	types []string
}

func newPokedex() Pokedex {
	return Pokedex{
		data: map[string]Pokemon{},
	}
}

func newPokemon(res pokeapi.PokemonResponse) Pokemon {
	pokemon := Pokemon{
		name: res.Name,
		height: res.Height,
		weight: res.Weight,
		stats: map[string]int{},
		types: []string{},
	}

	for _, stat := range res.Stats {
		pokemon.stats[stat.Stat.Name] = stat.BaseStat
	}

	for _, t := range res.Types {
		pokemon.types = append(pokemon.types, t.Type.Name)
	}

	return pokemon
}

func (dex *Pokedex) get(name string) (Pokemon, bool) {
	val, ok := dex.data[name]
	if !ok {
		return Pokemon{}, false
	}
	return val, true
}

func (dex *Pokedex) list() ([]string, error) {
	if len(dex.data) == 0 {
		return []string{}, errors.New("your pokedex is empty! you'll never be the best at this rate")
	}

	names := []string{}
	for name := range dex.data {
		names = append(names, name)
	}

	return names, nil
}

func (dex *Pokedex) add(name string, pokemon Pokemon) {
	dex.data[name] = pokemon
}