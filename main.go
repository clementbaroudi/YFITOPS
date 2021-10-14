import {
  "fmt"
  "container/list"
}

type Utilisateur struct(){
  num_utilisateur int 
  list_music list
  genre_fav string 
}

type Music struct() {
  titre string 
  album string
  artiste string
  genre string
}

func find_genre_fav (Utilisateur a) (Music) {
  
}



func find_musique ( param ... string, int distance ) (Musique){
  // distance est demandé à l'utilisateur et permet de savoir si il veut découvrir une musique donc on recherche une musique plus loin ou si il veut une musique qu'il est sûr d'aimer

  if param == nil {  //pas de filtre donc on parcourt tout 
  }
  
  if param == "genre" {
  }
  
  if param == "artiste" {
  }
  
  if param == "album" {
  }
