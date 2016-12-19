/*


species.go implementation of species of genomes.

@licstart   The following is the entire license notice for
the Go code in this page.

Copyright (C) 2016 jin yeom, whitewolf.studio

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

As additional permission under GNU GPL version 3 section 7, you
may distribute non-source (e.g., minimized or compacted) forms of
that code without the copy of the GNU GPL normally required by
section 4, provided you include this license notice and a URL
through which recipients can access the Corresponding Source.

@licend    The above is the entire license notice
for the Go code in this page.


*/

package neat

// Species is an implementation of species of genomes in NEAT, which
// are separated by measuring compatibility distance among genomes
// within a population.
type Species struct {
	sid     int       // species ID
	age     int       // species age
	genomes []*Genome // genomes in this species
}

// NewSpecies creates a new species given a species ID, and the genome
// that first populates the new species.
func NewSpecies(sid int, g *Genome) *Species {
	return &Species{
		sid:     sid,
		age:     0,
		genomes: []*Genome{g},
	}
}

// SID returns this species' species ID.
func (s *Species) SID() int {
	return s.sid
}

// Age returns this species' age.
func (s *Species) Age() int {
	return s.age
}

// Genomes returns this species' member genomes.
func (s *Species) Genomes() []*Genome {
	return s.genomes
}

// AddGenome adds a new genome to this species.
func (s *Species) AddGenome(g *Genome) {
	s.genomes = append(s.genomes, g)
}

// sh implements a part of the explicit fitness sharing function, sh.
// If a compatibility distance 'd' is larger than the compatibility
// threshold 'dt', return 0; otherwise, return 1.
func sh(d, dt float64) float64 {
	if d > dt {
		return 0.0
	}
	return 1.0
}

// FitnessShare computes and assigns the shared fitness of genomes in
// this species, via explicit fitness sharing.
func (s *Species) FitnessShare(dt float64) {
	adjusted := make(map[int]float64)
	for _, g0 := range s.genomes {
		adjustment := 0.0
		for _, g1 := range s.genomes {
			adjustment += sh(g0.Compatibility(g1), dt)
		}
		if adjustment != 0.0 {
			adjusted[g0.gid] = g0.fitness / adjustment
		}
	}
	for i := range s.genomes {
		s.genomes[i].fitness = adjusted[s.genomes[i].gid]
	}
}
