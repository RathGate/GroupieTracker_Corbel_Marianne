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


// ?CATEGORIES PAGE:

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


// ?CARDS TEMPLATE
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

// ?SEARCH PAGE - Form
let form = document.getElementById("search-form")
if (form) {
    form.addEventListener("submit", function(e) {
        e.preventDefault()
        
        // Hides all open filter options menu when submitting form
        removeClassFromAll(document.querySelectorAll(".filter-options-container"), "visible")

        // Sends the formdata and replaces the content of a div with the response
        $.ajax({
            type: "POST",
            url: "/search",
            data: $("#search-form").serialize(),
            success: function(data) {
                console.log(data)
                document.querySelector(".container").innerHTML = data
                currentPageIndex = 0
            },
            error: function() {
                console.log("nope")
            }
        })
    })
}


// ?SEARCH PAGE - Filters
let filters = document.querySelectorAll(".filter-container")
if (filters) {
    filters.forEach((filter, index) => {
        // Tracks the checkboxes
        var allcheckboxes = filter.querySelectorAll("input[type='checkbox']")
        if (allcheckboxes) {
            allcheckboxes.forEach(checkbox => {
                checkbox.addEventListener("change", function() {
                    let checkCount = countAllChecked(allcheckboxes)
                    document.querySelectorAll(".count .number")[index].innerHTML = checkCount == 0 || checkCount == allcheckboxes.length ? "all" : checkCount
                })
            })
        }

        // Show or hide full filter menu
        filter.querySelector(".filter-title").addEventListener("click", function() { 
            if (filter.querySelector(".filter-options-container").classList.contains("visible")) {
                filter.querySelector(".filter-options-container").classList.remove("visible")
            } else {
                removeClassFromAll(document.querySelectorAll(".filter-options-container"), "visible")
                filter.querySelector(".filter-options-container").classList.add("visible")    
            }
        })
    })
}

// Count all checkboxes checked in a array of checkboxes.
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


document.addEventListener("click", function(e) {
    let filters = document.querySelectorAll(".filter-options-container")
    let $target = $(e.target)

    if ($target.parents(".filter-options-container").length > 0 
    || $target.hasClass("filter-options-container")
    || $target.parents(".filter-title.activable").length > 0 
    || $target.hasClass("filter-title activable")){
        return
    }
    
    removeClassFromAll(filters, "visible")
})

// ?INDEX PAGE 

let refreshBtn = document.querySelector(".refresh")
let parentDiv = document.querySelector(".item-to-replace")
if (refreshBtn && parentDiv) {
    refreshBtn.addEventListener("click", function() {
        let randomInt = Math.floor(Math.random() * 389) + 1
        let jsonData;
        $.ajax({
            type: "POST",
            url: "/",
            data: {"retrieveNew": randomInt},
            success: function(template) {
                
                parentDiv.innerHTML = template
                document.querySelector("#id-to-replace").innerHTML = randomInt
            }})
        $.ajax({
            url: 'https://botw-compendium.herokuapp.com/api/v2/entry/' + randomInt,
            success: function(data) {
             document.querySelector("#json-to-replace").innerHTML = JSON.stringify(data, null, 2)
        }}); 
    })
}

let jsonToReplace = document.querySelector("#json-to-replace")
let initialData = {"data":{"category":"creatures","common_locations":["Gerudo Desert"],"description":"This is Riju's own sand seal. It may look intense, but she dotes on it regularly; the ribbon it wears was a gift form her, and it even has its own pen in Gerudo Town. It's far more agile than any other sand seal and far more outgoing. An ever-reliable partner to Riju, Patricia is always ready to take off through the desert at a moment's notice.","drops":[],"id":8,"image":"https://botw-compendium.herokuapp.com/api/v2/entry/patricia/image","name":"patricia"}}
if (jsonToReplace) {
    jsonToReplace.innerHTML = JSON.stringify(initialData, null, 2)
}