
let inputName = document.querySelector("#name-input");
let email = document.querySelector("#email");
let number = document.querySelector("#number");
let radioButton = document.querySelector("#radio-button");
let message = document.querySelector("#textarea");
const form = document.querySelector("#form");



form.addEventListener("submit", (e)=>{
    e.preventDefault();
    
    const valueName = inputName.value
    const valueEmail = email.value
    const valueNumber = number.value
    const valueRb = radioButton.value
    const valueMessage = message.value

    let dataStatic = {
        valueName,
        valueEmail,
        valueNumber,
        valueRb,
        valueMessage
    }


    console.log(dataStatic);


    const emailReceiver = "dragoriza123@gmail.com";

    let anchor = document.createElement("a");

    anchor.href = `mailto:${emailReceiver}?subject=${valueRb}&body=konnichiwa hajimimashite ${valueName},\n${valueMessage}`;
    anchor.click();
})