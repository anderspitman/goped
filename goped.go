package goped

import (
        //"fmt"
        "bufio"
        "strings"
        "errors"
)


type Pedigree struct {
        Individuals map[string]*Individual
}

type Individual struct {
        FamilyId string
        IndividualId string
        PaternalId string
        MaternalId string
        Sex string
        Phenotype string
}

func NewPedigree() *Pedigree {
        return &Pedigree{Individuals: make(map[string]*Individual)}
}

func Parse(pedStr string) (*Pedigree, error) {

        ped := NewPedigree()

        scanner := bufio.NewScanner(strings.NewReader(pedStr))
        for scanner.Scan() {
                line := scanner.Text()

                if strings.HasPrefix(line, "#") {
                        continue
                }

                indi, err := parseLine(line)
                if err != nil {
                        return nil, errors.New("Error parsing row")
                }

                ped.Individuals[indi.IndividualId] = indi
        }

        return ped, nil
}

func parseLine(line string) (*Individual, error) {
        fields := strings.Fields(line)

        if len(fields) < 6 {
                return nil, errors.New("Not enough fields in PED row")
        }

        indy := &Individual {
                FamilyId: fields[0],
                IndividualId: fields[1],
                PaternalId: fields[2],
                MaternalId: fields[3],
                Sex: fields[4],
                Phenotype: fields[5],
        }

        return indy, nil
}
