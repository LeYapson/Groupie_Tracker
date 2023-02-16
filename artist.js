//? Cette fonction prend une URL d'API JSON en tant que paramètre.
//? Elle utilise la méthode fetch() pour envoyer une requête HTTP GET à cette URL.
//? Elle affiche ensuite le contenu dans le terminal
function getDataFromApi(apiUrl) {
    fetch(apiUrl)
      .then(response => response.json())
      .then(data => {
        console.log(data);
      })
      .catch(error => console.error(error));
}

getDataFromApi("https://groupietrackers.herokuapp.com/api/artists")