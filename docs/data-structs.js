
let Room = {
    title: "Quarto solteiro com uma cama",
    code: "SOLT-1",
    beds: [
        {
            size: "small",
            type: "single"
        }
    ]
}

let Room = {
    title: "Quarto solteiro com duas camas",
    code: "SOLT-2",
    beds: [
        {
            size: "small",
            type: "single"
        },
        {
            size: "small",
            type: "single"
        }
    ]
}

let Room = {
    title: "Quarto solteiro com três camas",
    code: "SOLT-3",
    bed: [
        {
            size: "small",
            type: "single"
        },
        {
            size: "small",
            type: "single"
        },
        {
            size: "small",
            type: "single"
        }
    ]
}

let Room = {
    title: "Quarto casal cama king size",
    code: "CASAL-1",
    bed: [
        {
            size: "king",
            type: "double"
        }
    ]
}

let Room2 = {
    title: "Quarto casal com cama king size e uma cama de solteiro",
    code: "CASAL-2",
    bed: [
        {
            size: "king",
            type: "double"
        },
        {
            size: "small",
            type: "single"
        }
    ],
    note: "Ideal for casal + children"
}

let PricingConfig = {
    date: "2021-02-05"
}

let Ocupation = {
    date: "2021-02-20",
    rooms: []
}