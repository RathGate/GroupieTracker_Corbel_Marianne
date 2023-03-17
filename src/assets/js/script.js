// ? This JS file contains functions and variables that apply 
// ? to the "base.html" file.

// Function linked to the burger menu icon.
document.querySelector(".nav-menu-icon").addEventListener("click", function() {
    document.querySelector(".mobile-menu").classList.toggle("visible")
})

// Removes the visible attribute when resizing over the max size of 
// the burger menu.
window.addEventListener("resize", function() {
    if (this.innerWidth > 750) {
        this.document.querySelector(".mobile-menu").classList.remove("visible")
    } 
})

// Removes the `className` class from all elements in the `divArr` array.
function removeClassFromAll(divArr, className) {
    if (divArr.length == 0) {
        return
    }
    divArr.forEach(div => {
        div.classList.remove(className)
    })
}

let currentPageIndex = 0

// TODO ----------------------------------------|
// TODO: Other files ---------------------------|
// TODO ----------------------------------------|

// Switches the current category on click (category page only)
function switchCategories() {
    var categories = document.querySelectorAll(".cat-icon")
    if (categories) {
        categories.forEach((element, index) => {
            element.addEventListener("click", function() {
                if (!element.classList.contains("active")) {
                    $.ajax({
                        type: "POST",
                        url: "/categories",
                        data: { "category-id":  index},
                        success: function(data) {
                           document.querySelector(".cards-container").innerHTML = data
                           currentPageIndex = 0
                        }
                    })}
                removeClassFromAll(categories, "active")
                element.classList.add("active")
            })
        })
    }
}

switchCategories()


// !CARDS

let cards = document.querySelectorAll(".card-item")
if (cards) {
    cards.forEach(card => {
        let id = card.querySelector(".item-id").innerHTML
        card.addEventListener("click", function() {
            $.ajax({
                type: "POST",
                url: "/categories",
                data: { "item-id":  id},
                success: function(data) {
                    console.log("oui")
                    // document.getElementById("test").innerHTML = data
                    console.log(data)
                }
            })
        })
    })
}

let form = document.getElementById("search-form")
if (form) {
    form.addEventListener("submit", function(e) {
        e.preventDefault()

        removeClassFromAll(document.querySelectorAll(".filter-options-container"), "visible")

        const formData = new FormData(e.target)
        const formDataObj = {};
        formData.forEach((value, key) => (formDataObj[key] = value));
        console.log(formDataObj)
        $.ajax({
            type: "POST",
            url: "/search",
            data: formDataObj,
            success: function(data) {
                document.querySelector(".cards-container").innerHTML = data
                currentPageIndex = 0
            }
        })

    })
}


let filters = document.querySelectorAll(".filter-container")
if (filters) {
    filters.forEach(element => {
        element.querySelector(".filter-title").addEventListener("click", function() { 
            if (element.querySelector(".filter-options-container").classList.contains("visible")) {
                element.querySelector(".filter-options-container").classList.remove("visible")
            } else {
                removeClassFromAll(document.querySelectorAll(".filter-options-container"), "visible")
                element.querySelector(".filter-options-container").classList.add("visible")    
            }
        })
    })
}

if (filters) {
    filters.forEach((filter, index) => {
        var allcheckboxes = filter.querySelectorAll("input[type='checkbox'")
        if (allcheckboxes) {
            allcheckboxes.forEach(checkbox => {
                checkbox.addEventListener("change", function() {
                    let checkCount = countAllChecked(allcheckboxes)
                    document.querySelectorAll(".count .number")[index].innerHTML = checkCount == 0 || checkCount == allcheckboxes.length ? "all" : checkCount
                })
            })
        }
    })

}

function countAllChecked(checkboxes) {
    if (!checkboxes) {
        return 0
    }
    let total = 0
    checkboxes.forEach(checkbox => {
        if (checkbox.checked) {
            total++
        }
    })
    return total
}
// document.body.addEventListener('click', function( event ){
// 	if (!document.querySelector(".filter-options-container").contains( event.target ) && document.querySelector(".filter-options-container").classList.contains("visible") ) {
//         document.querySelector(".filter-options-container").classList.remove("visible")
//     }
// });