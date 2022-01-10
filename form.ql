form taxOfficeExample { 
    "What is your first name?" firstName: string
    "What is your last name?" lastName: string
    "Your Full name:" fullName: string = firstName + " " + lastName
    "Did you sell a house in 2010?" hasSoldHouse: boolean
    "Did you buy a house in 2010?" hasBoughtHouse: boolean
    "Did you enter a loan?" hasMaintLoan: boolean

    if (hasBoughtHouse && hasMaintLoan) {
        "What was the selling price?" sellingPrice: integer
        "Private debts for the sold house:" privateDebt: integer
        "Value residue:" valueResidue: integer = sellingPrice - privateDebt

        if (privateDebt > 100) {
            "Why do you have so much debt?" whySoMuchDebt: string
            if (whySoMuchDebt == "Because I can") {
                "Are you nuts?" isNuts: boolean
            }
        }
    }
}