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

// TODO ----------------------------------------|
// TODO: Other files ---------------------------|
// TODO ----------------------------------------|

// Closes the details window when clicking on the closing cross.
function addClosableListener() {
    var quits = document.querySelectorAll(".data-view-container")
    if (quits) {
        quits.forEach(element => {
            element.querySelector(".quit").addEventListener("click", function() {
                element.remove()
            })
        })
    }
}

addClosableListener()

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
                        }
                    })}
                removeClassFromAll(categories, "active")
                element.classList.add("active")
            })
        })
    }
}

switchCategories()
