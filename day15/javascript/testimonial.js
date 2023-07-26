

const promise = new Promise((resolve, reject) => {
    const req = new XMLHttpRequest();

    req.open("get", "https://api.npoint.io/d65c77368442553efc2e", true);

    req.onload = function () {

        if (req.status === 200 ) {
            resolve(JSON.parse(req.responseText));
            console.log(req.status);

        }else if (req.status >= 400) {
            reject("error broh");
        }

    }

    req.onerror = function () {
        reject("erros masbro")
    }

    req.send();
});


const ul = document.querySelector(".testimonial");


async function allData() {
    ul.innerHTML = "";
    try {
                const res = await promise
                dataTestimonial = res
        
                console.log(dataTestimonial);
                
                dataTestimonial.forEach((item) => {
                    ul.innerHTML += `<li class="testimonial-item">
                        <div class="wrapper-testimonial-item">
                            <p class="item-rating">${item.rating}<i class="fa fa-star"></i></p>
                            <img src="${item.img}" alt="gambar">
                            <p class="item-comment">${item.comment}</p>
                            <span><i>${item.user}</i></span>
                            
                        </div>
                    </li>`
                })
                
            } catch (error) {
                        console.log(error);
                        
                    }
    
}

allData()
//Memfilter dengan Array filter

function dataFilter(rating) {
    
    const tempData = dataTestimonial.filter((item) => item.rating === rating );

    ul.innerHTML = "";
    tempData.forEach((item) => {
        ul.innerHTML += `<li class="testimonial-item">
            <div class="wrapper-testimonial-item">
            <p class="item-rating">${item.rating}<i class="fa fa-star"></i></p>
                <img src="${item.img}" alt="gambar">
                <p class="item-comment">${item.comment}</p>
                <span><i>${item.user}</i></span>
                
            </div>
        </li>`
    })


}

