import {
  "fmt"
  "container/list"
}

type Utilisateur struct{
  num_utilisateur int 
  list_music list
  genre_fav string 
}

type Music struct {
  titre, album, artiste, genre string 
}


type Genre struct { // ce struct sert que pour la recherche du genre fav
  nom string
  nombre int
}

//sert à trouver le genre favori (si le genre favori compte dans la recehrche de la musqiue, on aura besoin de cette methode)
func find_genre_fav (Utilisateur a) (string) { 
  genres := list.New() // table à double entrée des genres avec en face de cahque genre le nb de musiques correspondant dans la liste de l'utilisateur
  existe := false // boolean pour savoir si le genre est dans la liste
  for _, msc := range a.list_music { // on parcours toutes les musiques de l'utilisateur
    for i := genres.Front(); i != nil ; i = genres.Next() { //on cehrche si le genre de la musique est dans la liste de genres
      if msc.genre == i.nom { // si oui on ajoute 1 au total de musiques
        i.nombre += 1  
        existe = true
      }
    }
    
    if existe == false { // si non on ajoute le genre dans la liste 
      genres.PushBack( Genre(nom: msc.titre, nombre: 1)) // pushback c'est pour ajouter un item à la fin de la liste
    }
    existe == false
  }
  
  //on détermine ensuite le genre fav en cherchant quel genre a le nb de musiques le plus élevé
  var genrefav string
  var nbgenrefav = 0
  for i := genres.Front(); i != nil ; i = genres.Next() { 
    if i.nombre > nbgenrefav {
      nbgenrefav = i.nombre
      genrefav = i.nom
    }
  return genrefav
}



func find_musique ( param ... string, int distance ) (Musique){
  // distance est demandé à l'utilisateur et permet de savoir si il veut découvrir une musique donc on recherche une musique plus loin ou si il veut une musique qu'il est sûr d'aimer
  
  switch {
    case param == nil :  //pas de filtre donc on parcourt tout 
  
  
    case param == "genre" :
  
  
    case param == "artiste" : 
  
  
    case param == "album" : 
  }
  
