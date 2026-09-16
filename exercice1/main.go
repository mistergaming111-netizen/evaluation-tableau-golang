package main

import "fmt"

type Soldat struct {
	nom     string
	vie     int
	attaque int
}


func main() {
	equipe := [6]Soldat{
		{"Arthas", 1200, 250},
		{"Kael", 850, 320},
		{"Thrall", 1500, 180},
		{"Sylvanas", 700, 400},
		{"Garrosh", 1000, 280},
		{"Jaina", 500, 450},
	}
	afficherEquipe(equipe)
}


func afficherEquipe(equipe [6]Soldat) {
	fmt.Println("=== ÉQUIPE ===")
	fmt.Println()
	for i := 0; i < len(equipe); i++ {
		fmt.Println(equipe[i].nom)
		fmt.Println("Vie :", equipe[i].vie)
		fmt.Println("Attaque :", equipe[i].attaque)
		fmt.Println()
	}
}


func trouverPlusDeVie(equipe [6]Soldat) Soldat {
	meilleur := equipe[0]
	for i := 0; i < len(equipe); i++ {
		if equipe[i].vie > meilleur.vie {
			meilleur = equipe[i]
		}
	}
	return meilleur
}

func trouverPlusDAttaque(equipe [6]Soldat) Soldat {
	meilleur := equipe[0]
	for i := 0; i< len(equipe); i++ {
		if equipe[i].attaque > meilleur.attaque {
			meilleur = equipe[i]
		}
	}
	return meilleur
}
