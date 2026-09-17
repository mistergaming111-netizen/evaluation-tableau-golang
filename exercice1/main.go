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

	
	fmt.Println("=== ANALYSE ===")
	fmt.Println()

	soldatVie := trouverPlusDeVie(equipe)
	fmt.Println("Soldat avec le plus de vie :", soldatVie.nom)
	fmt.Println("Vie :", soldatVie.vie)
	fmt.Println()

	soldatAttaque := trouverPlusDAttaque(equipe)
	fmt.Println("Soldat avec la plus grande attaque :", soldatAttaque.nom)
	fmt.Println("Attaque :", soldatAttaque.attaque)
	fmt.Println()

	moyenne := calculerVieMoyenne(equipe)
	fmt.Printf("Vie moyenne : %.2f\n", moyenne)
	fmt.Println()

	faibles := compterFaibles(equipe)
	fmt.Println("Soldats avec moins de 800 PV :", faibles)
	fmt.Println()

	
	fmt.Println("=== BATAILLE ===")
	fmt.Println()

	var nbAttaques int
	fmt.Print("Nombre d'attaques ennemies : ")
	fmt.Scan(&nbAttaques)
	fmt.Println()

	for i := 1; i <= nbAttaques; i++ {
		var degats int
		fmt.Printf("Attaque %d : ", i)
		fmt.Scan(&degats)
		fmt.Println()

		
		attaquerEquipe(&equipe, degats)

		
		fmt.Printf("=== APRÈS L'ATTAQUE %d ===\n\n", i)
		afficherEtat(equipe)
		fmt.Println()
	}
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
	for i := 0; i < len(equipe); i++ {
		if equipe[i].attaque > meilleur.attaque {
			meilleur = equipe[i]
		}
	}
	return meilleur
}

func calculerVieMoyenne(equipe [6]Soldat) float64 {
	somme := 0
	for i := 0; i < len(equipe); i++ {
		somme += equipe[i].vie
	}
	return float64(somme) / float64(len(equipe))
}

func compterFaibles(equipe [6]Soldat) int {
	compteur := 0
	for i := 0; i < len(equipe); i++ {
		if equipe[i].vie < 800 {
			compteur++
		}
	}
	return compteur
}


func attaquerEquipe(equipe *[6]Soldat, degats int) {
	for i := 0; i < len(equipe); i++ {
		
		if equipe[i].vie > 0 {
			equipe[i].vie -= degats
			
			if equipe[i].vie < 0 {
				equipe[i].vie = 0
			}
		}
	}
}


func afficherEtat(equipe [6]Soldat) {
	for i := 0; i < len(equipe); i++ {
		if equipe[i].vie > 0 {
			fmt.Printf("%s : %d PV\n", equipe[i].nom, equipe[i].vie)
		} else {
			fmt.Printf("%s : KO\n", equipe[i].nom)
		}
	}
}