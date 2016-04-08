form taxOfficeExample { 
    "Did you sell a house in 2010?" hasSoldHouse: string
    "Did you buy a house in 2010?" hasBoughtHouse: boolean
    "Did you enter a loan?" hasMaintLoan: boolean

    if (hasBoughtHouse && hasMaintLoan) {
        "What was the selling price?" sellingPrice: integer
        "Private debts for the sold house:" privateDebt: integer
        "Value residue:" valueResidue: string = "d" + hasSoldHouse

        if (privateDebt > 100) {
            "Why do you have so much debt?" whySoMuchDebt: string
            if (whySoMuchDebt == "Because I can") {
                "Are you nuts?" isNuts: boolean
            }
        }
    }
}