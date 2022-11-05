let container = document.getElementById('container')
document.querySelector('body').style.backgroundColor = '#222528'
const {origin} = window.location;

fetch(`${origin}/getRestaurants/`).then(res => res.json()).then( data => {
        console.log(data)
        for (let restaurant of data["eateries"]) {
            let x = document.createElement('div')
            x.classList.add('card');
            x.style.display = "flex";
            x.style.flex = "20rem"
            x.style.backgroundColor = '#303134'
            x.style.margin = '0.5rem';
            x.style.alignItems = "center";

            let ig = document.createElement('img');
            ig.src = restaurant["image"];
            ig.classList.add("card-img-top");
            ig.alt = "...";

            let bdy = document.createElement('div');
            bdy.classList.add('card-body');

            let title = document.createElement('h5');
            title.innerText = restaurant['name'];
            title.style.color = 'aliceblue';
            title.style.fontWeight = 'bold';
            title.classList.add('card-title');

            let text = document.createElement('p');
            text.innerText = restaurant['location'];
            text.style.color = 'aliceblue';
            text.classList.add('card-text');

            let btn = document.createElement('a');
            btn.classList.add('btn');
            btn.innerText = "View Details";
            btn.classList.add('btn-primary');
            btn.style.width = "280px"
            btn.style.margin = "2px"

            bdy.appendChild(title);
            bdy.appendChild(text);

            x.appendChild(ig);
            x.appendChild(bdy);
            x.appendChild(btn);

            container.appendChild(x);
        }
    }
)