const hamburgerMenu = document.querySelector(".hamburgerMenu");
const navMenu = document.querySelector(".nav-menu");


hamburgerMenu.addEventListener("click", ()=>{
    hamburgerMenu.classList.toggle("active");
    navMenu.classList.toggle("active");
})

