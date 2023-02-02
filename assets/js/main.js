function switcher() {
    const city1 = document.getElementById('city1');
    const city2 = document.getElementById('city2');
    var option1 = city1.value;
    var option2 = city2.value;
    city1.value = option2;
    city2.value = option1;
    disable(2, false);
    disable(1, false);
    render(0);
}

function disableDropdown() {
    var selectobject;
    selectobject = document.getElementById("dropdownBranch").getElementsByTagName("option");
    selectobject[3].disabled = true;
    $('#dropdownBranch').selectpicker('render');
    selectobject = document.getElementById("pureDropDown").getElementsByTagName("option");
    for (z = 0; z < selectobject.length; z++) {
        selectobject[z].disabled = true;
    }
}

function disable(cityNumber, f) {
    const city1 = document.getElementById('city1');
    const city2 = document.getElementById('city2');
    if (cityNumber == 1) {
        const option2 = city2.value;
        for (z = 0; z < city1.length; z++) {
            city1[z].disabled = false;
        }
        city1[option2].disabled = true;
        if (f == true) {
            render(1);
        }
    }
    if (cityNumber == 2) {
        const option1 = city1.value;
        for (z = 0; z < city2.length; z++) {
            city2[z].disabled = false;
        }
        city2[option1].disabled = true;
        if (f == true) {
            render(2);
        }
    }
}

function render(cityNumber) {
    if (cityNumber == 0) {
        $('#city1').selectpicker('render');
        $('#city2').selectpicker('render');
    }
    if (cityNumber == 1) {
        $('#city1').selectpicker('render');
    }
    if (cityNumber == 2) {
        $('#city2').selectpicker('render');
    }
}

function handler() {
    let std = document.querySelector("#std");
    std.innerHTML = "";
    // window.alert("فرم را پر کنید")
    for (let i = 1; i < 6;) {
        cardmaker();
        i = i + 1;
    }
}

function cardmaker() {
    let std = document.querySelector("#std");
    let div = document.createElement("div");
    div.style = "margin: 10px;";
    let div2 = document.createElement("div");
    div2.classList.add("card");
    div.append(div2);
    let h5 = document.createElement("h5");
    h5.classList.add("card-header");
    h5.innerText = "Featured";
    div2.append(h5);
    let div3 = document.createElement("div");
    div3.classList.add("card-body");
    div2.append(div3);
    let h52 = document.createElement("h5");
    h52.classList.add("card-title");
    h52.innerText = "هواپیمایی X";
    div3.append(h52);
    let p1 = document.createElement("p");
    p1.classList.add("card-text");
    p1.innerText = "مشخصات سفر";
    div3.append(p1);
    let a = document.createElement("a");
    a.href = "/payment";
    let button = document.createElement("button");
    button.onclick = "/payment"
    button.classList.add("btn");
    button.classList.add("btn-primary");
    button.innerText = "خرید";
    a.append(button);
    div3.append(a);

    std.append(div);

}

function changeNumber(f) {
    const passengerNumber = document.getElementById('passenger-number');
    var pNum = parseInt(passengerNumber.value);
    if (f == 0) {
        passengerNumber.value = pNum - 1;
    }
    if (f == 1) {
        passengerNumber.value = pNum + 1;
    }
    checkNumber();
}

function checkNumber() {
    const passengerNumber = document.getElementById('passenger-number');
    var pNum = parseInt(passengerNumber.value);
    if(pNum <= 0) {
        passengerNumber.value = 1;
    }
}