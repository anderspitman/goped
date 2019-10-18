package goped

import (
        "testing"
        "fmt"
)

func TestParse(t *testing.T) {
        ped, err := Parse("FAM001  1  0 0  1  2\nFAM001  2  0 0  1  2")

        if err != nil {
                t.Errorf("Parse failed: %s", err)
        }

        fmt.Println(ped)
}

func TestParseLine(t *testing.T) {
        indi, err := parseLine("FAM001  1  0 0  1  2")
        if err != nil {
                t.Errorf("parseLine failed: %s", err)
        }

        if indi.FamilyId != "FAM001" {
                t.Errorf("Wrong FamilyId")
        }

        if indi.IndividualId != "1" {
                t.Errorf("Wrong IndividualId")
        }

        if indi.PaternalId != "0" {
                t.Errorf("Wrong PaternalId")
        }

        if indi.MaternalId != "0" {
                t.Errorf("Wrong MaternalId")
        }

        if indi.Sex != "1" {
                t.Errorf("Wrong Sex")
        }

        if indi.Phenotype != "2" {
                t.Errorf("Wrong Phenotype")
        }
}
